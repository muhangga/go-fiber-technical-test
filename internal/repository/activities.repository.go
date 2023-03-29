package repository

import (
	"github.com/muhangga/internal/entity"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type ActivitiesRepository interface {
	GetAll() ([]entity.Activities, error)
	GetByID(id int64) (entity.Activities, error)
	Create(activities entity.Activities) (entity.Activities, error)
	Update(req entity.Activities) (entity.Activities, error)
	Delete(id int) error
	FindByEmail(email string) bool
}

type ActivitiesRepositoryImpl struct {
	db *gorm.DB
}

func NewActivitiesRepository(db *gorm.DB) ActivitiesRepository {
	return &ActivitiesRepositoryImpl{db}
}

func (r *ActivitiesRepositoryImpl) GetAll() ([]entity.Activities, error) {
	var activities []entity.Activities
	err := r.db.Find(&activities).Error
	if err != nil {
		log.Error().Msgf("Error when trying to get all activities: %v", err)
		return nil, err
	}
	return activities, nil
}

func (r *ActivitiesRepositoryImpl) GetByID(id int64) (entity.Activities, error) {
	var activities entity.Activities
	err := r.db.Where("id = ?", id).First(&activities).Error
	if err != nil {
		log.Error().Msgf("Error when trying to get activities by id: %v", err)
		return activities, err
	}
	return activities, nil
}

func (r *ActivitiesRepositoryImpl) Create(activities entity.Activities) (entity.Activities, error) {
	tx := r.db.Begin()

	save := tx.Table("activities").Create(&activities)
	if save.Error != nil {
		if tx.Rollback().Error != nil {
			log.Error().Msgf("Error when trying to rollback: %v", save.Error)
			return activities, save.Error
		}
		return activities, save.Error
	}

	tx.Commit()
	return activities, nil
}

func (r *ActivitiesRepositoryImpl) Update(req entity.Activities) (entity.Activities, error) {

	var activities entity.Activities

	tx := r.db.Begin()

	err := r.db.Save(&req).Error
	if err != nil {
		if tx.Rollback().Error != nil {
			log.Error().Msgf("Error when trying to rollback: %v", err)
			return activities, err
		}
		return activities, err
	}

	tx.Commit()
	return activities, nil
}

func (r *ActivitiesRepositoryImpl) Delete(id int) error {

	delele := r.db.Table("activities").Where("id = ?", id).Delete(&entity.Activities{})
	if delele.Error != nil {
		return delele.Error
	}
	return nil
}

func (r *ActivitiesRepositoryImpl) FindByEmail(email string) bool {
	var activities entity.Activities

	data := r.db.Table("activities")
	data.Select("email")
	data.Where("email = ?", email)
	data.Take(&activities)

	if data.Error != nil {
		return false
	}
	return true
}
