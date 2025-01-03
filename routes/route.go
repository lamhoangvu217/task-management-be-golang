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
	app.Get("/api/plans", controllers.GetAllPlans)

	authorizedApp := app.Group("/app", middlewares.AuthRequired)

	authorizedApp.Get("/user-detail", controllers.GetUserDetail)
	authorizedApp.Put("/user", controllers.UpdateUserDetail)

	authorizedApp.Post("/project", controllers.CreateProject)
	authorizedApp.Get("/projects", controllers.GetProjectByUserId)
	authorizedApp.Post("/add-collaborator", controllers.AddCollaboratorToProject)
	authorizedApp.Get("collaborators", controllers.GetCollaboratorsByProjectId)
	authorizedApp.Put("/remove-collaborator", controllers.UpdateCollaboratorFromProject)
	authorizedApp.Get("/project/:id", controllers.GetProjectById)

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
	authorizedApp.Put("/label/:id", controllers.UpdateLabel)

	authorizedApp.Post("/comment", controllers.CreateComment)
	authorizedApp.Get("/comments-by-user", controllers.GetCommentByUser)
	authorizedApp.Delete("/comment/:id", controllers.DeleteComment)
	authorizedApp.Get("/comments-in-task", controllers.GetAllCommentInTask)
	authorizedApp.Put("/comment/:id", controllers.UpdateComment)
	authorizedApp.Get("/roles", controllers.GetAllRoles)

	authorizedApp.Post("/subscribe-plan", controllers.SubscribePlan)
	authorizedApp.Get("/current-user-plan", controllers.GetCurrentUserPlan)

	admin := app.Group("/admin", middlewares.AuthRequired)
	admin.Post("/role", controllers.CreateRole)
	admin.Delete("/role/:id", controllers.DeleteRole)
	admin.Get("/users", controllers.GetUsers)

	admin.Post("/plan", controllers.CreatePlan)
	admin.Delete("/plan/:id", controllers.DeletePlan)
	admin.Put("/plan/:id", controllers.UpdatePlan)
}
