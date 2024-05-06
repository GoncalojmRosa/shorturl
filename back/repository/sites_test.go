package repository

import (
	"context"
	"log"
	"testing"

	"github.com/goncalojmrosa/shorturl/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoConnection() *mongo.Client {
	
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(""))

	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to MongoDB")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	return mongoTestClient
}

func testeMongoOperations(t *testing.T){
	mongoTestClient := newMongoConnection()
	defer mongoTestClient.Disconnect(context.Background())

	site1 := uuid.New().String()

	coll := mongoTestClient.Database("Cluster0").Collection("sites")

	empRepo := SiteRepo{MongoCollection: coll}

	t.Run("Create site 1", func(t *testing.T) {
		site := model.Site{
			Id: 	  site1,
			Url:      "https://www.google.com",
			ShortUrl: "http://localhost:8080/" + site1,
			Clicks:   0,
		}

		result, err := empRepo.CreateSite(&site)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
	})
}