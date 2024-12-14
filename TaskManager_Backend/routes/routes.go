package routes

import (
	"taskmanager/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	//enable cors (middleware) - very imp
	router.Use(cors.Default())

	//routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/tasks/:userID", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	return router
}
