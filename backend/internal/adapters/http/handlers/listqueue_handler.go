package handlers

import (
	"fmt"
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
	res                utils.IResponse
}

func NewListQueueHandler(s services.ListQueueService, sf services.StaffStatusService, o services.OrderService, u services.UsersService) *ListQueueHandler {
	return &ListQueueHandler{listQueueService: s, staffStatusService: sf, orderService: o, userService: u, res: utils.NewResponse()}
}

func (h *ListQueueHandler) badTime(c *fiber.Ctx, field string, err error) error {
	return h.res.BadRequest(c, fmt.Sprintf("invalid %s: %v", field, err), utils.CodeInvalidRequest)
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
func (h *ListQueueHandler) CreateRequest(c *fiber.Ctx) error {
	var in utils.CreateListQueueReq
	if err := c.BodyParser(&in); err != nil {
		return h.res.BadRequest(c, "Cannot parse JSON", utils.CodeInvalidRequest)
	}

	loc, _ := time.LoadLocation("Asia/Bangkok") // ปรับตามที่ต้องการ

	wordfileSubmit, err := utils.ParseTimeFlexible(in.WordfileSubmit, loc)
	if err != nil {
		return h.badTime(c, "wordfile_submit", err)
	}
	infoSubmit, err := utils.ParseTimeFlexible(in.InfoSubmit, loc)
	if err != nil {
		return h.badTime(c, "info_submit", err)
	}
	infoSubmit14, err := utils.ParseTimeFlexible(in.InfoSubmit14Days, loc)
	if err != nil {
		return h.badTime(c, "info_submit_14days", err)
	}
	timeRegister, err := utils.ParseTimeFlexible(in.TimeRegister, loc)
	if err != nil {
		return h.badTime(c, "time_register", err)
	}
	onWeb, err := utils.ParseTimeFlexible(in.OnWeb, loc)
	if err != nil {
		return h.badTime(c, "on_web", err)
	}
	apptAW, err := utils.ParseTimeFlexible(in.AppointmentDateAW, loc)
	if err != nil {
		return h.badTime(c, "appointment_date_aw", err)
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
		return h.res.BadRequest(c, err.Error(), utils.CodeValidationFailed)
	}

	newReq, err := h.listQueueService.CreateListQueue(&req)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to create request", err.Error(), utils.CodeInternalError)
	}
	return h.res.Created(c, "Create request successfully", newReq)
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
		return h.res.BadRequest(c, "Cannot parse JSON", utils.CodeInvalidRequest)
	}

	updatedReq, err := h.listQueueService.UpdateListQueue(&req)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to update request", err.Error(), utils.CodeInternalError)
	}

	return h.res.Updated(c, "Update request successfully", updatedReq)
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
		return h.res.BadRequest(c, "Invalid ID or status ID", utils.CodeInvalidID)
	}

	id := uint(intID)
	statusID := uint(intStatusID)

	statusUpdated, err := h.listQueueService.UpdateStaffStatus(id, statusID)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to update staff status", err.Error(), utils.CodeInternalError)
	}

	return h.res.Updated(c, "Update staff status successfully", statusUpdated)
}

func (h *ListQueueHandler) UpdatePriority(c *fiber.Ctx) error {
	intID, err1 := c.ParamsInt("id")
	intPriority, err2 := c.ParamsInt("priority")

	if err1 != nil || intID <= 0 || err2 != nil || intPriority < 0 {
		return h.res.BadRequest(c, "Invalid ID or priority", utils.CodeInvalidID)
	}

	result, err := h.listQueueService.UpdatePriority(uint(intID), uint(intPriority))
	if err != nil {
		return h.res.InternalServerError(c, "Failed to update priority", err.Error(), utils.CodeInternalError)
	}

	return h.res.Updated(c, "Update priority successfully", result)
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
		return h.res.InternalServerError(c, "Failed to get list queue", err.Error(), utils.CodeInternalError)
	}

	return h.res.Item(c, "Get list queue successfully", list)
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
		return h.res.InternalServerError(c, "Failed to get list queue where not finished", err.Error(), utils.CodeInternalError)
	}

	return h.res.Item(c, "Get list queue successfully", list)
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
		return h.res.BadRequest(c, "expected [2,5,7]", utils.CodeInvalidRequest)
	}
	result, err := h.listQueueService.GetListQueueByStaffStatus(ids)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to get list queue by staff status", err.Error(), utils.CodeInternalError)
	}

	return h.res.Item(c, "Get list queue successfully", result)
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
		return h.res.BadRequest(c, "expected [2,5,7]", utils.CodeInvalidRequest)
	}

	user, err := h.userService.FindEmail(c.Locals("email").(string))
	if err != nil {
		return h.res.InternalServerError(c, "Failed to get user by email", err.Error(), utils.CodeInternalError)
	}

	list, err := h.listQueueService.GetListQueueByFaculty(user.OrganizationNameTH, ids)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to get list queue by faculty", err.Error(), utils.CodeInternalError)
	}

	return h.res.Item(c, "Get list queue successfully", list)
}

func (h *ListQueueHandler) GetListQueueByCourseStatus(c *fiber.Ctx) error {
	var ids []uint

	if err := c.BodyParser(&ids); err != nil || len(ids) == 0 {
		return h.res.BadRequest(c, "expected [2,5,7]", utils.CodeInvalidRequest)
	}
	result, err := h.listQueueService.GetListQueueByCourseStatus(ids)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to get list queue by course status", err.Error(), utils.CodeInternalError)
	}

	return h.res.Item(c, "Get list queue successfully", result)
}

func (h *ListQueueHandler) GetListQueueByOwner(c *fiber.Ctx) error {
	var ids []uint
	email := c.Locals("email").(string)

	if err := c.BodyParser(&ids); err != nil {
		return h.res.BadRequest(c, "expected [2,5,7]", utils.CodeInvalidRequest)
	}

	result, err := h.listQueueService.GetListQueueByOwner(email, ids)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to get list queue by owner", err.Error(), utils.CodeInternalError)
	}

	return h.res.Item(c, "Get list queue successfully", result)
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
		return h.res.BadRequest(c, "Invalid status ID", utils.CodeInvalidID)
	}

	if err := h.listQueueService.RemoveListQueueForDev(uint(id)); err != nil {
		return h.res.InternalServerError(c, "Failed to remove list queue", err.Error(), utils.CodeInternalError)
	}

	return h.res.Deleted(c, "List queue removed successfully")
}
