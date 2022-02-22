package route

import (
	"demo/app/handler"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router) {
	auth := app.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}
