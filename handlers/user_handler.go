package handlers

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	firestoreClient *firestore.Client
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(client *firestore.Client) *UserHandler {
	return &UserHandler{
		firestoreClient: client,
	}
}

// User represents the user data structure
type User struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Age               int    `json:"age"`
	Email             string `json:"email"`
	MobileNo          string `json:"mobile_no"`
	PreferredLanguage string `json:"preferred_language"`
	MaritalStatus     string `json:"marrital_status"`
	City              string `json:"city"`
	CareerStage       string `json:"career_stage"`
}

// UserRegistrationRequest represents the request body for user registration
type UserRegistrationRequest struct {
	Name              string `json:"name" binding:"required"`
	Age               int    `json:"age" binding:"required"`
	Email             string `json:"email" binding:"required"`
	MobileNo          string `json:"mobile_no" binding:"required"`
	PreferredLanguage string `json:"preferred_language" binding:"required"`
	MaritalStatus     string `json:"marrital_status" binding:"required"`
	City              string `json:"city" binding:"required"`
	CareerStage       string `json:"career_stage" binding:"required"`
}

// Goal represents the goal data structure
type Goal struct {
	ID              string  `json:"id"`
	GoalAmount      float64 `json:"goal_amount"`
	GoalDescription string  `json:"goal_description"`
	GoalLine        string  `json:"goal_line"`
	GoalTimeline    int     `json:"goal_timeline"`
}

// GetAllUsers retrieves all users from Firestore
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	ctx := context.Background()

	// Query all documents in the users collection
	iter := h.firestoreClient.Collection("users").Documents(ctx)
	var users []User

	for {
		doc, err := iter.Next()
		if err != nil {
			// Check if we've reached the end
			if status.Code(err) == codes.NotFound {
				break
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to fetch users",
				"details": err.Error(),
			})
			return
		}

		// Convert document data to User struct
		var user User
		user.ID = doc.Ref.ID

		data := doc.Data()
		if name, ok := data["name"].(string); ok {
			user.Name = name
		}
		if age, ok := data["age"].(int64); ok {
			user.Age = int(age)
		}
		if email, ok := data["email"].(string); ok {
			user.Email = email
		}
		if mobileNo, ok := data["mobile_no"].(string); ok {
			user.MobileNo = mobileNo
		}
		if preferredLanguage, ok := data["preferred_language"].(string); ok {
			user.PreferredLanguage = preferredLanguage
		}
		if maritalStatus, ok := data["marrital_status"].(string); ok {
			user.MaritalStatus = maritalStatus
		}
		if city, ok := data["city"].(string); ok {
			user.City = city
		}
		if careerStage, ok := data["career_stage"].(string); ok {
			user.CareerStage = careerStage
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No users found",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a specific user by ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("userId")
	ctx := context.Background()

	// Get the user document
	doc, err := h.firestoreClient.Collection("users").Doc(userID).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "User not found",
				"user_id": userID,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch user",
			"details": err.Error(),
		})
		return
	}

	// Convert document data to User struct
	var user User
	user.ID = doc.Ref.ID

	data := doc.Data()
	if name, ok := data["name"].(string); ok {
		user.Name = name
	}
	if age, ok := data["age"].(int64); ok {
		user.Age = int(age)
	}
	if email, ok := data["email"].(string); ok {
		user.Email = email
	}
	if mobileNo, ok := data["mobile_no"].(string); ok {
		user.MobileNo = mobileNo
	}
	if preferredLanguage, ok := data["preferred_language"].(string); ok {
		user.PreferredLanguage = preferredLanguage
	}
	if maritalStatus, ok := data["marrital_status"].(string); ok {
		user.MaritalStatus = maritalStatus
	}
	if city, ok := data["city"].(string); ok {
		user.City = city
	}
	if careerStage, ok := data["career_stage"].(string); ok {
		user.CareerStage = careerStage
	}

	c.JSON(http.StatusOK, user)
}

// GetUserGoals retrieves all goals for a specific user
func (h *UserHandler) GetUserGoals(c *gin.Context) {
	userID := c.Param("userId")
	ctx := context.Background()

	// First check if user exists
	_, err := h.firestoreClient.Collection("users").Doc(userID).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "User not found",
				"user_id": userID,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to verify user",
			"details": err.Error(),
		})
		return
	}

	// Query goal_info subcollection
	iter := h.firestoreClient.Collection("users").Doc(userID).Collection("goal_info").Documents(ctx)
	var goals []Goal

	for {
		doc, err := iter.Next()
		if err != nil {
			// Check if we've reached the end
			if err.Error() == "iterator done" {
				break
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to fetch goals",
				"details": err.Error(),
			})
			return
		}

		// Convert document data to Goal struct
		var goal Goal
		goal.ID = doc.Ref.ID

		data := doc.Data()

		// Get goal fields
		if amount, ok := data["goal_amount"].(float64); ok {
			goal.GoalAmount = amount
		}
		if description, ok := data["goal_description"].(string); ok {
			goal.GoalDescription = description
		}
		if line, ok := data["goal_line"].(string); ok {
			goal.GoalLine = line
		}
		if timeline, ok := data["goal_timeline"].(int64); ok {
			goal.GoalTimeline = int(timeline)
		}

		goals = append(goals, goal)
	}

	if len(goals) == 0 {
		c.JSON(http.StatusOK, []Goal{})
		return
	}

	c.JSON(http.StatusOK, goals)
}

// GetSpecificGoal retrieves a specific goal for a user
func (h *UserHandler) GetSpecificGoal(c *gin.Context) {
	userID := c.Param("userId")
	goalID := c.Param("goalId")
	ctx := context.Background()

	// First check if user exists
	_, err := h.firestoreClient.Collection("users").Doc(userID).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "User not found",
				"user_id": userID,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to verify user",
			"details": err.Error(),
		})
		return
	}

	// Get the specific goal document
	doc, err := h.firestoreClient.Collection("users").Doc(userID).Collection("goal_info").Doc(goalID).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Goal not found",
				"user_id": userID,
				"goal_id": goalID,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch goal",
			"details": err.Error(),
		})
		return
	}

	// Convert document data to Goal struct
	var goal Goal
	goal.ID = doc.Ref.ID

	data := doc.Data()

	// Get goal fields
	if amount, ok := data["goal_amount"].(float64); ok {
		goal.GoalAmount = amount
	}
	if description, ok := data["goal_description"].(string); ok {
		goal.GoalDescription = description
	}
	if line, ok := data["goal_line"].(string); ok {
		goal.GoalLine = line
	}
	if timeline, ok := data["goal_timeline"].(int64); ok {
		goal.GoalTimeline = int(timeline)
	}

	c.JSON(http.StatusOK, goal)
}

// RegisterUser creates or updates a user in Firestore
func (h *UserHandler) RegisterUser(c *gin.Context) {
	userID := c.Param("userId")
	ctx := context.Background()

	// Parse request body
	var req UserRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Prepare user data for Firestore
	userData := map[string]interface{}{
		"name":               req.Name,
		"age":                req.Age,
		"email":              req.Email,
		"mobile_no":          req.MobileNo,
		"preferred_language": req.PreferredLanguage,
		"marrital_status":    req.MaritalStatus,
		"city":               req.City,
		"career_stage":       req.CareerStage,
	}

	// Save user data to Firestore
	_, err := h.firestoreClient.Collection("users").Doc(userID).Set(ctx, userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to save user data",
			"details": err.Error(),
		})
		return
	}

	// Return success response with the created user data
	user := User{
		ID:                userID,
		Name:              req.Name,
		Age:               req.Age,
		Email:             req.Email,
		MobileNo:          req.MobileNo,
		PreferredLanguage: req.PreferredLanguage,
		MaritalStatus:     req.MaritalStatus,
		City:              req.City,
		CareerStage:       req.CareerStage,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// HealthCheck provides a simple health check endpoint
func (h *UserHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
