package repository

import (
	"github.com/muhangga/internal/entity"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAll() ([]entity.Todo, error)
	GetAllByActivitiesGroupID(id int64) ([]entity.Todo, error)
	GetById(id int64) (entity.Todo, error)
	Create(todo entity.Todo) (entity.Todo, error)
	Update(todo entity.Todo) (entity.Todo, error)
	Delete(id int64) error
}

func (r *TodoRepositoryImpl) GetAll() ([]entity.Todo, error) {
	var todos []entity.Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepositoryImpl) GetAllByActivitiesGroupID(id int64) ([]entity.Todo, error) {
	var todos []entity.Todo
	err := r.db.Where("activity_group_id = ?", id).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepositoryImpl) GetById(id int64) (entity.Todo, error) {
	var todo entity.Todo
	err := r.db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) Create(todo entity.Todo) (entity.Todo, error) {
	tx := r.db.Begin()

	save := tx.Table("todos").Create(&todo).Error
	if save != nil {
		if tx.Rollback().Error != nil {
			return todo, save
		}
		return todo, save
	}
	tx.Commit()
	return todo, nil
}

func (r *TodoRepositoryImpl) Update(todo entity.Todo) (entity.Todo, error) {

	tx := r.db.Begin()

	save := tx.Table("todos").Save(&todo).Error
	if save != nil {
		if tx.Rollback().Error != nil {
			return todo, save
		}
		return todo, save
	}
	tx.Commit()
	return todo, nil
}

func (r *TodoRepositoryImpl) Delete(id int64) error {
	var todo entity.Todo

	data := r.db.Table("todos").Where("id = ?", id).Delete(&todo)
	if data != nil {
		return data.Error
	}

	return nil
}

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{db}
}
