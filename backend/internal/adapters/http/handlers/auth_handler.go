package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

type signInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerReq struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	PreNameID          string `json:"prename_id"`
	PreNameTH          string `json:"prename_th"`
	PreNameEN          string `json:"prename_en"`
	FirstNameTH        string `json:"firstname_th"`
	FirstNameEN        string `json:"firstname_en"`
	LastNameTH         string `json:"lastname_th"`
	LastNameEN         string `json:"lastname_en"`
	OrganizationCode   string `json:"organization_code"`
	OrganizationNameTH string `json:"organization_name_th"`
	OrganizationNameEN string `json:"organization_name_en"`
	ITAccountTypeID    string `json:"itaccounttype_id"`
	ITAccountTypeTH    string `json:"itaccounttype_th"`
	ITAccountTypeEN    string `json:"itaccounttype_en"`
}

type SignInHandler struct {
	usersService services.UsersService
}

func NewSigninHandler(u services.UsersService, _ services.FacultyService) *SignInHandler {
	return &SignInHandler{usersService: u}
}

// SignIn godoc
// @Summary ลงชื่อเข้าใช้งานด้วย Email และ Password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body signInReq true "Email and Password"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth [post]
func (h *SignInHandler) SignIn(c *fiber.Ctx) error {
	var req signInReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email and password are required"})
	}

	user, err := h.usersService.SignIn(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	token, err := utils.GenerateJWT(
		user.Name, user.Email,
		user.PreNameID, user.PreNameTH, user.PreNameEN,
		user.FirstNameTH, user.FirstNameEN,
		user.LastNameTH, user.LastNameEN,
		user.OrganizationCode, user.OrganizationNameTH, user.OrganizationNameEN,
		user.ITAccountTypeID, user.ITAccountTypeTH, user.ITAccountTypeEN,
	)
	if err != nil {
		log.Println("Error generating JWT:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

// Register godoc
// @Summary ลงทะเบียนผู้ใช้งานใหม่
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body registerReq true "ข้อมูลการลงทะเบียน"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /api/auth/register [post]
func (h *SignInHandler) Register(c *fiber.Ctx) error {
	var req registerReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name, email, and password are required"})
	}

	profile := &entities.Users{
		Name:               req.Name,
		Email:              req.Email,
		PreNameID:          req.PreNameID,
		PreNameTH:          req.PreNameTH,
		PreNameEN:          req.PreNameEN,
		FirstNameTH:        req.FirstNameTH,
		FirstNameEN:        req.FirstNameEN,
		LastNameTH:         req.LastNameTH,
		LastNameEN:         req.LastNameEN,
		OrganizationCode:   req.OrganizationCode,
		OrganizationNameTH: req.OrganizationNameTH,
		OrganizationNameEN: req.OrganizationNameEN,
		ITAccountTypeID:    req.ITAccountTypeID,
		ITAccountTypeTH:    req.ITAccountTypeTH,
		ITAccountTypeEN:    req.ITAccountTypeEN,
	}

	user, err := h.usersService.Register(profile, req.Password)
	if err != nil {
		if err.Error() == "email already exists" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already registered"})
		}
		log.Println("Error registering user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registration successful",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
