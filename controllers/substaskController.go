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
	"time"
)

func GetSubtaskByTask(c *fiber.Ctx) error {
	// Extract categoryId from query parameters
	taskIdStr := c.Query("taskId")
	if taskIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "taskIdStr is required",
		})
	}
	// Convert categoryId from string to uint
	taskId, err := strconv.ParseUint(taskIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid taskId",
		})
	}
	// Call the service to get products by category
	subtasks, err := services.GetSubtaskByTask(uint(taskId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message":  "get subtasks successfully",
		"subtasks": subtasks,
	})
}

func CreateSubtask(c *fiber.Ctx) error {
	subtask := new(models.Subtask)
	if err := c.BodyParser(subtask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	subtask.CreatedAt = time.Now()
	createdSubtask, err := services.CreateSubtask(subtask)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "subtask created successfully",
		"subtask": createdSubtask,
	})
}

func DeleteSubtask(c *fiber.Ctx) error {
	subtaskIdStr := c.Params("id")
	if subtaskIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "subtask id is required",
		})
	}
	// Convert task id from string to uint
	subtaskId, err := strconv.ParseUint(subtaskIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid subtask id",
		})
	}
	var subtask models.Subtask
	if err := database.DB.First(&subtask, subtaskId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "subtask id not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve subtask",
		})
	}

	if err := services.DeleteSubtask(&subtask, uint(subtaskId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete subtask",
		})
	}
	return c.JSON(fiber.Map{
		"message": "delete subtask successfully",
	})
}

func UpdateSubtask(c *fiber.Ctx) error {
	subtaskIdStr := c.Params("id")
	if subtaskIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "subtask id is required",
		})
	}
	// Convert product id from string to uint
	subtaskId, err := strconv.ParseUint(subtaskIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid subtask id",
		})
	}
	var subtask models.Subtask
	if err := database.DB.First(&subtask, subtaskId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "subtask id not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve subtask",
		})
	}
	var updateSubtaskData models.Subtask
	if err := c.BodyParser(&updateSubtaskData); err != nil {
		// Return 400 if request body is invalid
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Update only fields that are non-zero or non-empty
	if updateSubtaskData.Title != "" {
		subtask.Title = updateSubtaskData.Title
	}
	if updateSubtaskData.Status != "" {
		subtask.Status = updateSubtaskData.Status
	}

	// Update the `UpdatedAt` field to current time
	subtask.UpdatedAt = time.Now()
	if err := services.UpdateSubtask(&subtask); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Subtask updated successfully",
		"task":    subtask,
	})
}
