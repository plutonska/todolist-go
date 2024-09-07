package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/plutonska/todolist-go/internal/app/controller"
	"github.com/plutonska/todolist-go/internal/app/repository"
	"github.com/plutonska/todolist-go/internal/app/service"
	"github.com/plutonska/todolist-go/internal/pkg/databases"
	_ "github.com/plutonska/todolist-go/internal/pkg/databases"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	var todoRepo repository.TodoRepository
	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case "mysql", "sqlite", "postgres":
		todoRepo = repository.NewSQLTodoRepository(databases.DB)
	case "redis":
		todoRepo = repository.NewRedisTodoRepository(databases.RDB)
	default:
		log.Fatalf("Unsupported DB_TYPE: %v", dbType)
	}

	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	app.Post("/v1/todos", todoController.CreateTodo)
	app.Get("/v1/todos", todoController.GetAllTodos)
	app.Get("/v1/todos/:uuid", todoController.GetTodoByID)
	app.Put("/v1/todos/:uuid", todoController.UpdateTodo)
	app.Delete("/v1/todos/:uuid", todoController.DeleteTodo)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
