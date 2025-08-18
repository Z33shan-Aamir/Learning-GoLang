package main

import (
	"fmt"
	"os"
	"strings"

	// "golang.org/x/text/date"

	"example.com/todo/models"
	"example.com/todo/utils"
)

var file_path string = "./todo.json"

// saves data to a json file which in this case are going to be todo items.
func main() {
	todo_items_map := utils.LoadTodoItems("./todo.json")

	// todo_items_slice_uuid := make([]uuid.UUID, 0)
	fmt.Println("Welcome to Zesshan's Over complicated Todo App")
	command_line_arguments := os.Args
	// command_line_arguments := []string{"list"}

	// checks the command line arguments
	result := check_args(command_line_arguments)
	if result != nil {
		fmt.Println(result)
	}

	switch result := result.(type) {
	case models.Todo:

		// fmt.Println("reached the case in the main")
		// fmt.Println(result.name, result.created_on, result.is_done
		todo_items_map[result.UUID] = []any{
			result.Name,
			result.CreatedOn,
			result.IsDone,
		}
		utils.SaveTodoItems("./todo.json", todo_items_map)
		fmt.Println("New Tdod was added")
		// fmt.Println(todo_items_map)
		// fmt.Println(todo_items_slice_uuid)

	case int:

		fmt.Println("Yeah case got an int")
	case nil:
		fmt.Println("")

	}

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
				fmt.Println("Reached the remove case")
			}
		case "list":
			_ = list_todo(false)

		}
	}
	return nil
}

func list_todo(called_for_uuid_list bool) any {

	loaded_data := utils.LoadTodoItems(file_path)
	i := 1
	fmt.Printf(``)
	var keys []any
	for key, values := range loaded_data {
		if !called_for_uuid_list {
			fmt.Printf("%d. Name:  '%s'    Created on: '%s'    is done: '%t'\n", i, values[0], values[1], values[2])
		}
		i += 1
		keys = append(keys, key)
	}
	if called_for_uuid_list {
		return keys
	}
	return nil
}

func create_todo(todo_name string) models.Todo {
	Todo := models.NewTodo(todo_name)
	return Todo

}
