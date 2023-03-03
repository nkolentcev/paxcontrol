package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nkolentcev/paxcontrol/paxprod/database"
	"github.com/nkolentcev/paxcontrol/paxprod/models"
)

func CreateBoardinPass(c *fiber.Ctx) error {
	bp := new(models.BoardingPass)
	if err := c.BodyParser(bp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// TODO проверка заполнения
	if err := database.DB.DB.Create(&bp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(bp)

}
