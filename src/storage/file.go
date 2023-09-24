package storage

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/vapor05/todo/src/app"
)

type JSONTodoData struct {
	Order     []int
	Todos     map[int]*app.Todo
	NameIndex map[string][]int
}

type JSONStorage struct {
	File string
	Data *JSONTodoData
}

func NewJSONStorage(filename string) (*JSONStorage, error) {
	var file *os.File
	var err error
	var todoData JSONTodoData
	file, err = os.Open(filename)
	if errors.Is(err, os.ErrNotExist) {
		todoData.Todos = make(map[int]*app.Todo)
		todoData.NameIndex = make(map[string][]int)
		jsonStore := JSONStorage{
			File: filename,
			Data: &todoData,
		}
		return &jsonStore, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &todoData); err != nil {
		return nil, err
	}
	jsonStore := JSONStorage{
		File: filename,
		Data: &todoData,
	}
	return &jsonStore, nil
}

func (s *JSONStorage) GetNewId() int {
	var maxId int
	for id := range s.Data.Todos {
		if id > maxId {
			maxId = id
		}
	}
	return maxId + 1
}

func (s *JSONStorage) NewTodo(todo app.Todo) error {
	s.Data.Todos[todo.Id] = &todo
	s.Data.Order = append(s.Data.Order, todo.Id)
	ni, ok := s.Data.NameIndex[todo.Name]
	if !ok {
		ni = make([]int, 0, 1)
	}
	ni = append(ni, todo.Id)
	s.Data.NameIndex[todo.Name] = ni
	if err := s.Save(); err != nil {
		return err
	}
	return nil
}

func (s JSONStorage) Save() error {
	file, err := os.Create(s.File)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(s.Data)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (s JSONStorage) ListTodos() ([]*app.Todo, error) {
	list := make([]*app.Todo, 0, len(s.Data.Todos))
	for _, t := range s.Data.Todos {
		list = append(list, t)
	}
	return list, nil
}
