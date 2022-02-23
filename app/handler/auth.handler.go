package handler

import (
	"demo/app/service"
	"demo/postgres"
	"github.com/gofiber/fiber/v2"
)

func Profile(ctx *fiber.Ctx) error {
	return ctx.JSON("Profile")
}

func Register(ctx *fiber.Ctx) error {
	var body postgres.InsertUserParams

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.JSON("Registration form")
}

func Login(ctx *fiber.Ctx) error {
	var body struct {
		Email, Password string
	}

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	token, err := service.Login(body.Email, body.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
