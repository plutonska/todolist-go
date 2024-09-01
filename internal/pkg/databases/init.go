package databases

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB
var RDB *redis.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB_NAME"),
		)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed connect to MySQL: %v", err)
		}

	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(os.Getenv("SQLITE_FILE")), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed connect to SQLite: %v", err)
		}

	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"), os.Getenv("POSTGRES_PORT"),
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL default database: %v", err)
		}

	case "redis":
		RDB = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})

		_, err := RDB.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("Failed connect to Redis: %v", err)
		}

	default:
		log.Fatalf("Unsupported DB_TYPE: %s", dbType)
	}

	log.Printf("Connected to %s database successfully", dbType)

	if os.Getenv("APP_ENV") == "development" {
		MigrateDB(DB, RDB, dbType)
	}

}
