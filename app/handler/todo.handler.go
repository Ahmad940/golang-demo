package handler

import (
	"demo/app/model"
	"demo/app/service"
	"demo/pkg/validator"
	"demo/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllTodos(ctx *fiber.Ctx) error {
	todos, err := service.GetTodos()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(todos)
}

func GetATodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	todo, err := service.GetATodo(uuid.MustParse(id))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(todo)
}

func InsertTodo(ctx *fiber.Ctx) error {
	var body model.Todo

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	errors := validator.ValidateStruct(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	todo, err := service.InsertTodo(body.Title)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(todo)
}

func UpdateTodo(ctx *fiber.Ctx) error {
	var body postgres.UpdateTodoParams
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	todo, err := service.UpdateTodo(body.ID, body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(todo)
}

func DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := service.DeleteTodo(uuid.MustParse(id))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(fiber.Map{
		"success": true,
	})
}
