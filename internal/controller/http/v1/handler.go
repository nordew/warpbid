package v1

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nordew/go-errx"
	"github.com/nordew/warpbid/internal/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Get("/nonce", h.getNonce)
	auth.Post("/verify", h.verify)
}

func handleError(c *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	log.Printf("Error: %v", err)

	switch {
	case errx.IsCode(err, errx.NotFound):
		return writeError(c, fiber.StatusNotFound, err)
	case errx.IsCode(err, errx.BadRequest):
		return writeError(c, fiber.StatusBadRequest, err)
	case errx.IsCode(err, errx.Internal):
		return writeError(c, fiber.StatusInternalServerError, err)
	case errx.IsCode(err, errx.Unauthorized):
		return writeError(c, fiber.StatusUnauthorized, err)
	case errx.IsCode(err, errx.Forbidden):
		return writeError(c, fiber.StatusForbidden, err)
	default:
		return writeError(c, fiber.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err))
	}
}

func writeError(c *fiber.Ctx, statusCode int, err error) error {
	response := fiber.Map{
		"error": err.Error(),
	}

	return c.Status(statusCode).JSON(response)
}
