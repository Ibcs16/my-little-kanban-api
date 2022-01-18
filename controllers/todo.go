package controllers

import (
	"backend/my-little-kanvas/configs"
	"backend/my-little-kanvas/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var todoCollection *mongo.Collection = configs.GetCollection(configs.DB, "todos")
var listsCollection *mongo.Collection = configs.GetCollection(configs.DB, "lists")
var validate = validator.New()

// structured functions?
func GetTodos() gin.HandlerFunc {
    return func(c *gin.Context) {
		// get context
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var todos []models.Todo
		defer cancel()
		
		// find all todos
		results, err := todoCollection.Find(ctx, bson.M{})
		
		// handle error if given
		if err != nil {
			// if nothing found, returns JSON object with message
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"InternalError"})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var todo models.Todo
			if err = results.Decode(&todo); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"InternalError"})
			}
			todos = append(todos, todo)
		}
		// if found, returns JSON object of todo
		c.IndentedJSON(http.StatusOK, todos)
	}
}

func EditTodo() gin.HandlerFunc {
    return func(c *gin.Context) {
		// get context
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		var todo models.Todo
		defer cancel()
		
		// get object id from param string
		oid, _ := primitive.ObjectIDFromHex(id)

		//validate the request body
        if err := c.BindJSON(&todo); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"message":"Bad request"})
            return
        }

		// get updated object
		update := bson.M{"title": todo.Title, "status": todo.Status}
		// persist update
        result, err := todoCollection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": update})

		if err != nil  {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Internal server error"})
			return
		}

        if result.MatchedCount < 1  {
            c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Not found"})
            return
		}

		todo.Id = oid;
        

		// if updated, returns JSON object of todo
		c.IndentedJSON(http.StatusOK, todo)
	}
}

// structured functions?
func GetLists() gin.HandlerFunc {
    return func(c *gin.Context) {
		// get context
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var lists []models.TodoList
		defer cancel()
		
		// find all todos
		results, err := listsCollection.Find(ctx, bson.M{})
		
		// handle error if given
		if err != nil {
			// if nothing found, returns JSON object with message
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"InternalError"})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var list models.TodoList
			if err = results.Decode(&list); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
				return;
			}
			lists = append(lists, list)
		}
		// if found, returns JSON object of todo
		c.IndentedJSON(http.StatusOK, lists)
	}
}

func EditListCardIds() gin.HandlerFunc {
    return func(c *gin.Context) {
		// get context
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		var list models.TodoList
		defer cancel()
		
		// get object id from param string
		oid, _ := primitive.ObjectIDFromHex(id)

		//validate the request body
        if err := c.BindJSON(&list); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"message":"Bad request"})
            return
        }

		// get updated object
		update := bson.M{"cardIds": list.CardIds}
		// persist update
        result, err := listsCollection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": update})

		if err != nil  {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Internal server error", "error": err.Error()})
			return
		}

        if result.MatchedCount < 1  {
            c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Not found"})
            return
		}

		var updatedUser models.TodoList
        if result.MatchedCount == 1 {
            err := listsCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&updatedUser)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error", "error": err.Error()})
                return
            }
        }
        
		// if updated, returns JSON object of todo
		c.IndentedJSON(http.StatusOK, updatedUser)
	}
}