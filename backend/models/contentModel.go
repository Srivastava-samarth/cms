package models

import "time"

type ContentModel struct {
	ID         string    `json:"id" bson:"_id"`
	AuthorID   string    `json:"author_id" bson:"author_id"`
	Title      string    `json:"title" bson:"title"`
	Body       string    `json:"body" bson:"body"`
	Categories []string  `json:"categories" bson:"categories"`
	Tags       []string  `json:"tags" bson:"tags"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `jaon:"updated_at" bson:"updated_at"`
}
