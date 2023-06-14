package handlers

import (
	"strconv"

	"github.com/AlejandroEmilianoDamian21/listGamesGO/storage"
	"github.com/gofiber/fiber/v2"
)

type juegosHandler struct {
	BaseDatos storage.JuegoDB
}

func NuevoJuegosHandler() *juegosHandler {
	/*Regresa nuevo handler*/
	return &juegosHandler{
		BaseDatos: storage.NuevoJuegoDB(),
	}
}

func (j *juegosHandler) ObtenerJuego(c *fiber.Ctx) error {
	/*Obtener el ID desde los parametros, es uno de los metodos de fiber*/
	ID := c.Params("id")
	/*Si el ID esta vacio, no se envio asi que regresa un error*/
	if len(ID) < 0 {
		/*Esta es la forma de regresar errores  en fiber*/
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}
	intID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error converting in Integer", "data": err.Error()})
	}

	juego, err := j.BaseDatos.ObtenerJuego(intID)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error in DB", "data": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(juego)
}

func (j *juegosHandler) EliminarJuego(c *fiber.Ctx) error {
	ID := c.Params("id")

	if len(ID) < 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}

	intID, err := strconv.Atoi(ID)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error converting in Integer", "data": err.Error()})
	}

	exitoso, err := j.BaseDatos.EliminarJuego(intID)

	if err != nil || !exitoso {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error in DB", "data": err.Error()})
	}

	return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "ok", "message": "Game delete"})

}
