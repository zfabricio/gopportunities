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
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

// LoadEnvVariables loads environment variables from the .env file
func LoadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error loading .env file: %w", err)
	}
	return nil
}

// InitializeMongoDB establishes a connection to MongoDB and assigns it to a global variable.
func InitializeMongoDB() error {
	// Load environment variables
	if err := LoadEnvVariables(); err != nil {
		return fmt.Errorf("Error loading environment variables: %w", err)
	}

	// Get MongoDB credentials from environment variables
	dns := os.Getenv("MONGO_DNS")

	if dns == "" {
		return fmt.Errorf("MongoDB connection details are missing in environment variables")
	}

	// Construct the MongoDB connection URI
	uri := fmt.Sprintf(dns)

	// MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Establish the MongoDB connection with a timeout
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %w", err)
	}

	// Assign the client to the global variable
	mongoClient = client

	// Ping the database to verify the connection
	if err := pingMongoDB(); err != nil {
		return fmt.Errorf("Error pinging MongoDB: %w", err)
	}

	log.Println("Successfully connected to MongoDB Atlas")

	return nil
}

// pingMongoDB verifies if MongoDB is reachable by pinging the server.
func pingMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping the MongoDB server to ensure the connection is established
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return nil
}

// CloseMongoConnection gracefully closes the MongoDB connection.
func CloseMongoConnection() error {
	if mongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Disconnect from MongoDB
		if err := mongoClient.Disconnect(ctx); err != nil {
			return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
		}
		log.Println("MongoDB connection closed successfully")
	}

	return nil
}

// GetMongoClient - Returns the MongoDB client
func GetMongoClient() *mongo.Client {
	return mongoClient
}
