package types

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	PublicHost   string
	Port         string
	Database_URI string
}

type SiteStore interface {
	FindByUrlCode(ctx context.Context, urlCode string) (*Site, error)
	FindAll(ctx context.Context) (*[]Site, error)
	Insert(ctx context.Context, site *Site) (*Site, error)
	Update(site *Site) error
	Delete(site *Site) error
}

type RegisterSitePayload struct {
	Url string `json:"url"`
}

type Site struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Url      string             `json:"url" bson:"url"`
	UrlCode  string             `json:"urlCode" bson:"urlCode"`
	ShortUrl string             `json:"shortUrl" bson:"shortUrl"`
	Clicks   int                `json:"clicks" bson:"clicks"`
	CreateAt time.Time          `json:"createAt" bson:"createAt"`
}
