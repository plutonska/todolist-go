package service

import (
	"context"
	"github.com/plutonska/todolist-go/internal/app/models/domain"
	"github.com/plutonska/todolist-go/internal/app/repository"
)

type TodoService interface {
	CreateTodo(ctx context.Context, todo *domain.Todo) error
	GetTodoByUUID(ctx context.Context, uuid string) (*domain.Todo, error)
	GetAllTodos(ctx context.Context) ([]*domain.Todo, error)
	UpdateTodo(ctx context.Context, todo *domain.Todo) error
	DeleteTodo(ctx context.Context, uuid string) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	return s.repo.Create(ctx, todo)
}

func (s *todoService) GetTodoByUUID(ctx context.Context, uuid string) (*domain.Todo, error) {
	return s.repo.GetByUUID(ctx, uuid)
}

func (s *todoService) GetAllTodos(ctx context.Context) ([]*domain.Todo, error) {
	return s.repo.GetAll(ctx)
}

func (s *todoService) UpdateTodo(ctx context.Context, todo *domain.Todo) error {
	return s.repo.Update(ctx, todo)
}

func (s *todoService) DeleteTodo(ctx context.Context, uuid string) error {
	return s.repo.Delete(ctx, uuid)
}
