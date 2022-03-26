package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Label  string             `json:"label"`
	Status bool               `json:"status"`
}

func (t *Task) GetTaskById(Dbhandler Db, Client *mongo.Client, id primitive.ObjectID) error {
	return GetTaskColl(Dbhandler, Client).FindOne(context.TODO(), bson.M{"_id": id}).Decode(t)
}

func GetTaskColl(Database Db, client *mongo.Client) *mongo.Collection {
	return Database.GetDb(client).Collection("Task")
}

func CreateNextTask(Label string) *Task {
	return &Task{
		Label:  Label,
		Status: false,
	}
}
