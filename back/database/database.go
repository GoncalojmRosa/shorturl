package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insertSite(site string, ctx context.Context, db *mongo.Database) {
	sites := db.Collection("sites")
	sites.InsertOne(ctx, site)
}

func allSites(ctx context.Context, db *mongo.Database) {
	sites := db.Collection("sites")
	sites.Find(ctx, bson.M{})
}

func goDotEnvVariable(key string) string {
  err := godotenv.Load(".env.local")

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