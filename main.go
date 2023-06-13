package main

import (
	"log"

	"github.com/AlejandroEmilianoDamian21/listGamesGO/storage"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// log.Println("Start Proyect")
	//Crear nuestra aplicacion de Fiber

	app := fiber.New()

	/*Esto creara una peticion GET en la ruta base
	Regresara un simple string
	Primero se agrega un string que sera el path de la ruta base y luego
	se agrega la funcion handler, el cual tiene siempre esa forma
	*/

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!!")
	})

	DB := storage.ConnectDB()
	defer DB.Close()

	log.Println(DB.RowsAffected)

	log.Fatal(app.Listen(":3001"))
}
