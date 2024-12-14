package controllers

import (
	"log"
	"net/http"
	"taskmanager/config"
	"taskmanager/models"
	"taskmanager/utils"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "insert into users (username, email, password) values (?,?,?)"
	res, err := config.DB.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		log.Fatalf("error inserting user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	userID, _ := res.LastInsertId()
	token, err := utils.GenerateToken(int(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to generate token"})
		return
	}

	//output token for user
	c.JSON(http.StatusOK, gin.H{"message": "user created succss", "token": token})

}
