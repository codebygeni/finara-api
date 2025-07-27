package router

import (
	"finara-api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	// Health check endpoint
	r.GET("/health", userHandler.HealthCheck)

	// Dashboard endpoint - serves the HTML file
	r.GET("/dashboard", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.File("dashboard.html")
	})

	// User-specific dashboard endpoint
	r.GET("/dashboard/:userId", func(c *gin.Context) {
		_ = c.Param("userId") // User ID for future use
		// You can add user-specific logic here if needed
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.File("dashboard.html")
	})

	// Simple user routes
	r.GET("/users", userHandler.GetAllUsers)                               // GET /users
	r.GET("/users/:userId", userHandler.GetUserByID)                       // GET /users/:userId
	r.POST("/users/:userId", userHandler.RegisterUser)                     // POST /users/:userId
	r.GET("/users/:userId/goal_info", userHandler.GetUserGoals)            // GET /users/:userId/goal_info
	r.GET("/users/:userId/goal_info/:goalId", userHandler.GetSpecificGoal) // GET /users/:userId/goal_info/:goalId
	r.POST("/users/:userId/goal_info/:goalId", userHandler.RegisterGoal)   // POST /users/:userId/goal_info/:goalId

	// Root endpoint with API info
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service":     "Geni Firestore API",
			"version":     "v1.0.0",
			"description": "Simple REST API for retrieving data from Firebase Firestore",
			"endpoints": map[string]string{
				"health":         "GET /health",
				"dashboard":      "GET /dashboard",
				"user_dashboard": "GET /dashboard/:userId",
				"users":          "GET /users",
				"user_by_id":     "GET /users/:userId",
				"register_user":  "POST /users/:userId",
				"user_goals":     "GET /users/:userId/goal_info",
				"specific_goal":  "GET /users/:userId/goal_info/:goalId",
				"register_goal":  "POST /users/:userId/goal_info/:goalId",
			},
		})
	})
}
