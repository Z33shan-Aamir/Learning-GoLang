package file_io

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/google/uuid"
)

var permision os.FileMode = os.FileMode(0666)

func save_todo_items(file_path string, todo_item any) {
	t := reflect.TypeOf(todo_item)
	if t.Kind() != reflect.Struct {
		fmt.Println("Please pass a struct of type Todo")
		return
	}
	data, err := json.Marshal(todo_item)
	if err != nil {
		panic(err)
	}
	os.WriteFile(file_path, data, permision)
	fmt.Println("data was succesfully written to ", file_path)
}

// loads the data
func load_todo_items(file_path string) map[uuid.UUID]any {
	data, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	var data_map map[uuid.UUID]any
	err = json.Unmarshal(data, &data_map)
	if err != nil {
		return make(map[uuid.UUID]any)
		// panic(err)
	}
	return data_map
}
