package route

import (
	"github.com/gofiber/fiber/v2"
)

func RoutersInitializer(app *fiber.App) {
	api := app.Group("/api")
	AuthRouter(api)
	UserRouter(api)
	TodoRouter(api)
}
