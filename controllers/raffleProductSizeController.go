package controllers

import (
	"github.com/akadendry/go-jdsport-raffle/v1/database"
	"github.com/akadendry/go-jdsport-raffle/v1/models"
	"github.com/gofiber/fiber/v2"
)

func GetProductSizeStock(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var raffleProductSizeStock []models.RaffleProductSizeStock
	raffle_product_id := data["raffle_product_id"]
	database.DB.Where("raffle_product_id = ?", raffle_product_id).Where("deleted_at IS NULL").Find(&raffleProductSizeStock)

	return c.JSON(fiber.Map{
		"status":  "200",
		"message": "Success",
		"data":    raffleProductSizeStock,
	})
}
