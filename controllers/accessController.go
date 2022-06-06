package controllers

import (
	"github.com/akadendry/go-jdsport-raffle/database"
	"github.com/akadendry/go-jdsport-raffle/models"
	"github.com/gofiber/fiber/v2"
)

func AllAccess(c *fiber.Ctx) error {
	var access []models.Access

	database.DB.Find(&access)

	return c.JSON(access)
}
