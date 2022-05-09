package handler

import (
	database "kotaku/Database"
	model "kotaku/Models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var userModel model.User

	if err := c.BodyParser(&userModel); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid register",
		})
	}
	database.DB.Create(&userModel)

	return c.JSON(fiber.Map{
		"message": "Successfully created",
	})
}

func Login(c *fiber.Ctx) error {
	var userModel model.User

	if err := c.BodyParser(&userModel); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid login",
		})
	}

	var user model.User
	database.DB.Where("email = ? AND password = ?", userModel.Email, userModel.Password).First(&user)

	if user.Email == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid login",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully login",
		"data":    user,
	})
}
