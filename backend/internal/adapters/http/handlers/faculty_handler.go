package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

type FacultyHandler struct {
	facultyService services.FacultyService
	res            utils.IResponse
}

func NewFacultyHandler(s services.FacultyService) *FacultyHandler {
	return &FacultyHandler{facultyService: s, res: utils.NewResponse()}
}

// GetAllFaculty godoc
// @Summary Get all faculties
// @Description ดึงข้อมูลคณะทั้งหมดในระบบ
// @Tags Faculty
// @Produce json
// @Success 200 {array} entities.Faculty
// @Failure 500 {object} fiber.Map
// @Router /api/faculty [get]
// @Security BearerAuth
func (h *FacultyHandler) GetAllFaculty(c *fiber.Ctx) error {
	faculties, err := h.facultyService.GetAllFaculty()
	if err != nil {
		return h.res.InternalServerError(c, "Failed to get faculties", err.Error(), utils.CodeInternalError)
	}
	return h.res.Item(c, "Get all faculties successfully", faculties)
}
