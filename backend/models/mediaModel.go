package models


import "time"

type MediaModel struct {
	ID        string    `json:"id" bson:"_id"`
	Filename  string    `json:"filename" bson:"filename"`
	URL       string    `json:"url" bson:"url"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
