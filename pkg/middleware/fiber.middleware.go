package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddlewares(app *fiber.App) {
	// middleware
	app.Use(cors.New())
	app.Use(recover.New())

	// http middleware
	app.Get("/monitor", monitor.New())
}
