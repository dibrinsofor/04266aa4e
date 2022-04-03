package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Find out how bson arrays work and use that instead
type Playlist struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Urls     []string           `json:"urls,omitempty" bson:"urls,omitempty"`
	RandSlug string             `json:"rand_slug,omitempty" bson:"rand_slug,omitempty"`
}

// func FindPlaylistBySlug(slug string) ([]*Playlist, error) {
// 	filter := bson.D{
// 		primitive.E{Key: "rand_slug", Value: slug},
// 	}
// 	return filterTasks(filter)

// }

// TODO find out how to filter and also implement jwt
