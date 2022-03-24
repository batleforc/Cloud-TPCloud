package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Label       string
	Description string
	Status      bool
}
