package app

import "time"

type Storage interface {
	GetNewId() int
}

type Todo struct {
	Id            int
	Name          string
	Description   string
	Active        bool
	CreatedDate   *time.Time
	CompletedDate *time.Time
}

func NewTodo(name string, desc string, store Storage) *Todo {
	id := store.GetNewId()
	now := time.Now()
	todo := Todo{
		Id:          id,
		Name:        name,
		Description: desc,
		Active:      true,
		CreatedDate: &now,
	}
	return &todo
}
