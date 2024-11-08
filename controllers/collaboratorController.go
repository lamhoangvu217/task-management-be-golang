package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/task-management-be-golang/services"
	"strconv"
)

func GetCollaboratorsByProjectId(c *fiber.Ctx) error {
	// Retrieve projectId from the query parameters
	projectIdStr := c.Query("projectId")
	if projectIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "project id is required",
		})
	}

	// Convert projectId string to uint
	projectId, err := strconv.ParseUint(projectIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid project id",
		})
	}

	// Call service to get collaborators by project ID
	collaborators, err := services.GetCollaboratorsByProjectId(uint(projectId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch collaborators",
		})
	}

	// Return collaborators with user info
	return c.JSON(fiber.Map{
		"collaborators": collaborators,
	})
}
