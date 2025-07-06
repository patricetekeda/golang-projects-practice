package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive")


type NetflixMovie struct {
	ID 		primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title 	string `json:"movie,omitempty"`
	Year 	int `json:"year,omitempty"`
	Genre 	string `json:"genre,omitempty"`
	Director string `json:"director,omitempty"`
	Watched bool `json:"watched,omitempty"`
}
