package main

import (
	"context"
	"log"
	"os"

	"finara-api/config"
	"finara-api/handlers"
	"finara-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Firebase Admin SDK
	ctx := context.Background()
	firebaseApp, err := config.InitializeFirebase(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to initialize Firebase: %v", err)
	}
	log.Println("✅ Firebase Admin SDK initialized successfully")

	// Initialize Firestore client
	firestoreClient, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to initialize Firestore client: %v", err)
	}
	defer firestoreClient.Close()
	log.Println("✅ Firestore client initialized successfully")

	// Initialize handlers with Firestore client
	userHandler := handlers.NewUserHandler(firestoreClient)

	// Setup Gin router
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Setup routes
	router.SetupRoutes(r, userHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Starting Geni Firestore API server on port %s", port)
	log.Printf("📋 Available endpoints:")
	log.Printf("   GET /users - Get all users")
	log.Printf("   GET /users/:userId - Get specific user")
	log.Printf("   POST /users/:userId - Register/Update user")
	log.Printf("   GET /users/:userId/goal_info - Get user's goals")
	log.Printf("   GET /users/:userId/goal_info/:goalId - Get specific goal")
	log.Printf("   GET /health - Health check")

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
