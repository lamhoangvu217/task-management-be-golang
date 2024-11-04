package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
	"github.com/lamhoangvu217/task-management-be-golang/services"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func CreateComment(c *fiber.Ctx) error {
	comment := new(models.Comment)
	userId := c.Locals("userId").(uint)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	comment.UserID = userId
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	createdComment, err := services.CreateComment(comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"comment": comment,
		})
	}
	return c.JSON(fiber.Map{
		"message": "comment created successfully",
		"comment": createdComment,
	})
}

func GetCommentByUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	// Call the service to get tasks by user
	comments, err := services.GetCommentByUser(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message":  "get comments successfully",
		"comments": comments,
		"userId":   userId,
	})
}

func DeleteComment(c *fiber.Ctx) error {
	commentIdStr := c.Params("id")
	if commentIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "comment id is required",
		})
	}
	// Convert task id from string to uint
	commentId, err := strconv.ParseUint(commentIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid comment id",
		})
	}
	var comment models.Comment
	if err := database.DB.First(&comment, commentId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "comment id not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve comment",
		})
	}

	if err := services.DeleteComment(&comment, uint(commentId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete comment",
		})
	}
	return c.JSON(fiber.Map{
		"message": "delete comment successfully",
	})
}
