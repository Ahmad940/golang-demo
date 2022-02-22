package route

import (
	"demo/app/handler"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {
	user := app.Group("/user")
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetUserById)
	user.Get("/:email", handler.GetUserByEmail)
	user.Patch("/:", handler.UpdateUser)
	user.Delete("/", handler.DeleteUser)
}
