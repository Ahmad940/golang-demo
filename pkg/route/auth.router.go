package route

import (
	"demo/app/handler"
	"demo/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router) {
	auth := app.Group("/auth")
	auth.Get("/profile", middleware.JWTProtected(), handler.Profile)
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}
