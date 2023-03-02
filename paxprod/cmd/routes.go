package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nkolentcev/paxcontrol/paxprod/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
