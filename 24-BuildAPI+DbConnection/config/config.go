package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// InitializeMongoClient establishes a MongoDB connection
func InitializeMongoClient() (*mongo.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connStr := os.Getenv("MONGO_URI")
	if connStr == "" {
		return nil, fmt.Errorf("MONGO_URI environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	MongoClient = client
	return client, nil
}

// GetMongoClient provides access to the MongoDB client
func GetMongoClient() *mongo.Client {
	return MongoClient
}
