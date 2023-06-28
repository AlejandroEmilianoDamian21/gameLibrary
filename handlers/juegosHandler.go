package handlers

import "fmt"

func hi() {
	fmt.Println("Hola")
}

// import (
// 	"strconv"
// 	"strings"

// 	"github.com/AlejandroEmilianoDamian21/listGamesGO/models"
// 	"github.com/AlejandroEmilianoDamian21/listGamesGO/storage"
// 	"github.com/gofiber/fiber/v2"
// )

// type juegosHandler struct {
// 	BaseDatos storage.JuegoDB
// }

// func NuevoJuegosHandler() *juegosHandler {
// 	/*Regresa nuevo handler*/
// 	return &juegosHandler{
// 		BaseDatos: storage.NuevoJuegoDB(),
// 	}
// }

// /*OBTENER JUEGO POR ID*/
// func (j *juegosHandler) ObtenerJuego(c *fiber.Ctx) error {
// 	/*Obtener el ID desde los parametros, es uno de los metodos de fiber*/
// 	ID := c.Params("ID")
// 	/*Si el ID esta vacio, no se envio asi que regresa un error*/
// 	if len(ID) < 0 {
// 		/*Esta es la forma de regresar errores  en fiber*/
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Review your input"})
// 	}
// 	intID, err := strconv.Atoi(ID)

// 	if err != nil {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error converting in Integer", "data": err.Error()})
// 	}

// 	juego, err := j.BaseDatos.ObtenerJuego(intID)

// 	if err != nil {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error in DB", "data": err.Error()})
// 	}
// 	return c.Status(fiber.StatusAccepted).JSON(juego)
// }

// /*OBTENER TODOS LOS JUEGOS*/
// func (j *juegosHandler) ObtenerTodosJuegos(c *fiber.Ctx) error {
// 	juego, err := j.BaseDatos.ObtenerJuegos()
// 	if err != nil {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error in DB", "data": err.Error()})
// 	}
// 	return c.Status(fiber.StatusAccepted).JSON(juego)
// }

// /********************/

// /*CREAR UN NUEVO JUEGO*/
// func (j *juegosHandler) CrearJuego(c *fiber.Ctx) error {
// 	var nuevoJuego *models.Juego
// 	/*Obtener los datos del body*/
// 	if err := c.BodyParser(&nuevoJuego); err != nil {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Revisa tu body", "data": err.Error()})
// 	}
// 	/*Comprobar los datos del Juego*/
// 	/*Nombre*/
// 	if len(strings.TrimSpace(nuevoJuego.Nombre)) <= 0 || strings.TrimSpace(nuevoJuego.Nombre) == "" {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "El nombre del juego es obligatorio"})
// 	}
// 	/*Desarrollador*/
// 	if len(strings.TrimSpace(nuevoJuego.Desarrollador)) <= 0 || strings.TrimSpace(nuevoJuego.Desarrollador) == "" {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "El desarrollador del juego es obligatorio"})

// 	}
// 	/*Precio*/
// 	if nuevoJuego.Precio < 0 {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "El precio del juego no debe ser negativo"})
// 	}

// 	juego, existoso, err := j.BaseDatos.CrearJuego(nuevoJuego)

// 	if err != nil || !existoso {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error in DB", "data": err.Error()})
// 	}

// 	return c.Status(fiber.StatusAccepted).JSON(juego)
// }

// /********************/

// /*MODIFICAR UN JUEGO*/
// func (j *juegosHandler) ModificarJuego(c *fiber.Ctx) error {

// 	var body *models.Juego
// 	/*Obtener los datos del body*/
// 	if err := c.BodyParser(&body); err != nil {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Revisa tu body", "data": err.Error()})
// 	}

// 	/*Modificar el juego*/
// 	var nuevoJuego *models.Juego

// 	/*Comprobacion de los datos*/
// 	/*Nombre*/
// 	if len(strings.TrimSpace(body.Nombre)) < 0 || strings.TrimSpace(body.Nombre) == "" {
// 		nuevoJuego.Nombre = body.Nombre
// 	}

// 	if len(strings.TrimSpace(body.Desarrollador)) < 0 || strings.TrimSpace(body.Desarrollador) == "" {
// 		nuevoJuego.Desarrollador = body.Desarrollador
// 	}

// 	if body.Precio < 0 {
// 		nuevoJuego.Precio = body.Precio
// 	}
// 	/*Modificar el juego*/
// 	exitoso, err := j.BaseDatos.ModificarJuego(body)

// 	if err != nil || !exitoso {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error in DB", "data": err.Error()})
// 	}
// 	return c.SendStatus(fiber.StatusAccepted)
// }

// /********************/

// /*ELIMINAR JUEGO*/
// func (j *juegosHandler) EliminarJuego(c *fiber.Ctx) error {
// 	ID := c.Params("id")

// 	if len(ID) < 0 {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Review your input"})
// 	}

// 	intID, err := strconv.Atoi(ID)

// 	if err != nil {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error converting in Integer", "data": err.Error()})
// 	}

// 	exitoso, err := j.BaseDatos.EliminarJuego(intID)

// 	if err != nil || !exitoso {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error in DB", "data": err.Error()})
// 	}

// 	return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "ok", "message": "Game delete"})

// }

// /********************/
