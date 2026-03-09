package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {

	task := router.Group("/task")

	task.POST("/create",middleware.AdminAuthMiddleware(), controllers.CreateTask)

	task.GET("/getAllTask", controllers.GetAllTasks)

	//task.PATCH("/complete/:taskId", controllers.CompleteTask)

	// task.GET("/user/:id/tasks",controllers.GetUserTasks)


	

	//task.GET("/task:id")

}