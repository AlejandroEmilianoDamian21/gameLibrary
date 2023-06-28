package middleware

import (
	"fmt"
	"strings"

	"github.com/AlejandroEmilianoDamian21/listGamesGO/initializers"
	"github.com/AlejandroEmilianoDamian21/listGamesGO/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string

	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer") {
		tokenString = strings.TrimPrefix(authorization, "Beaber")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	config, _ := initializers.LoadConfig(".")

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("Invalidate toke: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)

	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid token claim"})
	}

	var user models.User

	initializers.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	if user.ID.String() != claims["sub"] {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "The user belongin to this toke no logger exists"})
	}
	c.Locals("user", models.FilterUserRecord(&user))
	return c.Next()
}
