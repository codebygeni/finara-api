package main

import (
	"context"
	"log"
	"os"
	"time"

	"finara-api/config"
	"finara-api/handlers"
	"finara-api/router"

	"github.com/gin-contrib/cors"
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

	// Configure CORS middleware for React compatibility
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000", // React development server
			"http://localhost:5173", // Vite development server
			"http://127.0.0.1:3000",
			"http://127.0.0.1:5173",
			"https://your-react-app.vercel.app", // Add your production domain
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Content-Length", "Accept-Encoding",
			"X-CSRF-Token", "Authorization", "Accept", "Cache-Control",
			"X-Requested-With",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	router.SetupRoutes(r, userHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Starting Geni Firestore API server on port %s", port)
	log.Printf("🌐 CORS enabled for React development (localhost:3000, localhost:5173)")
	log.Printf("📋 Available endpoints:")
	log.Printf("   GET /dashboard - Financial dashboard")
	log.Printf("   GET /dashboard/:userId - User-specific dashboard")
	log.Printf("   GET /users - Get all users")
	log.Printf("   GET /users/:userId - Get specific user")
	log.Printf("   POST /users/:userId - Register/Update user")
	log.Printf("   GET /users/:userId/goal_info - Get user's goals")
	log.Printf("   GET /users/:userId/goal_info/:goalId - Get specific goal")
	log.Printf("   POST /users/:userId/goal_info/:goalId - Register/Update goal")
	log.Printf("   GET /health - Health check")

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
