package main

import (
	"fmt"
	"log"

	"github.com/AlejandroEmilianoDamian21/listGamesGO/handlers"
	"github.com/AlejandroEmilianoDamian21/listGamesGO/initializers"
	"github.com/AlejandroEmilianoDamian21/listGamesGO/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

var isLogoutAllowed = true

func main() {

	// log.Println("Start Proyect")
	//Crear nuestra aplicacion de Fiber

	app := fiber.New()
	micro := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		return c.Next()
	})

	app.Mount("/api", micro)
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", handlers.SignUpUser)
		router.Post("/login", handlers.SignInUser)
		router.Get("/logout", middleware.DeserializeUser, handlers.GetMe)
	})

	micro.Get("/users/me", middleware.DeserializeUser, handlers.GetMe)

	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Welcome to Golang, Fiber, and GORM"})
	})
	nuevoHandlerJuegos := handlers.NuevoJuegosHandler()
	/*Rutas de los juegos*/
	micro.Get("/juego", middleware.DeserializeUser, nuevoHandlerJuegos.ObtenerTodosJuegos)
	micro.Get("/juego/:id", middleware.DeserializeUser, nuevoHandlerJuegos.ObtenerJuego)
	micro.Post("/juego", middleware.DeserializeUser, nuevoHandlerJuegos.CrearJuego)
	micro.Put("/juego", middleware.DeserializeUser, nuevoHandlerJuegos.ModificarJuego)
	micro.Delete("/juego/:id", middleware.DeserializeUser, nuevoHandlerJuegos.EliminarJuego)

	nuevoHandlerGeneros := handlers.NuevoGenerosHandler()

	micro.Get("/genero", middleware.DeserializeUser, nuevoHandlerGeneros.ObtenerTodosGeneros)

	micro.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("Path: %v does not exists on this server", path)})
	})
	// log.Println(DB.RowsAffected)
	log.Fatal(app.Listen(":3030"))

}
