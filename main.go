package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID          int
	Description string
}

func getTodos(context *gin.Context) {
	var todos []todo
	for _, todoItem := range todoMap {
		todos = append(todos, todoItem)
	}
	context.JSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	context.BindJSON(&newTodo)
	latestID := len(todoMap) + 1
	newTodo.ID = latestID
	newIDString := strconv.Itoa(latestID)
	todoMap[newIDString] = newTodo
	context.JSON(http.StatusOK, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	theTodo, exist := todoMap[id]
	if exist {
		context.JSON(http.StatusOK, theTodo)
	} else {
		context.String(http.StatusNotFound, "Todo with id %s was not found", id)
	}
}

func updateTodo(context *gin.Context) {
	var newTodo todo
	context.BindJSON(&newTodo)
	id := context.Param("id")
	theTodo, exist := todoMap[id]
	if exist {
		theTodo.Description = newTodo.Description
		todoMap[id] = theTodo
		context.JSON(http.StatusOK, theTodo)
	} else {
		context.String(http.StatusNotFound, "Todo with id %s was not found", id)
	}
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	_, exist := todoMap[id]
	if exist {
		delete(todoMap, id)
		context.String(http.StatusOK, "Todo with id %s has been deleted", id)
	} else {
		context.String(http.StatusNotFound, "Todo with id %s was not found", id)
	}
}

var todoMap map[string]todo

func main() {
	todoMap = map[string]todo{
		"1": todo{1, "Learning golang"},
		"2": todo{2, "Learning node js"},
	}

	router := gin.Default()

	// /api/todos
	router.GET("/api/todos", getTodos)
	router.POST("/api/todos", addTodo)
	// /api/todos/:id
	router.GET("/api/todos/:id", getTodo)
	router.PATCH("/api/todos/:id", updateTodo)
	router.DELETE("api/todos/:id", deleteTodo)
	router.Run(":8080")
}
