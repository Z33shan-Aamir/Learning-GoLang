package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/google/uuid"
	// "example.com/todo/models"
)

var permission os.FileMode = os.FileMode(0666)

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

// loads the data and also returns an empty map if there is now data in the file
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
