package routes

import (
	"first_project/handlers"
	"first_project/internal/database"
	"first_project/internal/redis"

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
