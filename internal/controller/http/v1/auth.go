package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nordew/warpbid/internal/dto"
)

func (h *Handler) getNonce(c *fiber.Ctx) error {
	walletAddress := c.Query("walletAddress")

	nonce, err := h.service.GetNonce(c.Context(), walletAddress)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"nonce": "Sign this message: " + nonce,
	})
}

func (h *Handler) verify(c *fiber.Ctx) error {
	var req dto.VerifyAuthRequest

	if err := c.BodyParser(&req); err != nil {
		return handleError(c, err)
	}

	accessToken, refreshToken, err := h.service.Verify(c.Context(), req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
