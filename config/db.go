package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/naveenbalaji2001/employee_dashboard_app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ctx = context.Background()

func SetupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("postgres://postgres:Naveen1341@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	db.AutoMigrate(&model.Profile{})
	db.AutoMigrate(&model.Registers{})
	db.AutoMigrate(&model.Todo{})

	return db, nil
}

func SetupRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
	} else {
		fmt.Printf("Connected to Redis: %v\n", pong)
	}

	return client
}

func SetupRecentDeletedRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       1,                // Use a separate DB for recently deleted cache
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis (recent deleted): %v\n", err)
	} else {
		fmt.Printf("Connected to Redis (recent deleted): %v\n", pong)
	}

	return client
}
