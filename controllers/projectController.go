package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
	"github.com/lamhoangvu217/task-management-be-golang/services"
	"github.com/lamhoangvu217/task-management-be-golang/utils"
	"time"
)

type AddCollaboratorRequest struct {
	ProjectID uint `json:"project_id" validate:"required"`
	UserID    uint `json:"user_id" validate:"required"`
}

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

func GetProjectByUserId(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(uint)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	projects, err := services.GetProjectByUserId(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":  "get projects successfully",
		"projects": projects,
		"userId":   userId,
	})
}

func AddCollaboratorToProject(c *fiber.Ctx) error {
	var req AddCollaboratorRequest
	// Parse the JSON body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	// Find the project
	var project = models.Project{}
	if err := database.DB.First(&project, req.ProjectID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
	}
	// Find the label
	var user models.User
	if err := database.DB.First(&user, req.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	// Associate the label with the task
	if err := database.DB.Model(&project).Association("Collaborators").Append(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User added to project successfully"})
}
