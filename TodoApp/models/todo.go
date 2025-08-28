package models

// filepath: /home/zeeshan/Projects/Learning-GoLang/TodoApp/models/todo.go

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Name      string
	IsDone    bool
	CreatedOn string
	UUID      uuid.UUID
}

func NewTodo(todoName string) Todo {
	now := time.Now()
	year, month, day := now.Date()
	// currentTime := time.Now().Local()
	return Todo{
		Name:      todoName,
		IsDone:    false,
		CreatedOn: fmt.Sprintf("%d-%02d-%02d", year, month, day),
		UUID:      uuid.New(),
	}
}
