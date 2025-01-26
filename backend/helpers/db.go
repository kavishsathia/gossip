package helpers

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDatabase() (*gorm.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		print("NOT SET")
		return nil, fmt.Errorf("DATABASE_URL not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func OpenRedis() (*redis.Client, error) {
	dbURL := os.Getenv("REDIS_URL")
	if dbURL == "" {
		print("NOT SET")
		return nil, fmt.Errorf("DATABASE_URL not set")
	}

	return redis.NewClient(&redis.Options{
		Addr: dbURL,
		DB:   0,
	}), nil
}
