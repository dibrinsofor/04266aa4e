package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Urls        []string           `json:"urls,omitempty" bson:"urls,omitempty"`
	RandSlug    string             `json:"rand_slug,omitempty" bson:"rand_slug,omitempty"`
	Description string             `json:"description" bson:"description"`
	Title       string             `json:"title" bson:"title"`
}

type IdemKey struct {
	ID         string                 `json:"IdemKey"`
	StatusCode int                    `json:"status_code"`
	Response   map[string]interface{} `json:"response"`
	CreatedAt  string                 `json:"created_at"`
}
