package sites

import (
	"context"
	"fmt"

	"github.com/goncalojmrosa/shorturl/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db *mongo.Client
}

const (
	sitesDatabase   = "shorturl"
	sitesCollection = "sites"
)

func NewStore(db *mongo.Client) *Store {
	return &Store{db: db}
}

// Delete implements types.SiteStore.
func (s *Store) Delete(site *types.Site) error {
	panic("unimplemented")
}

// FindAll implements types.SiteStore.
func (s *Store) FindAll(ctx context.Context) (*[]types.Site, error) {
	col := s.db.Database(sitesDatabase).Collection(sitesCollection)
	fmt.Println(col)

	filter := bson.M{}

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []types.Site
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

// FindByUrlCode implements types.SiteStore.
func (s *Store) FindByUrlCode(ctx context.Context, urlCode string) (*types.Site, error) {
	col := s.db.Database(sitesDatabase).Collection(sitesCollection)

	var site types.Site
	filter := bson.D{{Key: "shortUrl", Value: urlCode}}

	if err := col.FindOne(ctx, filter).Decode(&site); err != nil {
		return nil, err
	}

	return &site, nil
}

// Insert implements types.SiteStore.
func (s *Store) Insert(ctx context.Context, site *types.Site) (*types.Site, error) {
	col := s.db.Database(sitesDatabase).Collection(sitesCollection)
	_, err := col.InsertOne(ctx, site)

	if err != nil {
		return nil, err
	}

	return site, nil
}

// Update implements types.SiteStore.
func (s *Store) Update(site *types.Site) error {
	panic("unimplemented")
}
