package connection

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client


func GetDatabase() *mongo.Client {
	 err := godotenv.Load()
    if err != nil {
        log.Fatal("Failed to load .env file")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := os.Getenv("MONGO_URI")
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("MongoDB connection error:", err)
    }

    // Optional ping to verify connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("MongoDB ping failed:", err)
    }

	log.Println("Connected to MongoDB")
	return client
    
}


