package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
)

type CourseStatusHandler struct {
	courseStatusService services.CourseStatusService
}

func NewCourseStatusHandler(s services.CourseStatusService) *CourseStatusHandler {
	return &CourseStatusHandler{courseStatusService: s}
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
func (h *CourseStatusHandler) GetCourseStatus(c *fiber.Ctx) error {
	statuses, err := h.courseStatusService.GetCourseStatus()
	if err != nil {
		log.Printf("Error getting course statuses: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get course statuses"})
	}
	return c.Status(fiber.StatusOK).JSON(statuses)
}
