package usecases

import "github.com/julianVelandia/todo-app/backend/entity"

func GetTodos(repo TodosRepository)([]entity.Todo, error){
	todos, err := repo.GetAllTodos()

	if err != nil {
		return nil, ErrInternal
	}
	return todos, nil
}
