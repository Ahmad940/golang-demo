package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddlewares(app *fiber.App) {
	// middlewares
	app.Use(cors.New())
	app.Use(recover.New())

	// http middlewares
	app.Get("/monitor", monitor.New())
}
