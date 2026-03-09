package routes


import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)



func UserRoutes(router *gin.Engine) {

	user:=router.Group("/user")

	user.POST("/register", controllers.CreateUser)
	user.POST("/login", controllers.LoginUser)

}