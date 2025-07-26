package router

import (
	"finara-api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	// Health check endpoint
	r.GET("/health", userHandler.HealthCheck)

	// Simple user routes
	r.GET("/users", userHandler.GetAllUsers)                               // GET /users
	r.GET("/users/:userId", userHandler.GetUserByID)                       // GET /users/:userId
	r.POST("/users/:userId", userHandler.RegisterUser)                     // POST /users/:userId
	r.GET("/users/:userId/goal_info", userHandler.GetUserGoals)            // GET /users/:userId/goal_info
	r.GET("/users/:userId/goal_info/:goalId", userHandler.GetSpecificGoal) // GET /users/:userId/goal_info/:goalId

	// Root endpoint with API info
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service":     "Geni Firestore API",
			"version":     "v1.0.0",
			"description": "Simple REST API for retrieving data from Firebase Firestore",
			"endpoints": map[string]string{
				"health":        "GET /health",
				"users":         "GET /users",
				"user_by_id":    "GET /users/:userId",
				"register_user": "POST /users/:userId",
				"user_goals":    "GET /users/:userId/goal_info",
				"specific_goal": "GET /users/:userId/goal_info/:goalId",
			},
		})
	})
}
