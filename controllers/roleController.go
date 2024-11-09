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

func GetAllRoles(c *fiber.Ctx) error {
	roles, err := services.GetAllRoles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "get all roles successfully",
		"roles":   roles,
	})
}

func CreateRole(c *fiber.Ctx) error {
	role := new(models.Role)
	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	createdRole, err := services.CreateRole(role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "role created successfully",
		"role":    createdRole,
	})
}

func DeleteRole(c *fiber.Ctx) error {
	roleIdStr := c.Params("id")
	if roleIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "role id is required",
		})
	}
	// Convert task id from string to uint
	roleId, err := strconv.ParseUint(roleIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid label id",
		})
	}
	var role models.Role
	if err := database.DB.First(&role, roleId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "role id not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve role",
		})
	}

	if err := services.DeleteRole(&role, uint(roleId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete role",
		})
	}
	return c.JSON(fiber.Map{
		"message": "delete role successfully",
	})
}
