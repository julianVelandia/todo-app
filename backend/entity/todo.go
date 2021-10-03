package entity

type Todo struct {
	title string
	description string
	isCompleted bool
}

func NewTodo(title string, description string, isCompleted bool) Todo{
	return Todo{
		title: title,
		description: description,
		isCompleted: isCompleted,
	}
}

func (t Todo) Title() string{
	return t.title
}

func (t Todo) Description() string{
	return t.description
}

func (t Todo) IsCompleted() bool{
	return t.isCompleted
}