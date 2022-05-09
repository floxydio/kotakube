package handler

import (
	"fmt"
	database "kotaku/Database"
	models "kotaku/Models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllVoucher(c *fiber.Ctx) error {
	var voucherModel []models.Voucher

	database.DB.Find(&voucherModel)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  200,
		"data":    voucherModel,
		"message": "Successfully Load Data Voucher",
	})
}

func VoucherCreate(c *fiber.Ctx) error {
	voucherModel := models.Voucher{}

	if err := c.BodyParser(&voucherModel); err != nil {
		fmt.Println("ERROR IN BODYPARSER")
	}

	database.DB.Create(&voucherModel)

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"data":    voucherModel,
		"message": "Successfully Created Voucher",
	})
}

func VoucherDecrement(c *fiber.Ctx) error {
	var voucherModel models.Voucher
	id := c.Params("id")

	database.DB.Model(voucherModel).Where("id = ?", id).Update("voucher_quantity", gorm.Expr("voucher_quantity - ?", 1))

	return c.JSON(fiber.Map{
		"message": "Decrement",
	})
}
