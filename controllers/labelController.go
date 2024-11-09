package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
	"github.com/lamhoangvu217/task-management-be-golang/services"
	"gorm.io/gorm"
	"net/http"
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

func UpdateLabel(c *fiber.Ctx) error {
	labelIdStr := c.Params("id")
	if labelIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "label id is required",
		})
	}
	// Convert product id from string to uint
	labelId, err := strconv.ParseUint(labelIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid subtask id",
		})
	}
	var label models.Label
	if err := database.DB.First(&label, labelId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "label id not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve label",
		})
	}
	var updateLabelData models.Label
	if err := c.BodyParser(&updateLabelData); err != nil {
		// Return 400 if request body is invalid
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Update only fields that are non-zero or non-empty
	if updateLabelData.Name != "" {
		label.Name = updateLabelData.Name
	}
	if err := services.UpdateLabel(&label); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Label updated successfully",
		"label":   label,
	})
}
