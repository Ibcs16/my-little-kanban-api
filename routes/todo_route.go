package routes

import (
	"backend/my-little-kanvas/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoute(router *gin.Engine)  {
	router.GET("/todos", controllers.GetTodos())
	router.PUT("/todos/:id", controllers.EditTodo())
	router.GET("/lists", controllers.GetLists())
	router.PUT("/lists/:id", controllers.EditListCardIds())
}