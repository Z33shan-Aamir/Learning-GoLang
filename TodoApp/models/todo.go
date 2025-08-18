package models

// filepath: /home/zeeshan/Projects/Learning-GoLang/TodoApp/models/todo.go

import (
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
	currentTime := time.Now().Local()
	return Todo{
		Name:      todoName,
		IsDone:    false,
		CreatedOn: currentTime.Format("03/02/2009"),
		UUID:      uuid.New(),
	}
}
