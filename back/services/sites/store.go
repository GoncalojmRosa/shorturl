package sites

import (
	"context"
	"fmt"

	"github.com/goncalojmrosa/shorturl/types"
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
func (s *Store) FindAll(ctx context.Context) ([]*types.Site, error) {
	col := s.db.Database(sitesDatabase).Collection(sitesCollection)
	fmt.Println(col)
	var sites types.Site
	cursor, err := col.Find(ctx, &sites)
	if err != nil {
		return nil, err
	}

	fmt.Println(cursor)

	var results []*types.Site
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

// FindByShortUrl implements types.SiteStore.
func (s *Store) FindByShortUrl(shortUrl string) (*types.Site, error) {
	panic("unimplemented")
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
