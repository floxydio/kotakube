package routes

import (
	"kotaku/Handler"
	"github.com/gofiber/fiber/v2"
)

func RoutesServer(c *fiber.App) {
	c.Get("/", func(s *fiber.Ctx) error {
		return s.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	})

	// Auth [ User ]

	c.Post("/register", handler.Register)
	c.Post("/login", handler.Login)
	// -- End Auth

	// Info By Kota
	c.Get("/info", handler.GetAllDataInfo)
	c.Get("/info/:kota/:tag", handler.GetDataInfoByTag)
	// -- End Info

	// Voucher
	c.Get("/voucher", handler.GetAllVoucher)
	c.Post("/vouchers", handler.VoucherCreate)
	c.Post("/voucher/:id", handler.VoucherDecrement)

	// -- End Voucher
}
