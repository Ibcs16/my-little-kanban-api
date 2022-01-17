// Define main as a module
package main

// Import needed libraries
import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// Define object typing
type todo struct {
	// the string tells how the attribute will be called when serialized to JSON
	ID string `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

// Create new variable to store todos
var todos = []todo{
	{
		ID: "1", Title: "Test", Status: "todo",
	},
}

// var clientOptions = options.Client().ApplyURI("mongodb+srv://iagobrayham:iagobrayham@cluster0.h2jsn.mongodb.net/my-little-kanban?retryWrites=true&w=majority")
// var client = mongo.Client()

// The start function of the module
func main() {
	// I guess this extracts the router object from gin.Default()
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	router := gin.Default()

	// Define endpoints and handlers/callbacks
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", postTodo)

	// Tell the server to run on port 8080
	router.Run("localhost:8080")
}

func getTodos(c *gin.Context) {
	// gets context from gin, like Req/Res from express
	// returns JSON response with status and object
	c.IndentedJSON(http.StatusOK, todos)
}

func postTodo(c *gin.Context) {
	// defines new variable of type todo
	var newTodo todo

	// binds the body of the request to newTodo variable
	// as it might throw error, check if binding returns error before moving on
	if err := c.BindJSON((&newTodo)); err != nil {
		return
	}

	// reassign variable with new object added to it
	todos = append(todos, newTodo)
	// returns JSON response
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(c *gin.Context) {
	// assign the id param of the request to a variable
	id := c.Param("id")

	// loops through todos array and check for id
	for _, a := range todos {
		if a.ID == id {
			// if found, returns JSON
			c.IndentedJSON(http.StatusOK, a)
			// returns to stop function execution
			return
		}
	}
	// if nothing found, returns JSON object with message
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}