package databases

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/plutonska/todolist-go/internal/app/models/domain"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"time"
)

func MigrateDB(db *gorm.DB, rdb *redis.Client, dbType string) {
	switch dbType {
	case "mysql", "sqlite", "postgres":
		migrateSQL(db)
	case "redis":
		migrateRedis(rdb)
	}
}

func migrateSQL(db *gorm.DB) {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}

	models := []any{
		&domain.Todo{},
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Failed to migrate model %T: %v", model, err)
		}
	}

	log.Println("Database migration completed successfully")
}

func migrateRedis(rdb *redis.Client) {
	if rdb == nil {
		log.Fatal("Redis connection is not initialized")
	}

	ctx := context.Background()
	now := time.Now()
	todos := []domain.Todos{
		{
			UUID:      uuid.New().String(),
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
			Item:      "Learn Go",
			Done:      true,
		},
		{
			UUID:      uuid.New().String(),
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
			Item:      "Read a book",
			Done:      false,
		},
		{
			UUID:      uuid.New().String(),
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
			Item:      "Go to Groceries",
			Done:      false,
		},
	}

	for _, todo := range todos {
		key := fmt.Sprintf("todo:%s", todo.UUID)
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			log.Fatalf("Failed to marshal todo: %v", err)
		}

		err = rdb.Set(ctx, key, todoJSON, 0).Err()
		if err != nil {
			log.Fatalf("Failed set todo in Redis: %v", err)
		}
	}

	log.Println("Redis migration completed successfully")
}
