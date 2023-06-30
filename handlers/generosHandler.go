package handlers

import (
	"github.com/AlejandroEmilianoDamian21/listGamesGO/storage"
	"github.com/gofiber/fiber/v2"
)

type generosHandler struct {
	BaseDatos storage.GeneroDB
}

func NuevoGenerosHandler() *generosHandler {
	/*Regresa nuevo handler*/
	return &generosHandler{
		BaseDatos: storage.NuevoGeneroDB(),
	}
}

func (g *generosHandler) ObtenerTodosGeneros(c *fiber.Ctx) error {
	if !isLogoutAllowed {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "Access denied"})
	}
	genero, err := g.BaseDatos.ObtenerGeneros()
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "fail", "message": "Error in DB", "data": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(genero)
}
