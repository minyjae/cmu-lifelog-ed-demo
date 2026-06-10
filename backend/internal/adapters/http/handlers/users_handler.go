package handlers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

type UsersHandler struct {
	usersService   services.UsersService
	facultyService services.FacultyService
	res            utils.IResponse
}

func NewUsersHandler(u services.UsersService, f services.FacultyService) *UsersHandler {
	return &UsersHandler{usersService: u, facultyService: f, res: utils.NewResponse()}
}

// CreateUser godoc
// @Summary สร้างผู้ใช้ใหม่
// @Description เพิ่ม user ด้วย role ที่ส่งเข้ามาผ่าน path
// @Tags Users
// @Produce json
// @Param email path string true "User Email"
// @Param role path string true "Role"
// @Success 201 {object} entities.Users
// @Failure 500 {object} map[string]string
// @Router /api/user/{email}/{role} [post]
// @Security BearerAuth
func (h *UsersHandler) CreateUser(c *fiber.Ctx) error {
	email := c.Params("email")
	role := c.Params("role")

	user, err := h.usersService.CreateUser(role, email)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to create user", err.Error(), utils.CodeInternalError)
	}

	return h.res.Created(c, "Create user successfully", user)
}

// RemoveUser godoc
// @Summary ลบผู้ใช้
// @Tags Users
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/{id} [delete]
// @Security BearerAuth
func (h *UsersHandler) RemoveUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return h.res.BadRequest(c, "Invalid user ID", utils.CodeInvalidID)
	}

	if err := h.usersService.RemoveUser(uint(id)); err != nil {
		return h.res.InternalServerError(c, "Failed to delete user", err.Error(), utils.CodeInternalError)
	}

	return h.res.Deleted(c, "Delete user successfully")
}

// UpdateUserInfo godoc
// @Summary แก้ไขข้อมูลผู้ใช้ (role และคณะ)
// @Tags Users
// @Accept json
// @Produce json
// @Param email path string true "User Email"
// @Param request body object true "Role and Organization"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/updateinfo/{email} [put]
// @Security BearerAuth
func (h *UsersHandler) UpdateUserInfo(c *fiber.Ctx) error {
	email := c.Params("email")
	user := new(struct {
		Role               string `json:"role"`
		OrganizationNameTH string `json:"organization_name_th"`
	})
	if err := c.BodyParser(user); err != nil {
		log.Printf("[UpdateInfo] failed to parse body: %v", err)
		return h.res.BadRequest(c, "Failed to parse body", utils.CodeInvalidRequest)
	}
	f, err := h.facultyService.CheckFacultyExist(user.OrganizationNameTH)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to check faculty", err.Error(), utils.CodeInternalError)
	}
	u := &entities.Users{
		Role:               user.Role,
		OrganizationNameTH: f.NameTH,
	}

	if err := h.usersService.UpdateInfo(email, u); err != nil {
		return h.res.InternalServerError(c, "Failed to update user info", err.Error(), utils.CodeInternalError)
	}
	return h.res.Updated(c, "Update user info successfully", nil)
}

// GetStaff godoc
// @Summary ดึงข้อมูล staff ทั้งหมด
// @Description ใช้สำหรับโชว์ตอนสร้าง listqueue ว่าใครเป็นผู้ดูแลหลักสูตร
// @Tags Users
// @Produce json
// @Success 200 {array} entities.Users
// @Failure 500 {object} map[string]string
// @Router /api/staff [get]
// @Security BearerAuth
func (s *UsersHandler) GetStaff(c *fiber.Ctx) error {
	staff, err := s.usersService.GetStaff()
	if err != nil {
		return s.res.InternalServerError(c, "Failed to get staff", err.Error(), utils.CodeInternalError)
	}

	return s.res.Item(c, "Get staff successfully", staff)
}

// GetAllUsers godoc
// @Summary ดึงข้อมูลผู้ใช้ทั้งหมด
// @Description ใช้สำหรับหน้า manage ผู้ใช้ของ admin
// @Tags Users
// @Produce json
// @Success 200 {array} entities.Users
// @Failure 500 {object} map[string]string
// @Router /api/user/all [get]
// @Security BearerAuth
func (s *UsersHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := s.usersService.GetAllUsers()
	if err != nil {
		return s.res.InternalServerError(c, "Failed to get all users", err.Error(), utils.CodeInternalError)
	}

	return s.res.Item(c, "Get all users successfully", users)
}
