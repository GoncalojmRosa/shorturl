package repository

import (
	"context"

	"github.com/goncalojmrosa/shorturl/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SiteRepo struct {
	MongoCollection *mongo.Collection
}

func (s *SiteRepo) CreateSite(site *model.Site) (interface{}, error) {
	result, err := s.MongoCollection.InsertOne(context.Background(), site)
	if err != nil {
		return nil, err
	}
	return result, nil
}

 func (s *SiteRepo) GetAllSites() ([]model.Site, error) {
 	cursor, err := s.MongoCollection.Find(context.Background(), bson.D{})
 	if err != nil {
 		return nil, err
 	}
 	defer cursor.Close(context.Background())

 	var sites []model.Site
 	if err = cursor.All(context.Background(), &sites); err != nil {
 		return nil, err
 	}
 	return sites, nil
 }