package services

import (
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
)

func GetSubtaskByTask(taskId uint) ([]models.Subtask, error) {
	var subtasks []models.Subtask
	if err := database.DB.Where("task_id = ?", taskId).Preload("Task").Find(&subtasks).Error; err != nil {
		return nil, err
	}
	return subtasks, nil
}

func CreateSubtask(subtask *models.Subtask) (*models.Subtask, error) {
	if err := database.DB.Create(&subtask).Error; err != nil {
		return nil, err
	}
	return subtask, nil
}

func DeleteSubtask(subtask *models.Subtask, subtaskId uint) error {
	if err := database.DB.Delete(&subtask, subtaskId).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSubtask(updatedSubtask *models.Subtask) error {
	if err := database.DB.Save(&updatedSubtask).Error; err != nil {
		return err
	}
	return nil
}
