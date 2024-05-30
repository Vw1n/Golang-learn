package handlers

import (
	"Demo/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (uh *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users := uh.userRepo.GetAllUsers()
	return c.JSON(users)
}
