package types

import (
	"context"
	"time"
)

type Config struct {
	PublicHost   string
	Port         string
	Database_URI string
}

type SiteStore interface {
	FindByShortUrl(shortUrl string) (*Site, error)
	FindAll(ctx context.Context) ([]*Site, error)
	Insert(ctx context.Context, site *Site) (Site, error)
	Update(site *Site) error
	Delete(site *Site) error
}

type RegisterSitePayload struct {
	ShortUrl string `json:"url"`
}

type Site struct {
	Id       string    `json:"id" bson:"_id"`
	Url      string    `json:"url" bson:"url"`
	ShortUrl string    `json:"shortUrl" bson:"shortUrl"`
	UrlCode  string    `json:"urlCode" bson:"urlCode"`
	Clicks   int       `json:"clicks" bson:"clicks"`
	CreateAt time.Time `json:"createAt" bson:"createAt"`
}
