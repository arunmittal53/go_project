package routes

import (
	"go_project/handlers"
	"go_project/internal/database"
	"go_project/internal/redis"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(db *database.PostgresClient, redis *redis.RedisClient) *gin.Engine {
	router := gin.Default()

	handlers := handlers.NewUserHandler(db, redis)

	router.POST("/users", handlers.CreateUser)
	router.GET("/users", handlers.GetUsers)
	router.GET("/user/:id", handlers.GetUser)
	router.DELETE("/user/:id", handlers.DeleteUser)

	router.GET("/redis/:id", handlers.GetRedisValue)
	router.DELETE("/redis/:id", handlers.DeleteRedisKey)

	return router
}
