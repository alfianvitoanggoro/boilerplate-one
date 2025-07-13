package rbac

import (
	"boilerplate-one/internal/domain/user"

	"github.com/gofiber/fiber/v2"
)

// MiddlewareRBAC memeriksa apakah user memiliki permission tertentu
func MiddlewareRBAC(permissionID uint) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u, ok := c.Locals("user").(*user.User)
		if !ok || u == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		rbacSvc := &RBACService{}
		if !rbacSvc.HasPermission(u.RoleID, permissionID) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: insufficient permission"})
		}
		return c.Next()
	}
}
