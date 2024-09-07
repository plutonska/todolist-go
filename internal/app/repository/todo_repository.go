package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/plutonska/todolist-go/internal/app/models/domain"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *domain.Todo) error
	GetByUUID(ctx context.Context, uuid string) (*domain.Todo, error)
	GetAll(ctx context.Context) ([]*domain.Todo, error)
	Update(ctx context.Context, todo *domain.Todo) error
	Delete(ctx context.Context, uuid string) error
}

type SQLTodoRepository struct {
	DB *gorm.DB
}

func NewSQLTodoRepository(db *gorm.DB) TodoRepository {
	return &SQLTodoRepository{DB: db}
}

func (r *SQLTodoRepository) Create(ctx context.Context, todo *domain.Todo) error {
	todo.UUID = uuid.New().String()
	return r.DB.WithContext(ctx).Create(todo).Error
}

func (r *SQLTodoRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Todo, error) {
	var todo domain.Todo
	err := r.DB.WithContext(ctx).Where("uuid = ?", uuid).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *SQLTodoRepository) GetAll(ctx context.Context) ([]*domain.Todo, error) {
	var todos []*domain.Todo
	err := r.DB.WithContext(ctx).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *SQLTodoRepository) Update(ctx context.Context, todo *domain.Todo) error {
	existingTodo := &domain.Todo{}
	err := r.DB.WithContext(ctx).Where("uuid = ?", todo.UUID).First(existingTodo).Error
	if err != nil {
		return err
	}

	updateFields := map[string]interface{}{
		"updated_at": time.Now(),
		"item":       todo.Item,
		"done":       todo.Done,
	}
	return r.DB.WithContext(ctx).Model(existingTodo).Updates(updateFields).Error
}

func (r *SQLTodoRepository) Delete(ctx context.Context, uuid string) error {
	return r.DB.WithContext(ctx).Where("uuid = ?", uuid).Delete(&domain.Todo{}).Error
}

type RedisTodoRepository struct {
	Client *redis.Client
}

func NewRedisTodoRepository(client *redis.Client) TodoRepository {
	return &RedisTodoRepository{Client: client}
}

func (r *RedisTodoRepository) Create(ctx context.Context, todo *domain.Todo) error {
	todo.UUID = uuid.New().String()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("todo:%s", todo.UUID)
	return r.Client.Set(ctx, key, jsonData, 0).Err()
}

func (r *RedisTodoRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Todo, error) {
	key := fmt.Sprintf("todo:%s", uuid)
	data, err := r.Client.Get(ctx, key).Bytes()
	if err != nil {

		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("todo not found")
		}
		return nil, err
	}

	var todo domain.Todo
	err = json.Unmarshal(data, &todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *RedisTodoRepository) GetAll(ctx context.Context) ([]*domain.Todo, error) {
	keys, err := r.Client.Keys(ctx, "todo:*").Result()
	if err != nil {
		return nil, err
	}

	var todos []*domain.Todo
	for _, key := range keys {
		data, err := r.Client.Get(ctx, key).Bytes()
		if err != nil {
			continue
		}

		var todo domain.Todo
		err = json.Unmarshal(data, &todo)
		if err != nil {
			continue
		}

		todos = append(todos, &todo)
	}
	return todos, nil

}

func (r *RedisTodoRepository) Update(ctx context.Context, todo *domain.Todo) error {
	key := fmt.Sprintf("todo:%s", todo.UUID)
	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	var existingTodo domain.Todo
	err = json.Unmarshal([]byte(value), &existingTodo)
	if err != nil {
		return err
	}

	todo.CreatedAt = existingTodo.CreatedAt

	todo.UpdatedAt = time.Now()
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return r.Client.Set(ctx, key, jsonData, 0).Err()
}

func (r *RedisTodoRepository) Delete(ctx context.Context, uuid string) error {
	key := fmt.Sprintf("todo:%s", uuid)
	return r.Client.Del(ctx, key).Err()
}
