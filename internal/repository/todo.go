package repository

import (
	"go-todo/domain"
	"go-todo/domain/value"
)

// type ToDoRepository interface {
// 	FindAll() ([]domain.ToDo, error)
// 	FindByID(id value.ToDoID) (*domain.ToDo, error)
// 	Store(todo domain.ToDo) error
// 	Update(id value.ToDoID, todo domain.ToDo) error
// 	Delete(id value.ToDoID) error
// }

type ToDoRepository struct{}

func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{}
}

var (
	todos = map[value.ToDoID]*domain.ToDo{}
)

func (r *ToDoRepository) FindAll() ([]domain.ToDo, error) {
	allTodos := make([]domain.ToDo, 0, len(todos))
	for _, todo := range todos {
		allTodos = append(allTodos, *todo)
	}
	return allTodos, nil
}

func (r *ToDoRepository) FindByID(id value.ToDoID) (*domain.ToDo, error) {
	todo := todos[id]
	if todo == nil {
		return nil, domain.ErrToDoNotFound
	}
	return todo, nil
}

func (r *ToDoRepository) FindByUserID(userId value.UserID) ([]domain.ToDo, error) {
	userTodos := make([]domain.ToDo, 0)
	for _, todo := range todos {
		if todo.UserID == userId {
			userTodos = append(userTodos, *todo)
		}
	}
	return userTodos, nil
}

func (r *ToDoRepository) Store(todo domain.ToDo) error {
	todos[todo.ID] = &todo
	return nil
}

func (r *ToDoRepository) Update(todo domain.ToDo) error {
	todos[todo.ID] = &todo
	return nil
}

func (r *ToDoRepository) Delete(todoId value.ToDoID) error {
	delete(todos, todoId)
	return nil
}
