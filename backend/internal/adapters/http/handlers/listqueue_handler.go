package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

type ListQueueHandler struct {
	listQueueService   services.ListQueueService
	staffStatusService services.StaffStatusService
	orderService       services.OrderService
	userService        services.UsersService
}

func NewListQueueHandler(s services.ListQueueService, sf services.StaffStatusService, o services.OrderService, u services.UsersService) *ListQueueHandler {
	return &ListQueueHandler{listQueueService: s, staffStatusService: sf, orderService: o, userService: u}
}

// CreateRequest godoc
// @Summary สร้างคำร้องใหม่
// @Tags ListQueue
// @Accept json
// @Produce json
// @Param request body entities.ListQueue true "ListQueue Request"
// @Success 201 {object} entities.ListQueue
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /list-queue [post]

func badTime(c *fiber.Ctx, field string, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": fmt.Sprintf("invalid %s: %v", field, err),
	})
}

func (h *ListQueueHandler) CreateRequest(c *fiber.Ctx) error {
	var in utils.CreateListQueueReq
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	loc, _ := time.LoadLocation("Asia/Bangkok") // ปรับตามที่ต้องการ

	wordfileSubmit, err := utils.ParseTimeFlexible(in.WordfileSubmit, loc)
	if err != nil {
		return badTime(c, "wordfile_submit", err)
	}
	infoSubmit, err := utils.ParseTimeFlexible(in.InfoSubmit, loc)
	if err != nil {
		return badTime(c, "info_submit", err)
	}
	infoSubmit14, err := utils.ParseTimeFlexible(in.InfoSubmit14Days, loc)
	if err != nil {
		return badTime(c, "info_submit_14days", err)
	}
	timeRegister, err := utils.ParseTimeFlexible(in.TimeRegister, loc)
	if err != nil {
		return badTime(c, "time_register", err)
	}
	onWeb, err := utils.ParseTimeFlexible(in.OnWeb, loc)
	if err != nil {
		return badTime(c, "on_web", err)
	}
	apptAW, err := utils.ParseTimeFlexible(in.AppointmentDateAW, loc)
	if err != nil {
		return badTime(c, "appointment_date_aw", err)
	}

	req := entities.ListQueue{
		Title:                in.Title,
		StaffID:              in.StaffID,
		FacultyID:            in.FacultyID,
		StaffStatusID:        in.StaffStatusID,
		CourseStatusID:       in.CourseStatusID,
		DateWordFileSubmit:   wordfileSubmit,
		DateInfoSubmit:       infoSubmit,
		DateInfoSubmit14Days: infoSubmit14,
		DateRegister:         timeRegister,
		OnWeb:                onWeb,
		AppointmentDateAW:    apptAW,
		Owner:                in.Owner,
		Note:                 in.Note,
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newReq, err := h.listQueueService.CreateListQueue(&req)
	if err != nil {
		log.Printf("[CreateRequest] failed: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create request"})
	}
	return c.Status(fiber.StatusCreated).JSON(newReq)
}

// UpdateListQueue godoc
// @Summary แก้ไขคำร้อง
// @Tags ListQueue
// @Accept json
// @Produce json
// @Param request body entities.ListQueue true "ListQueue Update"
// @Success 200 {object} entities.ListQueue
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /list-queue [put]
func (h *ListQueueHandler) UpdateListQueue(c *fiber.Ctx) error {
	var req entities.ListQueue
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	updatedReq, err := h.listQueueService.UpdateListQueue(&req)
	if err != nil {
		log.Printf("[UpdateListQueue] failed to update request: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update request",
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedReq)
}

// UpdateStaffStatus godoc
// @Summary อัปเดตสถานะของเจ้าหน้าที่
// @Tags ListQueue
// @Param id path int true "ListQueue ID"
// @Param staff_status_id path int true "Staff Status ID"
// @Produce json
// @Success 200 {object} entities.ListQueue
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /list-queue/{id}/staff-status/{staff_status_id} [put]
func (h *ListQueueHandler) UpdateStaffStatus(c *fiber.Ctx) error {
	intID, err1 := c.ParamsInt("id")
	intStatusID, err2 := c.ParamsInt("staff_status_id")
	if err1 != nil || err2 != nil || intID <= 0 || intStatusID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID or status ID",
		})
	}

	id := uint(intID)
	statusID := uint(intStatusID)

	statusUpdated, err := h.listQueueService.UpdateStaffStatus(uint(id), uint(statusID))
	if err != nil {
		log.Printf("[UpdateStaffStatusListQueue] failed to update staff status: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update staff status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(statusUpdated)
}

func (h *ListQueueHandler) UpdatePriority(c *fiber.Ctx) error {
	intID, err1 := c.ParamsInt("id")
	intPriority, err2 := c.ParamsInt("priority")

	if err1 != nil || intID <= 0 || err2 != nil || intPriority < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID or priority",
		})
	}

	result, err := h.listQueueService.UpdatePriority(uint(intID), uint(intPriority))
	if err != nil {
		log.Printf("[UpdatePriority] failed to update priority: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update priority",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// GetListQueue godoc
// @Summary ดึงคำร้องทั้งหมด
// @Tags ListQueue
// @Produce json
// @Success 200 {array} entities.ListQueue
// @Failure 500 {object} map[string]string
// @Router /list-queue [get]
func (h *ListQueueHandler) GetListQueue(c *fiber.Ctx) error {
	list, err := h.listQueueService.GetListQueue()
	if err != nil {
		log.Printf("[GetListQueue] failed to get list queue: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get list queue",
		})
	}

	return c.Status(fiber.StatusOK).JSON(list)
}

// GetListQueueNotYet godoc
// @Summary Get list queue not yet finished
// @Description ดึงรายการที่ยังไม่เสร็จ (เช่น priority != 0 หรือเงื่อนไขอื่นที่คุณกำหนดไว้ใน service)
// @Tags ListQueue
// @Produce json
// @Success 200 {array} entities.ListQueue
// @Failure 500 {object} fiber.Map
// @Router /listqueue/notyet [get]
// @Security BearerAuth
func (h *ListQueueHandler) GetListQueueNotYet(c *fiber.Ctx) error {
	list, err := h.listQueueService.GetListQueueNotYet()
	if err != nil {
		log.Printf("[GetListQueueNotYet] failed to get list queue where not finished: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get list queue where not finished",
		})
	}

	return c.Status(fiber.StatusOK).JSON(list)
}

// GetListQueueByStaffStatus godoc
// @Summary ดึงคำร้องตาม Staff Status
// @Tags ListQueue
// @Param staff_status_id path int true "Staff Status ID"
// @Produce json
// @Success 200 {array} entities.ListQueue
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /list-queue/staff-status/{staff_status_id} [get]
func (h *ListQueueHandler) GetListQueueByStaffStatus(c *fiber.Ctx) error {
	var ids []uint
	if err := c.BodyParser(&ids); err != nil || len(ids) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "expected [2,5,7]"})
	}
	result, err := h.listQueueService.GetListQueueByStaffStatus(ids)
	if err != nil {
		log.Printf("[GetListQueueByStaffStatus] failed to get list queue by staff status: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get list queue by staff status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// GetListQueueByFaculty godoc
// @Summary Get list queue by faculty
// @Description ดึงรายการคิวตามชื่อคณะของผู้ใช้ (จาก JWT Claims: organizationname_th)
// @Tags ListQueue
// @Produce json
// @Success 200 {array} entities.ListQueue
// @Failure 500 {object} fiber.Map
// @Router /listqueue/faculty [get]
// @Security BearerAuth
func (h *ListQueueHandler) GetListQueueByFaculty(c *fiber.Ctx) error {
	var ids []uint

	if err := c.BodyParser(&ids); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "expected [2,5,7]"})
	}

	user, err := h.userService.FindEmail(c.Locals("email").(string))
	if err != nil {
		log.Printf("[GetListQueueByFaculty] failed to get user by email: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user by email",
		})
	}

	list, err := h.listQueueService.GetListQueueByFaculty(user.OrganizationNameTH, ids)
	if err != nil {
		log.Printf("[GetListQueueByFaculty] failed to get list queue by faculty: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get list queue by faculty",
		})
	}

	return c.Status(fiber.StatusOK).JSON(list)
}

func (h *ListQueueHandler) GetListQueueByCourseStatus(c *fiber.Ctx) error {
	var ids []uint

	if err := c.BodyParser(&ids); err != nil || len(ids) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "expected [2,5,7]"})
	}
	result, err := h.listQueueService.GetListQueueByCourseStatus(ids)
	if err != nil {
		log.Printf("[GetListQueueByStaffStatus] failed to get list queue by staff status: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get list queue by staff status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *ListQueueHandler) GetListQueueByOwner(c *fiber.Ctx) error {
	var ids []uint
	email := c.Locals("email").(string)

	if err := c.BodyParser(&ids); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "expected [2,5,7]"})
	}

	result, err := h.listQueueService.GetListQueueByOwner(email, ids)
	if err != nil {
		log.Printf("[GetListQueueByOwner] failed to get list queue by owner: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get list queue by owner",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// RemoveListQueueForDev godoc
// @Summary Remove a list queue by ID (For Dev Only)
// @Description ใช้ลบรายการคิวตาม ID (ใช้เฉพาะในโหมดพัฒนาเท่านั้น)
// @Tags ListQueue
// @Param id path int true "ListQueue ID"
// @Produce json
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /listqueue/dev/{id} [delete]
// @Security BearerAuth
func (h *ListQueueHandler) RemoveListQueueForDev(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid status ID",
		})
	}

	err = h.listQueueService.RemoveListQueueForDev(uint(id))
	if err != nil {
		log.Printf("[RemoveListQueueForDev] failed to remove list queue: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to remove list queue",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "List queue removed successfully",
	})
}
