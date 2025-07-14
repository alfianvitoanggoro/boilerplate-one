package middleware

import (
	"boilerplate-one/internal/domain/auth"
	"boilerplate-one/internal/domain/user"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(userRepo user.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ParseJWT(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Ambil user dari database
		u, err := userRepo.FindByID(c.Context(), claims.UserID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		// Simpan ke context Fiber
		c.Locals("user", u)
		return c.Next()
	}
}
