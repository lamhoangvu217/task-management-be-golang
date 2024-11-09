package services

import (
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
)

func CreateComment(comment *models.Comment) (*models.Comment, error) {
	if err := database.DB.Create(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func GetCommentByUser(userId uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := database.DB.Where("user_id = ?", userId).Preload("User").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func DeleteComment(comment *models.Comment, commentId uint) error {
	if err := database.DB.Delete(&comment, commentId).Error; err != nil {
		return err
	}
	return nil
}

func UpdateComment(updatedComment *models.Comment) error {
	if err := database.DB.Save(&updatedComment).Error; err != nil {
		return err
	}
	return nil
}

func GetAllCommentInTask(taskId uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := database.DB.Where("task_id = ?", taskId).Preload("Task").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
