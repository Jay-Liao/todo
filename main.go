package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID          int
	DESCRIPTION string
}

func getTodos(context *gin.Context) {
	context.JSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	context.BindJSON(&newTodo)
	lastTodo := todos[len(todos)-1]
	lastID := lastTodo.ID
	newTodo.ID = lastID + 1
	todos = append(todos, newTodo)
	context.JSON(http.StatusOK, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	theTodo := todo{-1, "Not Found"}
	for _, todo := range todos {
		if strconv.Itoa(todo.ID) == id {
			theTodo = todo
		}
	}
	if theTodo.ID == -1 {
		context.String(http.StatusNotFound, "Todo with id %s was not found", id)
	} else {
		context.JSON(http.StatusOK, theTodo)
	}
}

func updateTodo(context *gin.Context) {
	var newTodo todo
	context.BindJSON(&newTodo)
	id := context.Param("id")
	theTodo := todo{-1, "Not Found"}
	target := -1
	for index, todo := range todos {
		if strconv.Itoa(todo.ID) == id {
			target = index
		}
	}
	if target >= 0 {
		newTodo.ID = todos[target].ID
		todos[target] = newTodo
		theTodo = newTodo
	}
	if theTodo.ID == -1 {
		context.String(http.StatusNotFound, "Todo with id %s was not found", id)
	} else {
		context.JSON(http.StatusOK, theTodo)
	}
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	target := -1
	for index, todo := range todos {
		if strconv.Itoa(todo.ID) == id {
			target = index
		}
	}

	if target >= 0 {
		todos[target] = todos[len(todos)-1]
		todos = todos[:len(todos)-1]
		context.String(http.StatusOK, "Todo with id %s has been deleted", id)
	} else {
		context.String(http.StatusNotFound, "Todo with id %s was not found", id)
	}
}

var todos []todo

func main() {
	todos = append(todos, todo{1, "Learning golang"})
	todos = append(todos, todo{2, "Learning node js"})

	router := gin.Default()

	router.GET("/api/todos", getTodos)
	router.POST("/api/todos", addTodo)

	router.GET("/api/todos/:id", getTodo)
	router.PATCH("/api/todos/:id", updateTodo)
	router.DELETE("api/todos/:id", deleteTodo)
	router.Run(":8080")
}
