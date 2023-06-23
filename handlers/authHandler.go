package handlers

import (
	"strings"

	"github.com/AlejandroEmilianoDamian21/listGamesGO/models"
	"github.com/AlejandroEmilianoDamian21/listGamesGO/storage"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *fiber.Ctx) error {
	var payload *models.SignUpInput

	// Analizar el cuerpo de la solicitud y asignarlo a la variable payload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Validar la estructura de los datos del payload
	errors := models.ValidateStruct(payload)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": errors})
	}

	// Verificar que las contraseñas coincidan
	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwors do not match"})
	}

	// Generar un hash de la contraseña proporcionada
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Crear un nuevo usuario con los datos proporcionados
	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
		Photo:    &payload.Photo,
	}

	// Crear el registro del nuevo usuario en la base de datos
	result := storage.DB.Create(&newUser)

	// Verificar si ocurrieron errores durante la creación del usuario
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	// Devolver la respuesta con el usuario filtrado
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})
}
