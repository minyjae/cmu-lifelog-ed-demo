package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

type RoleHandler struct {
	RoleService services.RoleService
	res         utils.IResponse
}

func NewRoleHandler(s services.RoleService) *RoleHandler {
	return &RoleHandler{RoleService: s, res: utils.NewResponse()}
}

// GetRole godoc
// @Summary Get all roles
// @Description ดึงบทบาททั้งหมดในระบบ
// @Tags Role
// @Produce json
// @Success 200 {array} entities.Role
// @Failure 500 {object} fiber.Map
// @Router /role [get]
// @Security BearerAuth
func (h *RoleHandler) GetRole(c *fiber.Ctx) error {
	roles, err := h.RoleService.GetRole()
	if err != nil {
		return h.res.InternalServerError(c, "Failed to get roles", err.Error(), utils.CodeInternalError)
	}
	return h.res.Item(c, "Get all roles successfully", roles)
}
