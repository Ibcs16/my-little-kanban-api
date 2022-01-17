package controllers

import (
	"backend/my-little-kanvas/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoController struct {
	session *mgo.Session
}

func GetController(session *mgo.Session) *TodoController{
	return &TodoController{session}
}

// structured functions
func (todoController TodoController) GetTodos (c *gin.Context) {

}

func (todoController TodoController) CreateTodo (c *gin.Context) {

}

func (todoController TodoController) GetTodo (c *gin.Context) {
	// assign the id param of the request to a variable
	id := c.Param("id")

	// loops through todos array and check for id
	if !bson.IsObjectIdHex(id) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	oid := bson.ObjectIdHex(id)

	// create empty object of type Todo
	todo := models.Todo{}

	// find todo by id and bind it to todo model variable
	// handle error if given
	if err := todoController.session.DB("my-little-kanban").C("todos").FindId(oid).One(&todo); err != nil {
		// if nothing found, returns JSON object with message
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"})
		return
	}

	// if found, returns JSON object of todo
	c.IndentedJSON(http.StatusOK, todo)
}

func (todoController TodoController) DeleteTodo (c *gin.Context) {

}




// func getTodos(c *gin.Context) {
// 	// gets context from gin, like Req/Res from express
// 	// returns JSON response with status and object
// 	c.IndentedJSON(http.StatusOK, todos)
// }

// func postTodo(c *gin.Context) {
// 	// defines new variable of type todo
// 	var newTodo Todo

// 	// binds the body of the request to newTodo variable
// 	// as it might throw error, check if binding returns error before moving on
// 	if err := c.BindJSON((&newTodo)); err != nil {
// 		return
// 	}

// 	// reassign variable with new object added to it
// 	todos = append(todos, newTodo)
// 	// returns JSON response
// 	c.IndentedJSON(http.StatusCreated, newTodo)
// }

// func getTodoByID(c *gin.Context) {
	
// }