package routes

import (
	"net/http"

	"github.com/fromcode/to-do-app/src/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.GET("/todo", controllers.GetAllTodos)
	router.GET("/todo/create/form", func(context *gin.Context) {
		context.HTML(http.StatusOK, "create.html", nil)
	})
	router.POST("/todo/create", controllers.HandleSubmit)

	ApiRoutes := router.Group("/api")
	{
		ApiRoutes.POST("/todo/create", controllers.CreateTodo)
		ApiRoutes.GET("/todo", controllers.GetAllTodos)
		ApiRoutes.PUT("/todo/:idTodo", controllers.UpdateTodo)
		ApiRoutes.DELETE("/todo/:idTodo", controllers.DeleteTodo)
	}

	router.Run()
}
