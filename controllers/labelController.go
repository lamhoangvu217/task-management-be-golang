package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/task-management-be-golang/models"
	"github.com/lamhoangvu217/task-management-be-golang/services"
)

func CreateLabel(c *fiber.Ctx) error {
	label := new(models.Label)
	if err := c.BodyParser(label); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	createdLabel, err := services.CreateLabel(label)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "label created successfully",
		"label":   createdLabel,
	})
}
