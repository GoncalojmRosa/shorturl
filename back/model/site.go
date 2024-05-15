package model

import "time"

type Site struct {
	Id       string    `json:"id" bson:"_id"`
	Url      string    `json:"url" bson:"url"`
	ShortUrl string    `json:"shortUrl" bson:"shortUrl"`
	Clicks   int       `json:"clicks" bson:"clicks"`
	CreateAt time.Time `json:"createAt" bson:"createAt"`
}