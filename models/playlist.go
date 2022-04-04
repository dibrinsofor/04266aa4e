package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Urls     []string           `json:"urls,omitempty" bson:"urls,omitempty"`
	RandSlug string             `json:"rand_slug,omitempty" bson:"rand_slug,omitempty"`
}
