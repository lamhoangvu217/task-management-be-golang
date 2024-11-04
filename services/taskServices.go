package services

import (
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
)

func GetTasksByUserId(userId uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := database.DB.Where("user_id = ?", userId).Preload("User").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func CreateTask(task *models.Task) (*models.Task, error) {
	if err := database.DB.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func DeleteTask(task *models.Task, taskId uint) error {
	if err := database.DB.Delete(&task, taskId).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTask(updatedTask *models.Task) error {
	if err := database.DB.Save(&updatedTask).Error; err != nil {
		return err
	}
	return nil
}
