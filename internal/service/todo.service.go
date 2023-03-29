package service

import (
	"fmt"

	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/entity/dto"
	"github.com/muhangga/internal/repository"
)

type TodoService interface {
	FindAll() ([]entity.Todo, error)
	FindAllByActivityGroupID(id int64) ([]entity.Todo, error)
	FindById(id int64) (entity.Todo, error)
	Create(todo dto.TodoDTO) (entity.Todo, error)
	Update(id int, todo dto.TodoDTO) (entity.Todo, error)
	Delete(id int64) error
}

type TodoServiceImpl struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(service repository.TodoRepository) TodoService {
	return &TodoServiceImpl{todoRepository: service}
}

func (s *TodoServiceImpl) FindAll() ([]entity.Todo, error) {
	data, err := s.todoRepository.GetAll()
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *TodoServiceImpl) FindAllByActivityGroupID(id int64) ([]entity.Todo, error) {
	data, err := s.todoRepository.GetAllByActivitiesGroupID(id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *TodoServiceImpl) FindById(id int64) (entity.Todo, error) {
	data, err := s.todoRepository.GetById(id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *TodoServiceImpl) Create(todo dto.TodoDTO) (entity.Todo, error) {

	var todoEntity entity.Todo
	todoEntity.Title = todo.Title
	todoEntity.ActivityGroupID = int64(todo.ActivityGroupID)
	todoEntity.Priority = "very-high"

	data, err := s.todoRepository.Create(todoEntity)
	if err != nil {
		return entity.Todo{}, err
	}
	return data, nil
}

func (s *TodoServiceImpl) Update(id int, todo dto.TodoDTO) (entity.Todo, error) {

	todos, err := s.todoRepository.GetById(int64(id))
	if err != nil {
		return entity.Todo{}, err
	}

	todos.Title = todo.Title
	todos.Priority = todo.Priority
	todos.IsActive = todo.IsActive
	todos.UpdatedAt = todo.UpdatedAt

	if todo.Status != "ok" {
		return entity.Todo{}, fmt.Errorf("status is not ok")
	}

	data, err := s.todoRepository.Update(todos)
	if err != nil {
		return entity.Todo{}, err
	}
	return data, nil
}

func (s *TodoServiceImpl) Delete(id int64) error {

	findId, err := s.todoRepository.GetById(id)
	if err != nil {
		return err
	}

	if findId.ID == 0 {
		return err
	}

	err = s.todoRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
