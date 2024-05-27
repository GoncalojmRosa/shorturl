package types

import "time"

type Config struct {
	PublicHost string
	Port       string
	Database_URI string
}

type SiteStore interface {
	FindByShortUrl(shortUrl string) (*Site, error)
	FindAll() ([]*Site, error)
	Insert(site *Site) error
	Update(site *Site) error
	Delete(site *Site) error
}

type RegisterSitePayload struct {
	Url string `json:"url"`
}

type Site struct {
	Id       string    `json:"id" bson:"_id"`
	Url      string    `json:"url" bson:"url"`
	ShortUrl string    `json:"shortUrl" bson:"shortUrl"`
	BaseUrl  string    `json:"baseUrl" bson:"baseUrl"`
	Clicks   int       `json:"clicks" bson:"clicks"`
	CreateAt time.Time `json:"createAt" bson:"createAt"`
}