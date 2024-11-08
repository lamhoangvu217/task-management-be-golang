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
	RoleID    uint `json:"role_id" validate:"required"`
}

type UpdateCollaboratorRequest struct {
	UserID    uint `json:"user_id" validate:"required"`
	ProjectID uint `json:"project_id" validate:"required"`
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
	project.OwnerID = userId

	createdProject, err := services.CreateProject(project)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "project created successfully",
		"project": createdProject,
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
			"error":  err.Error(),
			"userId": userId,
		})
	}

	return c.JSON(fiber.Map{
		"message":  "get projects successfully",
		"projects": projects,
		"userId":   userId,
	})
}

func AddCollaboratorToProject(c *fiber.Ctx) error {
	var req models.UserProjectRole
	// Parse the JSON body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	// Find the project
	var project = models.Project{}
	if err := database.DB.First(&project, req.ProjectID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
	}
	// Find the user
	var user *models.User
	if err := database.DB.First(&user, req.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Find role
	var role models.Role
	if err := database.DB.First(&role, req.RoleID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}
	res, err := services.AddCollaboratorToProject(&req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User added to project successfully", "user": res})
}

func UpdateCollaboratorFromProject(c *fiber.Ctx) error {
	var bodyReq UpdateCollaboratorRequest
	if err := c.BodyParser(&bodyReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	if err := services.UpdateCollaboratorInProject(bodyReq.ProjectID, bodyReq.UserID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to remove collaborator",
		})
	}
	return c.JSON(fiber.Map{
		"message": "remove collaborator successfully",
	})
}
