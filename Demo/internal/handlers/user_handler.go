package handlers

import (
	"Demo/internal/models"
	"Demo/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (uh *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := uh.userRepo.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(users)
}

func (uh *UserHandler) GetUserByID(c *fiber.Ctx) error {
	str := c.Params("id")
	num, _ := strconv.ParseUint(str, 10, 64)
	id := uint(num)
	user, err := uh.userRepo.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	return c.JSON(user)
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := uh.userRepo.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	str := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	num, _ := strconv.ParseUint(str, 10, 64)
	id := uint(num)
	user.ID = id
	if err := uh.userRepo.UpdateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {
	str := c.Params("id")
	num, _ := strconv.ParseUint(str, 10, 64)
	id := uint(num)
	if err := uh.userRepo.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusNoContent).SendString("User deleted successfully")
}
