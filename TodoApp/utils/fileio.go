package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"slices"

	"github.com/google/uuid"

	"example.com/todo/models"
	// "example.com/todo"
)

var permission os.FileMode = os.FileMode(0666)

// checks if file exists, if it doesn't exist returns true otherwise false
func DoesFileExist(file_path string) bool {
	_, err := os.Stat(file_path)

	if !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

func SaveTodoItems(file_path string, todo_items_map map[uuid.UUID][]any) {

	todo_items_map_reflect := reflect.TypeOf(todo_items_map).Kind() == reflect.Map
	if todo_items_map_reflect {
		data, err := json.MarshalIndent(todo_items_map, " ", "\t")
		if err != nil {
			panic(err)
		} else {
			os.WriteFile(file_path, data, permission)
			fmt.Println("data was succesfully written to ", file_path)
		}
	} else {
		fmt.Println("Please enter a valid value")
	}

}

// loads the todo items and returns an empty map if there is now data in the file
func LoadTodoItems(file_path string) map[uuid.UUID][]any {
	data, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	var data_map map[uuid.UUID][]any
	err = json.Unmarshal(data, &data_map)
	if err != nil {
		return make(map[uuid.UUID][]any)
		// panic(err)
	}
	return data_map
}

// loads the deleted todo and  returns an empty slice if there is no data in the file
func LoadDeletedTodoItemsUUIDs(file_path string) []uuid.UUID {
	data, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println("File does not exist creating file...")
		os.Create(file_path)
		return make([]uuid.UUID, 0)
	}
	var deleted_todo_slice []uuid.UUID
	err = json.Unmarshal(data, &deleted_todo_slice)
	if err != nil {
		// return make([]uuid.UUID, 0)
		fmt.Println("Could not decode json at line 57")
		panic(err)
	}
	return deleted_todo_slice
}

// automatically deletes todo items to keep in sync
func TrackDeletedTodo() {
	if DoesFileExist(models.FilePathDeletedTodo) {
		os.Create(models.FilePathDeletedTodo)
	}
	deleted_todo_items_uuid := LoadDeletedTodoItemsUUIDs(models.FilePathDeletedTodo)
	todo_items := LoadTodoItems(models.FilePathTodo)
	for key := range todo_items {
		if slices.Contains(deleted_todo_items_uuid, key) {
			delete(todo_items, key)
		}
	}
	// TODO: implement the logic to get all the todo items, iterate over the keys store them in a varaible and then compare them and if value of one exists in the other remove it from the todo.json file or what ever file it is store in

}
