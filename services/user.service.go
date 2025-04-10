package services

import (
	"example/web-service-gin/config"
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new user
func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING user_id"
    err := config.DB.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.Id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, user)
}