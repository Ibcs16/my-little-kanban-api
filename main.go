// Define main as a module
package main

// Import needed libraries
import (
	"backend/my-little-kanvas/controllers"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	// "github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// The start function of the module
func main() {
	// I guess this extracts the router object from gin.Default()
	router := gin.Default()
	todoController := controllers.GetController(getSession())

	// Define endpoints and handlers/callbacks
	router.GET("/todos", todoController.GetTodos)
	router.POST("/todos", todoController.CreateTodo)
	router.GET("/todos/:id", todoController.GetTodo)
	router.DELETE("/todos/:id", todoController.DeleteTodo)

	// Tell the server to run on port 8080
	router.Run("localhost:8080")
}

//  function to return mongo session
// * astherisks tels the return type
func getSession() *mgo.Session {
	// stablish connection to cursor
	session, err := mgo.DialWithTimeout("mongodb+srv://iagobrayham:iagobrayham@cluster0.h2jsn.mongodb.net", time.Duration(5 * time.Second))
	// check for error
	if err != nil {
		panic(err)
	}
	return session
}
