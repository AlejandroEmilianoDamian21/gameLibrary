package handlers

import (
	"strings"

	"github.com/AlejandroEmilianoDamian21/listGamesGO/models"
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

/*Manejador de obtener los generos*/
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

/*Manejador de crear generos*/
func (g *generosHandler) CrearGenero(c *fiber.Ctx) error {
	var nuevoGenero *models.Genero

	if !isLogoutAllowed {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "Access denied"})
	}

	/*Obtener los datos del body*/
	if err := c.BodyParser(&nuevoGenero); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Revisa tu body", "data": err.Error()})
	}
	/*Comprobacion de los datos*/
	/*Nombre*/
	if len(strings.TrimSpace(nuevoGenero.Nombre)) <= 0 || strings.TrimSpace(nuevoGenero.Nombre) == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "El nombre del genero es obligatorio"})
	}
	/*Descripcion*/
	if len(strings.TrimSpace(nuevoGenero.Descripcion)) <= 0 || strings.TrimSpace(nuevoGenero.Descripcion) == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "La descripcion es obligatoria"})
	}

	/*Imagen URL*/
	if len(strings.TrimSpace(nuevoGenero.ImagenURL)) <= 0 || strings.TrimSpace(nuevoGenero.ImagenURL) == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Pasar imagen es obligatoria"})
	}

	genero, exitoso, err := g.BaseDatos.CrearGenero(nuevoGenero)

	if err != nil || !exitoso {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error in DB", "data": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(genero)
}

/********************************************************************************/

/*Modificar un genero*/

func (g *generosHandler) ModificarGenero(c *fiber.Ctx) error {

	var body *models.Genero

	if !isLogoutAllowed {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "Access denied"})
	}

	/*Obtener los datos del body*/
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Revisa tu body", "data": err.Error()})
	}

	var nuevoGenero *models.Genero

	/*Comprobacion de los datos de la Estructura Genero*/
	/*Nombre*/
	if len(strings.TrimSpace(nuevoGenero.Nombre)) <= 0 || strings.TrimSpace(body.Nombre) == "" {
		nuevoGenero.Nombre = body.Nombre
	}

	if len(strings.TrimSpace(nuevoGenero.Descripcion)) <= 0 || strings.TrimSpace(body.Descripcion) == "" {
		nuevoGenero.Descripcion = body.Descripcion
	}

	if len(strings.TrimSpace(nuevoGenero.ImagenURL)) <= 0 || strings.TrimSpace(body.ImagenURL) == "" {
		nuevoGenero.ImagenURL = body.ImagenURL
	}

	exitoso, err := g.BaseDatos.ModificarGenero(body)

	if err != nil || !exitoso {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Errorn in DB", "data": err.Error()})
	}
	return c.SendStatus(fiber.StatusAccepted)
}
