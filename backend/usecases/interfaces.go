package usecases

import "github.com/julianVelandia/todo-app/backend/entity"

type TodosRepository interface {
	GetAllTodos()([]entity.Todo, error)
}
