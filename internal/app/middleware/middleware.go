package middleware

import (
	"boilerplate-one/internal/domain/user"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v4"
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

// JWTMiddleware memverifikasi JWT dan menyimpan user ke context
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid Authorization header"})
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "JWT secret not set"})
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
		}
		// Ambil data user dari claims (misal: id, email, role_id)
		userObj := &user.User{
			ID:     uint(claims["id"].(float64)),
			Email:  claims["email"].(string),
			RoleID: uint(claims["role_id"].(float64)),
		}
		c.Locals("user", userObj)
		return c.Next()
	}
}
