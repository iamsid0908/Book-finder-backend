package firebase

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// InitializeFirebase initializes Firebase using an environment variable
func InitializeFirebase() (*firebase.App, error) {
	base64Config := os.Getenv("FIREBASE_CONFIG_BASE64")
	if base64Config == "" {
		log.Fatal("FIREBASE_CONFIG_BASE64 environment variable is missing")
	}

	// Decode the Base64 string
	jsonData, err := base64.StdEncoding.DecodeString(base64Config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Firebase JSON: %w", err)
	}

	// Check if JSON is valid
	var check map[string]interface{}
	if err := json.Unmarshal(jsonData, &check); err != nil {
		return nil, fmt.Errorf("invalid Firebase JSON: %w", err)
	}

	// Initialize Firebase with credentials from decoded JSON
	opt := option.WithCredentialsJSON(jsonData)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase: %w", err)
	}

	return app, nil
}
