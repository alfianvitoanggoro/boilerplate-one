package user

import (
	"boilerplate-one/internal/domain/user/dto"
	"boilerplate-one/pkg/response"
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service Service
}

func NewUserHandler(s Service) *UserHandler {
	return &UserHandler{service: s}
}

// ...existing code...

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Invalid request body", err)
	}
	user, err := h.service.Register(context.Background(), &req)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Could not create user", err)
	}
	return response.SendCreated(c, "User created successfully", user)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Invalid user ID", err)
	}
	user, err := h.service.GetUserByID(context.Background(), uint(id))
	if err != nil {
		return response.SendError(c, fiber.StatusNotFound, "User not found", err)
	}
	return response.SendSuccess(c, "User found successfully", user)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUser(context.Background())
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Could not list users", err)
	}
	return response.SendSuccess(c, "Users listed successfully", users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Invalid user ID", err)
	}
	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Invalid request body", err)
	}
	user, err := h.service.UpdateUser(context.Background(), uint(id), &req)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Could not update user", err)
	}
	return response.SendSuccess(c, "User updated successfully", user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Invalid user ID", err)
	}
	err = h.service.DeleteUser(context.Background(), uint(id))
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Could not delete user", err)
	}
	return response.SendSuccess(c, "User deleted successfully", nil)
}
