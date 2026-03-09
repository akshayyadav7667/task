package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"backend/models"

	//"backend/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// create the user

var userCollection *mongo.Collection

func SetUserCollection(collection *mongo.Collection) {
	userCollection = collection
}



// creating new user 
func CreateUser(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check if email already exists
	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exists",
		})
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password hashing failed",
		})
		return
	}

	user.Password = string(hashedPassword)

	// insert user
	result, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User not created",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"id":      result.InsertedID,
	})
}


// login user
func LoginUser(c *gin.Context) { // Here c = request + response (req + res together)

	var request models.User //stores login data from client
	var user models.User    //stores user data from MongoDB

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	//If MongoDB query takes more than 5 seconds → cancel it
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Find User in MongoDB
	err := userCollection.FindOne(ctx, bson.M{
		"email": request.Email,
	}).Decode(&user)

	//If MongoDB does not find the email → return error.
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email",
		})
		return
	}

	// compare password

	//request.Password → plain password
	//user.Password → hashed password from database

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)

	//If password wrong:
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong password",
		})
		return
	}

	// create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID.Hex(),
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	//Here JWT is signed using secret stored in .env.
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Token generation failed",
		})
		return
	}

	//Send Login Response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}
