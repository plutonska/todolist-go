package databases

import (
	"github.com/plutonska/todolist-go/internal/app/models/domain"
	"gorm.io/gorm"
	"log"
)

func MigrateDB(db *gorm.DB) {
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
