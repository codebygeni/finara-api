package main

import (
	"context"
	"log"
	"os"
	"time"

	"finara-api/config"
	"finara-api/handlers"
	"finara-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Firebase Admin SDK
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
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
	r.Use(gin.Recovery())

	// Setup routes
	router.SetupRoutes(r, userHandler)


	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("🚀 Starting Geni Firestore API server on port %s", port)
	log.Printf("📋 Available endpoints:")
	log.Printf("   GET /users - Get all users")
	log.Printf("   GET /users/:userId - Get specific user")
	log.Printf("   GET /users/:userId/goal_info - Get user's goals")
	log.Printf("   GET /users/:userId/goal_info/:goalId - Get specific goal")
	log.Printf("   GET /health - Health check")

	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
