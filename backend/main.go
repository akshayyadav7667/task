package main

//import "github.com/gin-gonic/gin"

import (
	"log"
	"os"

	"backend/config"
	"backend/controllers"
	"backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	// connect MongoDB
	config.ConnectDB()


	// get users collection
	userCollection := config.DB.Collection("users")

	// pass collection to controller
	controllers.SetUserCollection(userCollection)


	router := gin.Default()

	//router.SetTrustedProxies(nil)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is  running",
			"port":    os.Getenv("PORT"),
		})
	})

	// load routes

	routes.UserRoutes(router)

	port := os.Getenv("PORT")

	router.Run(":" + port)
}


