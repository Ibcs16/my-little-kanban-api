// Define main as a module
package main

// Import needed libraries
import (
	"backend/my-little-kanvas/controllers"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// The start function of the module
func main() {
	// setup mongo connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://iagobrayham:iagobrayham@cluster0.4deg4.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	err  = client.Ping(ctx, readpref.Primary())
	
	if err != nil {
		log.Fatal(err)
	}
	
	database := client.Database("my-little-kanban")
	todosCollection := database.Collection("todos")

	// I guess this extracts the router object from gin.Default()
	router := gin.Default()
	todoController := controllers.GetController(todosCollection, ctx)

	// Define endpoints and handlers/callbacks
	router.GET("/todos", todoController.GetTodos)
	// router.POST("/todos", todoController.CreateTodo)
	// router.GET("/todos/:id", todoController.GetTodo)
	// router.DELETE("/todos/:id", todoController.DeleteTodo)

	// Tell the server to run on port 8080
	router.Run("localhost:8080")
	defer client.Disconnect(ctx)
	
}

// function to return mongo session
// // astherisks tels the return type
// func getSession() *mgo.Session {
// 	// stablish connection to cursor
// 	// session, err := mgo.Dial("mongodb://iagobrayham:iagobrayham@cluster0.4deg4.mongodb.net/my-little-kanban")
// 	// // // session, err := mgo.Dial("mongodb+srv://iagobrayham:iagobrayham@cluster0.4deg4.mongodb.net/my-little-kanban")
// 	// // // check for error
// 	// // if err != nil {
// 	// // 	panic(err)
// 	// // }
// 	// return mgo.Session{}
// }
