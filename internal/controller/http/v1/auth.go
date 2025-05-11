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

// // internal/controller/http/v1/auth.go
// func (h *AuthHandler) GetNonce(c *gin.Context) {
// 	addr := c.Query("walletAddress")

// 	// 1. validate eth address...
// 	nonce := uuid.New().String()
// 	expire := time.Now().Add(5 * time.Minute)

// 	h.challenges.Save(addr, nonce, expire)

// 	c.JSON(200, gin.H{"nonce": fmt.Sprintf("Sign this message: %s", nonce)})
// }

// func (h *AuthHandler) Verify(c *gin.Context) {
// 	var req dto.VerifyRequest
// 	if err := c.ShouldBindJSON(&req); err != nil { /* 400 */
// 	}

// 	ch, ok := h.challenges.Load(req.WalletAddress)

// 	if !ok || ch.ExpiresAt.Before(time.Now()) { /* 400 challenge */
// 	}

// 	recoveredAddr, err := crypto.VerifySignature(ch.Nonce, req.Signature)
// 	if err != nil || !strings.EqualFold(recoveredAddr, req.WalletAddress) {
// 		/* 401 invalid signature */
// 	}

// 	token, refresh := h.tokenSvc.GenerateTokens(req.WalletAddress)

// 	c.JSON(200, dto.VerifyResponse{AccessToken: token, RefreshToken: refresh})
// }
