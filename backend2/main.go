package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"backend2/config"
	"backend2/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	// MongoDB connection
	mongoDBURL := config.GetMongoDBURL()
	log.Printf("Attempting to connect to MongoDB at: %s", mongoDBURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB successfully")

	// Router setup
	router := mux.NewRouter()

	// Setup routes
	routes.SetupRoutes(router, client)

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Add your frontend URL
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Use the CORS middleware
	handler := c.Handler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	log.Printf("Server is starting on port %s", port)
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}