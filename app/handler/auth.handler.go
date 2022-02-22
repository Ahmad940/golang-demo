package handler

import (
	"demo/app/service"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx *fiber.Ctx) error  {
	return ctx.JSON("Registration form")
}

func Login(ctx *fiber.Ctx) error {
	var body struct {
		email, password string
	}

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = service.Login(body.email, body.password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON("")
}
