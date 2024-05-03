package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func goDotEnvVariable(key string) string {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func Connect() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	url := goDotEnvVariable("DATABASE_URL")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	return client.Database("Cluster0")
}