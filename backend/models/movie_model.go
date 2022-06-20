package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Id    primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Title string             `json:"title, omitempty" bson:"title, omitempty"`
	Genre string             `json:"genre, omitempty" bson:"genre, omitempty"`
	Year  string             `json:"year, omitempty" bson:"year, omitempty"`
}
