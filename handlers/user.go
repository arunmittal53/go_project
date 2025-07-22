package handlers

import (
	"fmt"
	"go_project/internal/database"
	"go_project/internal/models"
	"go_project/internal/redis"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	DB    *database.PostgresClient
	Redis *redis.RedisClient
}

func NewUserHandler(db *database.PostgresClient, redis *redis.RedisClient) *UserHandler {
	return &UserHandler{
		DB:    db,
		Redis: redis,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest models.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save to redis
	h.Redis.SaveKey(c, userRequest.Id, userRequest.Name, 5*time.Minute)

	// save to DB
	user := models.User{
		ID:   userRequest.Id,
		Name: userRequest.Name,
	}

	msg, err := h.DB.SaveUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": msg, "user": user})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.DB.FetchAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.DB.FetchUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	// delete fro Redis
	err := h.Redis.Remove(c, userID)
	if err != nil {
		fmt.Printf("Error while deleting value from redis %s, err: %+v\n", userID, err)
	} else {
		fmt.Printf("Successfully deleted from redis %s\n", userID)
	}
	// delete from DB
	user, err := h.DB.DeleteUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, user)
}

/*
Redis
*/

func (h *UserHandler) GetRedisValue(c *gin.Context) {
	userID := c.Param("id")

	value, err := h.Redis.Get(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, value)
}

func (h *UserHandler) DeleteRedisKey(c *gin.Context) {
	userID := c.Param("id")

	err := h.Redis.Remove(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Deleted Redis key")
}
