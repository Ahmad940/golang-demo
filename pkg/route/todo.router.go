package route

import (
	"demo/app/handler"
	"github.com/gofiber/fiber/v2"
)

func TodoRouter(app fiber.Router) {
	todo := app.Group("/todo")
	todo.Get("/", handler.GetAllTodos)
	todo.Get("/:id", handler.GetATodo)
	todo.Post("/", handler.InsertTodo)
	todo.Patch("/", handler.UpdateTodo)
	todo.Delete("/:id", handler.DeleteTodo)
}
