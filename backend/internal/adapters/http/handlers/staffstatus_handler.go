package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

type StaffStatusHandler struct {
	staffStatusService services.StaffStatusService
	res                utils.IResponse
}

func NewStaffStatusHandler(s services.StaffStatusService) *StaffStatusHandler {
	return &StaffStatusHandler{staffStatusService: s, res: utils.NewResponse()}
}

// CreateStaffStatus godoc
// @Summary สร้างสถานะของเจ้าหน้าที่ใหม่
// @Tags StaffStatus
// @Accept json
// @Produce json
// @Param request body entities.StaffStatus true "Staff Status Payload"
// @Success 201 {object} entities.StaffStatus
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /staff-status [post]
func (h *StaffStatusHandler) CreateStaffStatus(c *fiber.Ctx) error {
	var status entities.StaffStatus
	if err := c.BodyParser(&status); err != nil {
		return h.res.BadRequest(c, "Cannot parse JSON", utils.CodeInvalidRequest)
	}

	newStatus, err := h.staffStatusService.CreateStaffStatus(&status)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to create status of staff", err.Error(), utils.CodeInternalError)
	}

	return h.res.Created(c, "Create staff status successfully", newStatus)
}

// RemoveStaffStatus godoc
// @Summary ลบสถานะของเจ้าหน้าที่
// @Tags StaffStatus
// @Param id path int true "Staff Status ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /staff-status/{id} [delete]
func (h *StaffStatusHandler) RemoveStaffStatus(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return h.res.BadRequest(c, "Invalid ID format", utils.CodeInvalidID)
	}

	if err := h.staffStatusService.RemoveStaffStatus(uint(id)); err != nil {
		return h.res.BadRequest(c, err.Error(), utils.CodeInvalidRequest)
	}

	return h.res.Deleted(c, "Delete staff status successfully")
}

// GetStaffStatus godoc
// @Summary ดึงสถานะของเจ้าหน้าที่ทั้งหมด
// @Tags StaffStatus
// @Produce json
// @Success 200 {array} entities.StaffStatus
// @Failure 500 {object} map[string]string
// @Router /staff-status [get]
func (h *StaffStatusHandler) GetStaffStatus(c *fiber.Ctx) error {
	status, err := h.staffStatusService.GetStaffStatus()
	if err != nil {
		return h.res.InternalServerError(c, "Failed to get staff status", err.Error(), utils.CodeInternalError)
	}

	return h.res.Item(c, "Get staff status successfully", status)
}

func (h *StaffStatusHandler) UpdateStaffStatusName(c *fiber.Ctx) error {
	var req struct {
		ID     uint   `json:"id"`
		Status string `json:"status"`
	}
	if err := c.BodyParser(&req); err != nil {
		return h.res.BadRequest(c, "Cannot parse JSON", utils.CodeInvalidRequest)
	}
	updatedStatus, err := h.staffStatusService.UpdateStaffStatusName(req.ID, req.Status)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to update staff status name", err.Error(), utils.CodeInternalError)
	}
	return h.res.Updated(c, "Update staff status name successfully", updatedStatus)
}
