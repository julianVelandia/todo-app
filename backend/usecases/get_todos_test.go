package usecases_test

import (
	"fmt"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/julianVelandia/todo-app/backend/entity"
	"github.com/julianVelandia/todo-app/backend/usecases"
)

var dummyTodos = []entity.Todo{
	entity.NewTodo("todo 1", "description 1", true),
	entity.NewTodo("todo 2", "description 2", false),
	entity.NewTodo("todo 3", "description 3", true),
}

type BadTodosRepo struct{}

func (BadTodosRepo) GetAllTodos() ([]entity.Todo, error){
	return nil, fmt.Errorf("something went wrong")
}

type MockTodosRepo struct{}

func (MockTodosRepo) GetAllTodos() ([]entity.Todo, error){
	return dummyTodos, nil
}

func TestGetTodosWhenSuccessfullyShouldReturnOk (t *testing.T){
	expect := expectate.Expect(t)
	repo := new(MockTodosRepo)

	_ ,err := usecases.GetTodos(repo)

	expect(err).ToBe(nil)
	//expect(todos).ToBe(dummyTodos)
}


func TestGetTodosWhenFailShouldReturnError (t *testing.T){

	expect := expectate.Expect(t)
	repo := new(BadTodosRepo)

	todos ,err := usecases.GetTodos(repo)

	expect(err).ToBe(usecases.ErrInternal)
	if todos != nil{
		t.Fatalf("Expected todos to be nil; Got; %v", todos)
	}
}

