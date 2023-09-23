package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	NextId int
}

func (m *MockStorage) GetNewId() int {
	return m.NextId
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
