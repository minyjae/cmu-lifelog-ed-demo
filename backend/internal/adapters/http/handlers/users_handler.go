package handlers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
)

type UsersHandler struct {
	usersService   services.UsersService
	facultyService services.FacultyService
}

func NewUsersHandler(u services.UsersService, f services.FacultyService) *UsersHandler {
	return &UsersHandler{usersService: u, facultyService: f}
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

	_, err := h.usersService.CreateUser(role, email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"meassage": "complete to create staff",
	})
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
		log.Printf("[DeleteUser] failed to convert id to uint: %v", err)
	}

	err = h.usersService.RemoveUser(uint(id))
	if err != nil {
		log.Printf("[DeleteUser] failed to delete user: %v", err)
	}

	return c.JSON(fiber.Map{"message": "DeleteUser successfully"})
}

func (h *UsersHandler) UpdateUserInfo(c *fiber.Ctx) error {
	email := c.Params("email")
	user := new(struct {
		Role               string `json:"role"`
		OrganizationNameTH string `json:"organization_name_th"`
	})
	if err := c.BodyParser(user); err != nil {
		log.Printf("[UpdateInfo] failed to parse body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}
	f, err := h.facultyService.CheckFacultyExist(user.OrganizationNameTH)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	u := &entities.Users{
		Role:               user.Role,
		OrganizationNameTH: f.NameTH,
	}

	err = h.usersService.UpdateInfo(email, u)
	if err != nil {
		log.Printf("[UpdateInfo] failed to update user info: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user info",
		})
	}
	return c.JSON(fiber.Map{"message": "UpdateInfo successfully"})
}

func (s *UsersHandler) GetStaff(c *fiber.Ctx) error {
	staff, err := s.usersService.GetStaff()
	if err != nil {
		log.Printf("[GetStaff] failed to get staff: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get staff",
		})
	}

	return c.JSON(staff)
}

func (s *UsersHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := s.usersService.GetAllUsers()
	if err != nil {
		log.Printf("[GetAllUser] failed to get all user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get all user",
		})
	}

	return c.JSON(users)
}
