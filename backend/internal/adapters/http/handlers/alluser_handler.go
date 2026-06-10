package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

type AllUserHandler struct {
	UsersService services.UsersService
	res          utils.IResponse
}

func NewAllUserHandler(u services.UsersService) *AllUserHandler {
	return &AllUserHandler{UsersService: u, res: utils.NewResponse()}
}

// GetCurrentUser godoc
// @Summary ดึงข้อมูลผู้ใช้ที่ล็อกอินอยู่
// @Description ดึงข้อมูลจาก database ของผู้ใช้ที่กำลัง signin อยู่ (อิงจาก email ใน JWT)
// @Tags Users
// @Produce json
// @Success 200 {object} entities.Users
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/user/me [get]
// @Security BearerAuth
func (h *AllUserHandler) GetCurrentUser(c *fiber.Ctx) error {
	v := c.Locals("email")
	email, ok := v.(string)
	if !ok || email == "" {
		return h.res.Unauthorized(c, "Missing email in context", utils.CodeUnauthorized)
	}

	user, err := h.UsersService.FindEmail(email)
	if err != nil {
		return h.res.NotFound(c, "User not found", utils.CodeUserNotFound)
	}

	return h.res.Item(c, "Get current user successfully", user)
}
