package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/task-management-be-golang/controllers"
	"github.com/lamhoangvu217/task-management-be-golang/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/user-detail", middlewares.AuthRequired, controllers.GetUserDetail)

	authorizedApp := app.Group("/app", middlewares.AuthRequired)

	authorizedApp.Post("/project", controllers.CreateProject)
	authorizedApp.Get("/projects", controllers.GetProjectByUserId)
	authorizedApp.Post("/add-collaborator", controllers.AddCollaboratorToProject)

	authorizedApp.Get("/tasks", controllers.GetTasksByProject)
	authorizedApp.Post("/task", controllers.CreateTask)
	authorizedApp.Delete("/task/:id", controllers.DeleteTask)
	authorizedApp.Put("/task/:id", controllers.UpdateTask)

	authorizedApp.Get("/subtasks", controllers.GetSubtaskByTask)
	authorizedApp.Post("/subtask", controllers.CreateSubtask)
	authorizedApp.Delete("/subtask/:id", controllers.DeleteSubtask)
	authorizedApp.Put("/subtask/:id", controllers.UpdateSubtask)

	authorizedApp.Post("/label", controllers.CreateLabel)
	authorizedApp.Get("/labels", controllers.GetAllLabels)
	authorizedApp.Post("/assign-label", controllers.AssignLabelToTask)
	authorizedApp.Post("/remove-assign-label", controllers.RemoveLabelFromTask)
	authorizedApp.Delete("/label/:id", controllers.DeleteLabel)

	authorizedApp.Post("/comment", controllers.CreateComment)
	authorizedApp.Get("/comments", controllers.GetCommentByUser)
	authorizedApp.Delete("/comment/:id", controllers.DeleteComment)
}
