package controllers

import (
	"github.com/akadendry/go-jdsport-raffle/v2/database"
	"github.com/akadendry/go-jdsport-raffle/v2/models"
	"github.com/gofiber/fiber/v2"
)

func AllAccess(c *fiber.Ctx) error {
	var access []models.Access

	database.DB.Find(&access)

	return c.JSON(access)
}
