package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func validateToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, nil)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// Proceed to the next middleware or route handler
	c.Next()
}

// TodoItem struct to represent a todo list item
type TodoItem struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      string             `json:"userID"`
	Description string             `json:"description" binding:"required"`
	Completed   bool               `json:"completed"`
	PhotoURL    string             `json:"photoURL"`
	CreatedAt   string             `json:"createdAt"`
}

func todoRoutes(r *gin.Engine) {
	todoGroup := r.Group("/api/todos").Use(validateToken)
	{
		todoGroup.POST("", addTodoHandler)
		todoGroup.GET("", getTodosHandler)
		todoGroup.PUT("/:id", updateTodoHandler)
		todoGroup.DELETE("/:id", deleteTodoHandler)
	}
}

func addTodoHandler(c *gin.Context) {
	// Parse user ID from JWT token
	tokenString := c.GetHeader("Authorization")
	token, _ := jwt.Parse(tokenString, nil)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["username"].(string)

	var newTodo TodoItem
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo.UserID = userID
	newTodo.CreatedAt = "your-formatted-timestamp"

	// Insert the new todo item into the database
	result, err := todoCollection.InsertOne(context.TODO(), newTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"insertedID": result.InsertedID})
}

func getTodosHandler(c *gin.Context) {
	// Parse user ID from JWT token
	tokenString := c.GetHeader("Authorization")
	token, _ := jwt.Parse(tokenString, nil)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["username"].(string)

	// Retrieve todo items for the user from the database
	cursor, err := todoCollection.Find(context.TODO(), bson.M{"userID": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer cursor.Close(context.TODO())

	var todos []TodoItem
	err = cursor.All(context.TODO(), &todos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func updateTodoHandler(c *gin.Context) {
	// Parse user ID from JWT token
	tokenString := c.GetHeader("Authorization")
	token, _ := jwt.Parse(tokenString, nil)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["username"].(string)

	// Parse todo item ID from request parameters
	todoID := c.Param("id")

	// Convert todoID string to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	var updatedTodo TodoItem
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the todo item in the database
	filter := bson.M{"_id": objectID, "userID": userID}
	update := bson.M{"$set": updatedTodo}

	_, err = todoCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusOK)
}

func deleteTodoHandler(c *gin.Context) {
	// Parse user ID from JWT token
	tokenString := c.GetHeader("Authorization")
	token, _ := jwt.Parse(tokenString, nil)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["username"].(string)

	// Parse todo item ID from request parameters
	todoID := c.Param("id")

	// Convert todoID string to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	// Delete the todo item from the database
	_, err = todoCollection.DeleteOne(context.TODO(), bson.M{"_id": objectID, "userID": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusOK)
}
