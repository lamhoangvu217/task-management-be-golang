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

func GetTasksByUserId(c *fiber.Ctx) error {
	// Extract categoryId from query parameters
	userId := c.Locals("userId").(uint)
	// Call the service to get tasks by user
	tasks, err := services.GetTasksByUserId(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "get tasks successfully",
		"tasks":   tasks,
		"userId":  userId,
	})
}

func CreateTask(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(uint)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	task.UserID = userId
	createdTask, err := services.CreateTask(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "task created successfully",
		"task":    createdTask,
	})
}

func DeleteTask(c *fiber.Ctx) error {
	taskIdStr := c.Params("id")
	if taskIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task id is required",
		})
	}
	// Convert task id from string to uint
	taskId, err := strconv.ParseUint(taskIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id",
		})
	}
	var task models.Task
	if err := database.DB.First(&task, taskId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "task id not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve product",
		})
	}

	if err := services.DeleteTask(&task, uint(taskId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete task",
		})
	}
	return c.JSON(fiber.Map{
		"message": "delete task successfully",
	})
}

func UpdateTask(c *fiber.Ctx) error {
	taskIdStr := c.Params("id")
	if taskIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task id is required",
		})
	}
	// Convert product id from string to uint
	taskId, err := strconv.ParseUint(taskIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id",
		})
	}
	var task models.Task
	if err := database.DB.First(&task, taskId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "task id not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve task",
		})
	}
	var updateTaskData models.Task
	if err := c.BodyParser(&updateTaskData); err != nil {
		// Return 400 if request body is invalid
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if updateTaskData.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "title is required",
		})
	}
	task.Title = updateTaskData.Title
	task.Description = updateTaskData.Description
	task.Status = updateTaskData.Status
	task.Priority = updateTaskData.Priority
	task.DueDate = updateTaskData.DueDate
	task.CreatedAt = updateTaskData.CreatedAt
	task.UpdatedAt = updateTaskData.UpdatedAt
	if err := services.UpdateTask(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task updated successfully",
		"task":    task,
	})
}
