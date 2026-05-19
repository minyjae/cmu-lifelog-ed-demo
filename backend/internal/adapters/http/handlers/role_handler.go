package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
)

type RoleHandler struct {
	RoleService services.RoleService
}

func NewRoleHandler(s services.RoleService) *RoleHandler {
	return &RoleHandler{RoleService: s}
}

// GetCourseStatus godoc
// @Summary Get all course statuses
// @Description ดึงสถานะของคอร์สทั้งหมด
// @Tags CourseStatus
// @Produce json
// @Success 200 {array} entities.CourseStatus
// @Failure 500 {object} fiber.Map
// @Router /coursestatus [get]
// @Security BearerAuth
func (h *RoleHandler) GetRole(c *fiber.Ctx) error {
	roles, err := h.RoleService.GetRole()
	if err != nil {
		log.Printf("Error getting course statuses: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get course statuses"})
	}
	return c.Status(fiber.StatusOK).JSON(roles)
}
