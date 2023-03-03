package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nkolentcev/paxcontrol/paxprod/database"
	"github.com/nkolentcev/paxcontrol/paxprod/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("ping pong")
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	//TODO проверки заполнения

	if err := database.DB.DB.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
