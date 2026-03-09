package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)



func UserRoutes(router *gin.Engine) {

	user:=router.Group("/user")

	user.POST("/register", controllers.CreateUser)
	user.POST("/login", controllers.LoginUser)

	user.GET("/mytasks", middleware.UserAuthMiddleware() , controllers.GetMyTasks)


	user.PATCH("/:taskId",middleware.UserAuthMiddleware(), controllers.CompleteTask)


	

		// task.GET("/user/:id/tasks",controllers.GetUserTasks)

}