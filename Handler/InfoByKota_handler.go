package handler

import (
	"fmt"
	database "kotaku/Database"
	models "kotaku/Models"

	"github.com/gofiber/fiber/v2"
)

func GetAllDataInfo(c *fiber.Ctx) error {
	var infoKota []models.InfoKota
	database.DB.Find(&infoKota)
	return c.JSON(fiber.Map{
		"data":    infoKota,
		"message": "Successfully get all data info",
	})
}

func GetDataInfoByTag(c *fiber.Ctx) error {
	tag := c.Params("tag")
	kota := c.Params("kota")
	fmt.Println(tag + kota)
	var infoKota []models.InfoKota
	database.DB.Raw("SELECT * FROM info_kota WHERE kota = ? AND tag = ?", kota, tag).Scan(&infoKota)
	return c.JSON(fiber.Map{
		"data":    infoKota,
		"message": "Successfully get data info by tag",
	})
}
