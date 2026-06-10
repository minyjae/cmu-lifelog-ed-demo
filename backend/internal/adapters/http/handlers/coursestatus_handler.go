package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

type CourseStatusHandler struct {
	courseStatusService services.CourseStatusService
	res                 utils.IResponse
}

func NewCourseStatusHandler(s services.CourseStatusService) *CourseStatusHandler {
	return &CourseStatusHandler{courseStatusService: s, res: utils.NewResponse()}
}

// GetCourseStatus godoc
// @Summary Get all course statuses
// @Description ดึงสถานะของคอร์สทั้งหมด
// @Tags CourseStatus
// @Produce json
// @Success 200 {array} entities.CourseStatus
// @Failure 500 {object} fiber.Map
// @Router /api/course/status [get]
// @Security BearerAuth
func (h *CourseStatusHandler) GetCourseStatus(c *fiber.Ctx) error {
	statuses, err := h.courseStatusService.GetCourseStatus()
	if err != nil {
		return h.res.InternalServerError(c, "Failed to get course statuses", err.Error(), utils.CodeInternalError)
	}
	return h.res.Item(c, "Get all course statuses successfully", statuses)
}
