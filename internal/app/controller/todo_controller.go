package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/plutonska/todolist-go/internal/app/models/domain"
	"github.com/plutonska/todolist-go/internal/app/service"
)

type TodoController struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{TodoService: todoService}
}

func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	var todo domain.Todo
	err := ctx.BodyParser(&todo)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err = c.TodoService.CreateTodo(ctx.Context(), &todo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create todo"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&todo)
}

func (c *TodoController) GetTodoByID(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	todo, err := c.TodoService.GetTodoByUUID(ctx.Context(), uuid)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	return ctx.JSON(todo)
}

func (c *TodoController) GetAllTodos(ctx *fiber.Ctx) error {
	todos, err := c.TodoService.GetAllTodos(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve todos"})
	}

	return ctx.JSON(todos)
}

func (c *TodoController) UpdateTodo(ctx *fiber.Ctx) error {
	var todo domain.Todo
	err := ctx.BodyParser(&todo)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	todo.UUID = ctx.Params("uuid")

	err = c.TodoService.UpdateTodo(ctx.Context(), &todo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update todo"})
	}

	return ctx.JSON(todo)
}

func (c *TodoController) DeleteTodo(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	err := c.TodoService.DeleteTodo(ctx.Context(), uuid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete todo"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item successfully deleted"})
}
