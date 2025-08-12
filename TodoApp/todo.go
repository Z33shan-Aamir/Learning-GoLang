package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	name       string
	is_done    bool
	created_on string
	uuid       uuid.UUID
}

func NewTodo(todo_name string) *Todo {
	current_time := time.Now().Local()
	return &Todo{
		name:       "New todo",
		is_done:    false,
		created_on: current_time.Format("03/02/2009"),
		uuid:       uuid.New(),
	}
}

func RemoveTodo(index int) {

}

func printHelp() {
	fmt.Println("Usage: go run todo.go [command]")
	fmt.Println("Available commands:")
	fmt.Println("  help       Show this help message")
	fmt.Println("  start      Start the application")
	fmt.Println("  stop       Stop the application")
	fmt.Println("  status     Show the current status")
	fmt.Println("  version    Show application version")
}

func check_args(arguments []string) any {
	for i := 0; i < len(arguments); i++ {
		switch strings.ToLower(arguments[i]) {
		case "help":
			printHelp()
		case "add":
			if len(arguments) >= i+1 {
				todo_name := arguments[i+1]
				return create_todo(todo_name)
			} else {
				fmt.Println("Are you stoopid!? \nEnter the todo name nigga!")
			}
		case "remove":
			if len(arguments) >= i+1 {

			}
		}
	}

	return nil
}

func create_todo(todo_name string) *Todo {
	Todo := NewTodo(todo_name)
	return Todo

}
func main() {
	// todo_items := map[]interface{}
	fmt.Println("Welcome to Zesshan's Over complicated Todo App")
	command_line_arguments := os.Args
	// checks the command line arguments
	result := check_args(command_line_arguments)
	fmt.Println(result)
	switch result := result.(type) {
	case *Todo:

		fmt.Println("reached the case in the main")
		fmt.Println(result.uuid, result.name, result.created_on, result.is_done)
		// todo_items[todo.uuid] = map[string]interface{}{}

	case int:

		fmt.Println("Yeah case got an int")

	case nil:

		fmt.Println("`check_args` gave me nothing")

	default:

		fmt.Println("Can `check_args` give me an actual value")

	}

}
