package app

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	NextId int
	Data   *Todo
	Error  error
}

func (m *MockStorage) GetNewId() int {
	return m.NextId
}

func (m *MockStorage) NewTodo(todo Todo) error {
	if m.Error != nil {
		return m.Error
	}
	m.Data = &todo
	return nil
}

func TestNewTodo(t *testing.T) {
	mock := MockStorage{NextId: 5}
	want := &Todo{
		Id:          5,
		Name:        "test todo",
		Description: "test description",
		Active:      true,
	}
	actual := NewTodo("test todo", "test description", &mock)
	assert.NotNil(t, actual.CreatedDate)
	actual.CreatedDate = nil
	assert.Equal(t, want, actual)
}

func TestTodoSave(t *testing.T) {
	dt := time.Now()
	mock := MockStorage{}
	todo := &Todo{
		Id:          10,
		Name:        "Test Todo",
		Description: "a todo that needs to be done",
		Active:      true,
		CreatedDate: &dt,
	}
	err := todo.Save(&mock)
	if err != nil {
		t.Fatalf("error running Todo Save method, %v", err)
	}
	assert.Equal(t, todo, mock.Data)
	saveErr := fmt.Errorf("some error")
	mockErr := MockStorage{Error: saveErr}
	err = todo.Save(&mockErr)
	if err == nil {
		t.Fatalf("expected error not recieved")
	}
	assert.Equal(t, saveErr, err)
}
