package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// InitializeFirebase initializes the Firebase Admin SDK
func InitializeFirebase(ctx context.Context) (*firebase.App, error) {
	// Path to your Firebase service account key file
	serviceAccountKeyPath := "config/firebase-service-account.json"

	// Initialize Firebase Admin SDK
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	config := &firebase.Config{
		ProjectID: "geni-project",
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Printf("❌ Error initializing Firebase app: %v", err)
		return nil, err
	}

	return app, nil
}
