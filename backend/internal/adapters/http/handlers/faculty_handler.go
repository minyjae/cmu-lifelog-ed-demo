package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
)

type FacultyHandler struct {
	facultyService services.FacultyService
}

func NewFacultyHandler(s services.FacultyService) *FacultyHandler {
	return &FacultyHandler{facultyService: s}
}

// GetAllFaculty godoc
// @Summary Get all faculties
// @Description ดึงข้อมูลคณะทั้งหมดในระบบ
// @Tags Faculty
// @Produce json
// @Success 200 {array} entities.Faculty
// @Failure 500 {object} fiber.Map
// @Router /faculty [get]
// @Security BearerAuth
func (h *FacultyHandler) GetAllFaculty(c *fiber.Ctx) error {
	faculties, err := h.facultyService.GetAllFaculty()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get faculties"})
	}
	return c.Status(fiber.StatusOK).JSON(faculties)
}
