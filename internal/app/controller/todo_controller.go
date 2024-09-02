package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/plutonska/todolist-go/internal/app/models/domain"
	"github.com/plutonska/todolist-go/internal/app/service"
)

type TodoController struct {
	TodoService service.TodoService
}

func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	var todo domain.Todo
	err := ctx.BodyParser(&todo)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = c.TodoService.CreateTodo(ctx.Context(), &todo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&todo)
}

func (c *TodoController) GetTodoByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	todo, err := c.TodoService.GetTodoByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(todo)
}

func (c *TodoController) GetAllTodos(ctx *fiber.Ctx) error {
	todos, err := c.TodoService.GetAllTodos(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(todos)
}

func (c *TodoController) UpdateTodo(ctx *fiber.Ctx) error {
	var todo domain.Todo
	err := ctx.BodyParser(&todo)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	todo.UUID = ctx.Params("id")

	err = c.TodoService.UpdateTodo(ctx.Context(), &todo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(todo)
}

func (c *TodoController) DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.TodoService.DeleteTodo(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
