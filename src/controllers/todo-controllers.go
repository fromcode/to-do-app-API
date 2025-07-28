package controllers

import (
	"fmt"
	"net/http"

	"github.com/fromcode/to-do-app/src/config"
	"github.com/fromcode/to-do-app/src/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// Database Client
var db *gorm.DB = config.ConnectDB()

// todo struct untuk request body
type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// struct untuk response
type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

// buat Todo data ke database dengan menjalankan kode ini
func CreateTodo(context *gin.Context) {
	var data todoRequest

	// target request body json ke request body struct
	if err := context.ShouldBindBodyWithJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Matching todo models struct dengan todo request struct
	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	// Query ke database
	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}
	// matching result untuk create response
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// create POST http response
	context.JSON(http.StatusCreated, response)
}

// GET all todos data
func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	//Querying to find todo data
	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
	}

	// creating GET http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})
}

func UpdateTodo(context *gin.Context) {
	var data todoRequest

	// targeting request parameter untuk get todo id
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	// menghubungkan request body ke request body struct
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// inisialisasi models Todo
	todo := models.Todo{}

	// Querying find todo data by todo id from request parameter
	todoById := db.Where("id = ?", idTodo).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Task not found"})
		return
	}

	// Matching todo request dengan todo models
	todo.Name = data.Name
	todo.Description = data.Description

	// Update new todo data
	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	// Matching result to todo response struct
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Creating http response
	context.JSON(http.StatusCreated, response)
}

// Delete todo data function
func DeleteTodo(context *gin.Context) {
	// Initiate todo models
	todo := models.Todo{}
	// Getting request parameter id
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	// Querying delete todo by id
	delete := db.Where("id = ?", idTodo).Unscoped().Delete(&todo)
	fmt.Println(delete)

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idTodo,
	})

}
