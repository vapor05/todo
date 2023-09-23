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
	newFile := false
	file, err = os.Open(filename)
	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
		newFile = true
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	if !newFile {
		data, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &todoData); err != nil {
			return nil, err
		}
	}

	jsonStore := JSONStorage{
		File: filename,
		Data: &todoData,
	}
	return &jsonStore, nil
}
