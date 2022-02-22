package handler

import (
	"demo/app/service"
	"demo/pkg/validator"
	"demo/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUserById(ctx *fiber.Ctx) error {
	id := uuid.MustParse(ctx.Params("id"))
	user, err := service.GetUserById(id)

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(user)
}

func GetUserByEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")
	user, err := service.GetUserByEmail(email)

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(user)
}

func GetAllUsers(ctx *fiber.Ctx) error {
	users, err := service.GetAllUsers()

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(users)
}

func Adduser(ctx *fiber.Ctx) error {
	var body postgres.InsertUserParams

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	errors := validator.ValidateStruct(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	user, err := service.InsertUser(body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	var body postgres.UpdateUserParams

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	errors := validator.ValidateStruct(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	user, err := service.UpdateUser(body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := service.DeleteUser(uuid.MustParse(id))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(fiber.Map{
		"success": true,
	})
}
