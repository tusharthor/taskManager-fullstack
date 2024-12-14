package controllers

import (
	"net/http"
	"taskmanager/config"
	"taskmanager/models"
	"taskmanager/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input for login"})
		return
	}

	query := "select id from users where email=? and password=?"
	var userID int
	err := config.DB.QueryRow(query, user.Email, user.Password).Scan(&userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid cred for login"})
		return
	}

	token, err := utils.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate login token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}
