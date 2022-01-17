package controllers

import (
	"backend/my-little-kanvas/models"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type TodoController struct {
	collection *mongo.Collection
	ctx context.Context
}

func GetController(todoCollection *mongo.Collection, ctx context.Context) *TodoController{
	return &TodoController{collection: todoCollection, ctx: ctx}
}

// structured functions
func (tc TodoController) GetTodos (c *gin.Context) {
	// find all todos
	var todos []models.Todo

	// handle error if given
	cursor, err := tc.collection.Find(tc.ctx, bson.M{})
	
	if err != nil {
		log.Fatal(err)
		// if nothing found, returns JSON object with message
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"InternalError"})
		return
	}

	for cursor.Next(tc.ctx) {
		var todo models.Todo
		if err = cursor.Decode(&todo); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	// if found, returns JSON object of todo
	c.IndentedJSON(http.StatusOK, todos)
	defer cursor.Close(tc.ctx)
}

// func (tc TodoController) CreateTodo (c *gin.Context) {
// 	// defines new variable of type todo
// 	var newTodo models.Todo

// 	// binds the body of the request to newTodo variable
// 	// as it might throw error, check if binding returns error before moving on
// 	if err := c.BindJSON((&newTodo)); err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Internal error"})
// 		return
// 	}

// 	// get new _id from mongo
// 	newTodo.Id = bson.NewObjectId()

// 	// try to insert it to db
// 	if err := tc.session.DB("my-little-kanban").C("todos").Insert(newTodo); err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
// 		return
// 	}
	
// 	// returns JSON response
// 	c.IndentedJSON(http.StatusCreated, newTodo)
// }

// func (tc TodoController) GetTodo (c *gin.Context) {
// 	// assign the id param of the request to a variable
// 	id := c.Param("id")

// 	// loops through todos array and check for id
// 	if !bson.IsObjectIdHex(id) {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
// 		return
// 	}

// 	// convert string hex to object id type
// 	oid := bson.ObjectIdHex(id)

// 	// create empty object of type Todo
// 	todo := models.Todo{}

// 	// find todo by id and bind it to todo model variable
// 	// handle error if given
// 	if err := tc.session.DB("my-little-kanban").C("todos").FindId(oid).One(&todo); err != nil {
// 		// if nothing found, returns JSON object with message
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"})
// 		return
// 	}

// 	// if found, returns JSON object of todo
// 	c.IndentedJSON(http.StatusOK, todo)
// }

// func (tc TodoController) DeleteTodo (c *gin.Context) {
// 	// assign the id param of the request to a variable
// 	id := c.Param("id")

// 	// loops through todos array and check for id
// 	if !bson.IsObjectIdHex(id) {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
// 		return
// 	}

// 	// convert string hex to object id type
// 	oid := bson.ObjectIdHex(id)

// 	if err := tc.session.DB("my-little-kanban").C("todos").RemoveId(oid); err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted", "_id": id})
// }


