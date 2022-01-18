package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Define object typing for Todo model
type Todo struct {
	// define attribute and json notation when serialized
	// also tells how bson is gonna be stored
	// looks like decorators
	Id  primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title string `json:"title,omitempty" validate:"required"`
	Status string `json:"status,omitempty" validate:"required"`
}

type TodoList struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title string `json:"title,omitempty" validate:"required"`
	StatusName string `json:"statusName,omitempty" validate:"required"`
	CardIds []primitive.ObjectID `json:"cardIds"`
	Order string `json:"order,omitempty"`
}