package app

import (
	"time"
)

type Storage interface {
	GetNewId() int
	NewTodo(todo Todo) error
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

func (t *Todo) Save(store Storage) error {
	err := store.NewTodo(*t)
	return err
}

func New(name string, description string, store Storage) error {
	todo := NewTodo(name, description, store)
	err := todo.Save(store)
	if err != nil {
		return err
	}
	return nil
}
