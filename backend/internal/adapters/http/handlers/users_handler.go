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
// @Tags Users
// @Accept json
// @Produce json
// @Param request body entities.User true "User Payload"
// @Success 201 {object} entities.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
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
// @Router /users/{id} [delete]
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

func (s *UsersHandler) GetStaff(c *fiber.Ctx) error {
	staff, err := s.usersService.GetStaff()
	if err != nil {
		return s.res.InternalServerError(c, "Failed to get staff", err.Error(), utils.CodeInternalError)
	}

	return s.res.Item(c, "Get staff successfully", staff)
}

func (s *UsersHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := s.usersService.GetAllUsers()
	if err != nil {
		return s.res.InternalServerError(c, "Failed to get all users", err.Error(), utils.CodeInternalError)
	}

	return s.res.Item(c, "Get all users successfully", users)
}
