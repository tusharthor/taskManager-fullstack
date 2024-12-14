package controllers

import (
	"net/http"
	"taskmanager/config"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	userID := c.Param("userID")
	rows, err := config.DB.Query("select id, title, description, completed, user_id from tasks where user_id=?", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch tasks"})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		//single task
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.UserID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error scanning tasks"})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	//if correct
	_, err := config.DB.Exec("insert into tasks (title, description, completed) values (?,?,?)",
		task.Title, task.Description, task.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "task created successfully"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := config.DB.Exec("delete from tasks where id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
