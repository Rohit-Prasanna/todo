package models

import (
	"time"
)

// Todo represents a task in the to-do application
type Todo struct {
	ID        string    `bson:"_id,omitempty" json:"id"` // Changed to string
	Title     string    `bson:"title" json:"title"`
	Date      string    `bson:"date" json:"date"`
	Time      string    `bson:"time" json:"time"`
	Checked   bool      `bson:"checked" json:"checked"`
	Dropdown  bool      `bson:"dropdown" json:"dropdown"`
	UserID    string    `bson:"userId" json:"userId"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
