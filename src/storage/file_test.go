package storage

import (
	"encoding/json"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vapor05/todo/src/app"
)

func TestGetNewId(t *testing.T) {
	store := &JSONStorage{
		File: "test.json",
		Data: &JSONTodoData{},
	}
	want := 1
	actual := store.GetNewId()
	assert.Equal(t, want, actual)
	cd1 := time.Date(2023, 9, 23, 3, 30, 0, 0, time.Local)
	cd2 := time.Date(2023, 9, 23, 3, 32, 0, 0, time.Local)
	cd3 := time.Date(2023, 9, 23, 3, 35, 0, 0, time.Local)
	cmp3 := time.Date(2023, 9, 23, 3, 40, 0, 0, time.Local)
	store = &JSONStorage{
		File: "test.json",
		Data: &JSONTodoData{
			Order: []int{1, 2, 3},
			Todos: map[int]*app.Todo{
				1: {
					Id:          1,
					Name:        "test1",
					Description: "a test todo",
					Active:      true,
					CreatedDate: &cd1,
				},
				2: {
					Id:          2,
					Name:        "todo2",
					Description: "another test todo",
					Active:      true,
					CreatedDate: &cd2,
				},
				3: {
					Id:            3,
					Name:          "some task",
					Description:   "some todo i need to do",
					Active:        false,
					CreatedDate:   &cd3,
					CompletedDate: &cmp3,
				},
			},
			NameIndex: map[string][]int{
				"test1":     {1},
				"todo2":     {2},
				"some task": {3},
			},
		},
	}
	want = 4
	actual = store.GetNewId()
	assert.Equal(t, want, actual)
}

func TestNewJSONStorage(t *testing.T) {
	// no existing data file
	storage, err := NewJSONStorage("test.json")
	if err != nil {
		t.Fatalf("NewJSONStorage failed to run, %v", err)
	}
	want := &JSONStorage{
		File: "test.json",
		Data: &JSONTodoData{},
	}
	assert.Equal(t, want, storage)
	// existing data file
	cd1 := time.Date(2023, 9, 23, 3, 30, 0, 0, time.Local)
	cd2 := time.Date(2023, 9, 23, 3, 32, 0, 0, time.Local)
	cd3 := time.Date(2023, 9, 23, 3, 35, 0, 0, time.Local)
	cmp3 := time.Date(2023, 9, 23, 3, 40, 0, 0, time.Local)
	want = &JSONStorage{
		File: "test.json",
		Data: &JSONTodoData{
			Order: []int{1, 2, 3},
			Todos: map[int]*app.Todo{
				1: {
					Id:          1,
					Name:        "test1",
					Description: "a test todo",
					Active:      true,
					CreatedDate: &cd1,
				},
				2: {
					Id:          2,
					Name:        "todo2",
					Description: "another test todo",
					Active:      true,
					CreatedDate: &cd2,
				},
				3: {
					Id:            3,
					Name:          "some task",
					Description:   "some todo i need to do",
					Active:        false,
					CreatedDate:   &cd3,
					CompletedDate: &cmp3,
				},
			},
			NameIndex: map[string][]int{
				"test1":     {1},
				"todo2":     {2},
				"some task": {3},
			},
		},
	}
	data, err := json.Marshal(want.Data)
	if err != nil {
		t.Fatalf("failed to write test data, %v", err)
	}
	file, err := os.Create("test.json")
	if err != nil {
		t.Fatalf("failed to write test data, %v", err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		t.Fatalf("failed to write test data, %v", err)
	}
	file.Close()
	actual, err := NewJSONStorage("test.json")
	if err != nil {
		t.Fatalf("error running the NewJSONStorage func, %v", err)
	}
	assert.Equal(t, want, actual)
	if err = os.Remove("test.json"); err != nil {
		t.Errorf("failed to clean up after test, %v", err)
	}
}

func TestNewTodo(t *testing.T) {
	cd := time.Now()
	todo := app.Todo{
		Id:          12,
		Name:        "test todo",
		Description: "a task to do",
		Active:      true,
		CreatedDate: &cd,
	}
	cd1 := time.Date(2023, 9, 23, 3, 30, 0, 0, time.Local)
	cd2 := time.Date(2023, 9, 23, 3, 32, 0, 0, time.Local)
	store := JSONStorage{
		File: "testsave.json",
		Data: &JSONTodoData{
			Order: []int{1, 2},
			Todos: map[int]*app.Todo{
				1: {
					Id:          1,
					Name:        "first todo",
					Description: "some desc",
					Active:      true,
					CreatedDate: &cd1,
				},
				2: {
					Id:          2,
					Name:        "next todo",
					Description: "some desc2",
					Active:      true,
					CreatedDate: &cd2,
				},
			},
			NameIndex: map[string][]int{
				"first todo": {1},
				"next todo":  {2},
			},
		},
	}
	want := JSONStorage{
		File: "testsave.json",
		Data: &JSONTodoData{
			Order: []int{1, 2, 12},
			Todos: map[int]*app.Todo{
				1: {
					Id:          1,
					Name:        "first todo",
					Description: "some desc",
					Active:      true,
					CreatedDate: &cd1,
				},
				2: {
					Id:          2,
					Name:        "next todo",
					Description: "some desc2",
					Active:      true,
					CreatedDate: &cd2,
				},
				12: {
					Id:          12,
					Name:        "test todo",
					Description: "a task to do",
					Active:      true,
					CreatedDate: &cd,
				},
			},
			NameIndex: map[string][]int{
				"first todo": {1},
				"next todo":  {2},
				"test todo":  {12},
			},
		},
	}
	store.NewTodo(todo)
	assert.Equal(t, want, store)
	// new todo with same name as existing todo
	cd3 := time.Now()
	todo = app.Todo{
		Id:          15,
		Name:        "test todo",
		Description: "some new task to do",
		Active:      true,
		CreatedDate: &cd3,
	}
	store.NewTodo(todo)
	want = JSONStorage{
		File: "testsave.json",
		Data: &JSONTodoData{
			Order: []int{1, 2, 12, 15},
			Todos: map[int]*app.Todo{
				1: {
					Id:          1,
					Name:        "first todo",
					Description: "some desc",
					Active:      true,
					CreatedDate: &cd1,
				},
				2: {
					Id:          2,
					Name:        "next todo",
					Description: "some desc2",
					Active:      true,
					CreatedDate: &cd2,
				},
				12: {
					Id:          12,
					Name:        "test todo",
					Description: "a task to do",
					Active:      true,
					CreatedDate: &cd,
				},
				15: {
					Id:          15,
					Name:        "test todo",
					Description: "some new task to do",
					Active:      true,
					CreatedDate: &cd3,
				},
			},
			NameIndex: map[string][]int{
				"first todo": {1},
				"next todo":  {2},
				"test todo":  {12, 15},
			},
		},
	}
	assert.Equal(t, want, store)
}

func TestSave(t *testing.T) {
	cd1 := time.Date(2023, 9, 23, 3, 30, 0, 0, time.Local)
	cd2 := time.Date(2023, 9, 23, 3, 32, 0, 0, time.Local)
	store := JSONStorage{
		File: "testsave.json",
		Data: &JSONTodoData{
			Order: []int{1, 2},
			Todos: map[int]*app.Todo{
				1: {
					Id:          1,
					Name:        "first todo",
					Description: "some desc",
					Active:      true,
					CreatedDate: &cd1,
				},
				2: {
					Id:          2,
					Name:        "next todo",
					Description: "some desc2",
					Active:      true,
					CreatedDate: &cd2,
				},
			},
			NameIndex: map[string][]int{
				"first todo": {1},
				"next todo":  {2},
			},
		},
	}
	err := store.Save()
	if err != nil {
		t.Fatalf("error running Save method, %v", err)
	}
	f, err := os.Open("testsave.json")
	if err != nil {
		t.Fatalf("failed to open saved json file, %v", err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to load saved json file, %v", err)
	}
	var actual JSONTodoData
	if err := json.Unmarshal(data, &actual); err != nil {
		t.Fatalf("failed to load saved json file, %v", err)
	}
	assert.Equal(t, store.Data, &actual)
}
