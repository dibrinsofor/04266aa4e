package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// {
// 	"_id": "",
// 	"created_at": "",
// 	"urls": {
// 		"link1": "",
// 		"link2": "",
// 	}
// }
// hmm chale e check like mongodb already stores update time for each task

type Playlist struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Urls     []string           `json:"urls" bson:"urls"`
	RandSlug string             `json:"rand_slug"`
}

// type Urls struct {
// 	Link1  string `json:"link1" bson:"link1"`
// 	Link2  string `json:"link2" bson:"link2"`
// 	Link3  string `json:"link3" bson:"link3"`
// 	Link4  string `json:"link4" bson:"link4"`
// 	Link5  string `json:"link5" bson:"link5"`
// 	Link6  string `json:"link6" bson:"link6"`
// 	Link7  string `json:"link7" bson:"link7"`
// 	Link8  string `json:"link8" bson:"link8"`
// 	Link9  string `json:"link9" bson:"link9"`
// 	Link10 string `json:"link10" bson:"link10"`
// }
