package models

import "gopkg.in/mgo.v2/bson"

// Define object typing for Todo model
type Todo struct {
	// define attribute and json notation when serialized
	// also tells how bson is gonna be stored
	// looks like decorators
	Id  bson.ObjectId `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Status string `json:"status" bson:"status"`
}