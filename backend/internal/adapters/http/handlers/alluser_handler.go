package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
)

type AllUserHandler struct {
	UsersService services.UsersService
}

func NewAllUserHandler(u services.UsersService) *AllUserHandler {
	return &AllUserHandler{UsersService: u}
}

func (h *AllUserHandler) GetCurrentUser(c *fiber.Ctx) error {
	v := c.Locals("email")
	email, ok := v.(string)
	if !ok || email == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing email in context",
		})
	}

	user, err := h.UsersService.FindEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
