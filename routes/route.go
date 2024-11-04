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
	authorizedApp.Get("/tasks", controllers.GetTasksByUserId)
	authorizedApp.Post("/task", controllers.CreateTask)
	authorizedApp.Delete("/task/:id", controllers.DeleteTask)
	authorizedApp.Put("/task/:id", controllers.UpdateTask)
}
