package handlers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
)

type StaffStatusHandler struct {
	staffStatusService services.StaffStatusService
}

func NewStaffStatusHandler(s services.StaffStatusService) *StaffStatusHandler {
	return &StaffStatusHandler{staffStatusService: s}
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	newStatus, err := h.staffStatusService.CreateStaffStatus(&status)
	if err != nil {
		log.Printf("[CreateStaffStatus] failed to create status of staff: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create status of staff",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newStatus)
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
		log.Printf("[RemoveStaffStatus] failed to convert id to uint: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	err = h.staffStatusService.RemoveStaffStatus(uint(id))
	if err != nil {
		log.Printf("[RemoveStaffStatus] failed to delete status of staff: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Delete StaffStatus successfully",
	})
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
		log.Printf("[GetStaffStatus] failed to get staffstatus: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get staffstatus",
		})
	}

	return c.Status(fiber.StatusOK).JSON(status)
}

func (h *StaffStatusHandler) UpdateStaffStatusName(c *fiber.Ctx) error {
	var req struct {
		ID     uint   `json:"id"`
		Status string `json:"status"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	updatedStatus, err := h.staffStatusService.UpdateStaffStatusName(req.ID, req.Status)
	if err != nil {
		log.Printf("[UpdateStaffStatusName] failed to update staff status name: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update staff status name",
		})
	}
	return c.Status(fiber.StatusOK).JSON(updatedStatus)
}
