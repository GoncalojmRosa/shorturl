package sites

import (
	"github.com/goncalojmrosa/shorturl/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db *mongo.Client
}

func NewStore(db *mongo.Client) *Store {
	return &Store{db: db}
}

// Delete implements types.SiteStore.
func (s *Store) Delete(site *types.Site) error {
	panic("unimplemented")
}

// FindAll implements types.SiteStore.
func (s *Store) FindAll() ([]*types.Site, error) {
	panic("unimplemented")
}

// FindByShortUrl implements types.SiteStore.
func (s *Store) FindByShortUrl(shortUrl string) (*types.Site, error) {
	panic("unimplemented")
}

// Insert implements types.SiteStore.
func (s *Store) Insert(site *types.Site) error {
	panic("unimplemented")
}

// Update implements types.SiteStore.
func (s *Store) Update(site *types.Site) error {
	panic("unimplemented")
}
