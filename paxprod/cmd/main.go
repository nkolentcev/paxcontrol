package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nkolentcev/paxcontrol/paxprod/database"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	setupRoutes(app)

	app.Listen(":8000")
}
