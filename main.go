package main

import (
	"go_project/internal/database"
	"go_project/internal/models"
	"go_project/internal/redis"
	"go_project/routes"
	"log"
)

func main() {
	db, err := database.NewPostgresClient()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err := db.GormDB().AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	redis := redis.NewRedisClient()

	r := routes.RegisterRoutes(db, redis)
	r.Run(":8080")
}
