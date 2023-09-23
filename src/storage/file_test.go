package storage

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vapor05/todo/src/app"
)

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
	err = os.Remove("test.json")
	if err != nil {
		t.Fatalf("failed to clean up after test, %v", err)
	}
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
					Id:            2,
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