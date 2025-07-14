package middleware

import (
	"boilerplate-one/internal/domain/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterMiddleware(app *fiber.App) {
	app.Use(logger.New())
}

// RBACMiddleware returns a Fiber handler that checks if the user's role is allowed
func RBACMiddleware(allowedRoles ...uint) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(*user.User)
		if !ok || user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		for _, role := range allowedRoles {
			if user.RoleID == role {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: insufficient role"})
	}
}
