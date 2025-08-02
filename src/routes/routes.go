package routes

import (
	"github.com/fromcode/to-do-app/src/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.GET("/", controllers.GetAllTodos)
	router.GET("/todo", controllers.GetAllTodos)

	ApiRoutes := router.Group("/api")
	{
		ApiRoutes.POST("/todo/create", controllers.CreateTodo)
		ApiRoutes.GET("/todo", controllers.GetAllTodos)
		ApiRoutes.PUT("/todo/:idTodo", controllers.UpdateTodo)
		ApiRoutes.DELETE("/todo/:idTodo", controllers.DeleteTodo)
	}

	router.Run()
}
