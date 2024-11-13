package main

import (
	"Build_API_DB_Connection/config"
	"Build_API_DB_Connection/routes"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Initialize MongoDB connection
	_, err := config.InitializeMongoClient()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer config.GetMongoClient().Disconnect(context.Background())

	fmt.Println("Welcome to Building API with GoLang")

	// Register routes
	r := routes.RegisterRoutes()

	// Start HTTP server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
