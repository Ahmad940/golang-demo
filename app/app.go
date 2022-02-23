package app

import (
	"demo/pkg/config"
	"demo/pkg/middleware"
	"demo/pkg/route"
	"github.com/gofiber/fiber/v2"
	"log"
)

func InitialiseApp() {
	//  App
	app := fiber.New(config.NewFiberConfig())

	// Middlewares
	middleware.FiberMiddlewares(app)

	// Routers
	route.RoutersInitializer(app)

	log.Fatal(app.Listen(":4000"))
}
