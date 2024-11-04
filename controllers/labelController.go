package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
	"github.com/lamhoangvu217/task-management-be-golang/services"
	"gorm.io/gorm"
	"strconv"
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

func GetAllLabels(c *fiber.Ctx) error {
	labels, err := services.GetAllLabels()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "get all labels successfully",
		"labels":  labels,
	})
}

func DeleteLabel(c *fiber.Ctx) error {
	labelIdStr := c.Params("id")
	if labelIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "label id is required",
		})
	}
	// Convert task id from string to uint
	labelId, err := strconv.ParseUint(labelIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid label id",
		})
	}
	var label models.Label
	if err := database.DB.First(&label, labelId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "subtask id not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve label",
		})
	}

	if err := services.DeleteLabel(&label, uint(labelId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete label",
		})
	}
	return c.JSON(fiber.Map{
		"message": "delete label successfully",
	})
}
