package routes

import (
	"backend/my-little-kanvas/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoute(router *gin.Engine)  {
	router.GET("/todos", controllers.GetTodos())
	router.GET("/todos/:id", controllers.GetTodoById())
	router.PUT("/todos/:id", controllers.EditTodo())
	router.GET("/lists", controllers.GetLists())
	router.GET("/lists/:id", controllers.GetListById())
	router.PATCH("/lists/:id", controllers.EditListCardIds())
}