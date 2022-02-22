package app

import (
	"demo/pkg/config"
	"demo/pkg/middlewares"
	"demo/pkg/route"
	"github.com/gofiber/fiber/v2"
	"log"
)

func InitialiseApp() {
	//  App
	app := fiber.New(config.NewFiberConfig())

	// Middlewares
	middlewares.FiberMiddlewares(app)

	// Routers
	route.RoutersInitializer(app)

	log.Fatal(app.Listen(":4000"))
}
