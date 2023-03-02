package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nkolentcev/paxcontrol/paxprod/database"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":8000")
}
