package middleware

import (
	"strings"

	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		// fmt.Printf("In [middleware] : tokenStr = %v\n", tokenStr)
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		// Attach claims to context for later use
		c.Locals("name", claims.Name)
		c.Locals("email", claims.Email)
		c.Locals("prename_id", claims.PreNameID)
		c.Locals("prename_th", claims.PreNameTH)
		c.Locals("prename_en", claims.PreNameEN)
		c.Locals("firstname_th", claims.FirstNameTH)
		c.Locals("firstname_en", claims.FirstNameEN)
		c.Locals("lastname_th", claims.LastNameTH)
		c.Locals("lastname_en", claims.LastNameEN)
		c.Locals("organizationcode", claims.OrganizationCode)
		c.Locals("organizationname_th", claims.OrganizationNameTH)
		c.Locals("organizationname_en", claims.OrganizationNameEN)
		c.Locals("itaccounttype_id", claims.ITAccountTypeID)
		c.Locals("itaccounttype_th", claims.ITAccountTypeTH)
		c.Locals("itaccounttype_en", claims.ITAccountTypeEN)
		return c.Next()
	}
}

func RequireRole(role services.RequireRoleService, allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		emailInJWT := c.Locals("email")
		if emailInJWT == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		email := emailInJWT.(string)
		// fmt.Println("==> Allowed roles:", allowedRoles, " | Role in JWT:", roleInJWT)
		role, err := role.GetRoleByEmail(email)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "role not found"})
		}

		if slices.Contains(allowedRoles, role) {
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}
}
