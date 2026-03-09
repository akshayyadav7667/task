package controllers

import (
	"context"
	"net/http"
	"time"

	"backend/models"

	

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection

// set collection from main.go
func SetTaskCollection(collection *mongo.Collection) {
	taskCollection = collection
}

// Create Task (Admin)
func CreateTask(c *gin.Context) {

	var task models.Task

	// read request body
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := taskCollection.InsertOne(ctx, task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Task not created",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task created successfully",
		"id":      result.InsertedID,
	})
}

// Get All Tasks
func GetAllTasks(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot fetch tasks",
		})
		return
	}

	var tasks []models.Task

	if err = cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error decoding tasks",
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

