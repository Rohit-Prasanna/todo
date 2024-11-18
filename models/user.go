package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Todo represents a single Todo item
type Todo struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID   string             `json:"userId" bson:"userId"`
	Title    string             `json:"title" bson:"title"`
	Date     string             `json:"date" bson:"date"`
	Time     string             `json:"time" bson:"time"`
	Dropdown bool               `json:"dropdown" bson:"dropdown"`
	Checked  bool               `json:"checked" bson:"checked"`
}
