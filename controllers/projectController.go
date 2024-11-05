package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/task-management-be-golang/models"
	"github.com/lamhoangvu217/task-management-be-golang/services"
	"github.com/lamhoangvu217/task-management-be-golang/utils"
	"time"
)

func CreateProject(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(uint)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	project := new(models.Project)
	if err := c.BodyParser(project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	if !utils.IsValidProjectStatus(project.Status) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "project status is invalid",
		})
	}
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	project.UserID = userId

	createdProject, err := services.CreateProject(project)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "project created successfully",
		"task":    createdProject,
	})
}
