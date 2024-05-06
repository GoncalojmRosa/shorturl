package model

type Site struct {
	Id       string    `json:"id" bson:"_id"`
	Url      string    `json:"url" bson:"url"`
	ShortUrl string    `json:"shortUrl" bson:"shortUrl"`
	Clicks   int       `json:"clicks" bson:"clicks"`
	// createAt time.Time `json:"createAt" bson:"createAt"`
}