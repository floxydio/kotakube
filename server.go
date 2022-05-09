package main

import (
	database "kotaku/Database"
	route "kotaku/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	s := fiber.New()
	database.Connect()
	s.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	s.Use(logger.New())
	// For Routes
	route.RoutesServer(s)

	s.Listen(":2000")
}
