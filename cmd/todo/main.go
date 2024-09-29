package main

import (
	"fmt"
	"github.com/ramalabe/todo-app"
)

func main() {
	fmt.Println("Welcome to Todo App")
	todos := todo.Todos{}
	todos.Add("Buy milk")
}