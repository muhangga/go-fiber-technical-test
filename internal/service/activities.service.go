package service

import (
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/entity/dto"
	"github.com/muhangga/internal/repository"
)

type ActivitiesService interface {
	GetAllActivities() ([]entity.Activities, error)
	GetActivitiesByID(id int64) (entity.Activities, error)
	CreateActivities(activities dto.ActivitiesDTO) (entity.Activities, error)
	UpdateActivities(req dto.ActivitiesDTO, id int) (entity.Activities, error)
	DeleteActivities(id int) error
	FindByEmail(email string) bool
}

func (s *ActivitiesServiceImpl) GetAllActivities() ([]entity.Activities, error) {
	data, err := s.activitiesRepository.GetAll()
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *ActivitiesServiceImpl) GetActivitiesByID(id int64) (entity.Activities, error) {
	data, err := s.activitiesRepository.GetByID(id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *ActivitiesServiceImpl) CreateActivities(activities dto.ActivitiesDTO) (entity.Activities, error) {

	var activitiesEntity entity.Activities
	activitiesEntity.Title = activities.Title
	activitiesEntity.Email = activities.Email
	activitiesEntity.CreatedAt = activities.CreatedAt
	activitiesEntity.UpdatedAt = activities.UpdatedAt

	return s.activitiesRepository.Create(activitiesEntity)
}

func (s *ActivitiesServiceImpl) UpdateActivities(req dto.ActivitiesDTO, id int) (entity.Activities, error) {

	activities, err := s.activitiesRepository.GetByID(int64(id))
	if err != nil {
		return activities, err
	}

	activities.Title = req.Title

	updated, err := s.activitiesRepository.Update(activities)
	if err != nil {
		return updated, err
	}

	return updated, nil
}

func (s *ActivitiesServiceImpl) DeleteActivities(id int) error {

	findID, err := s.activitiesRepository.GetByID(int64(id))
	if err != nil {
		return err
	}

	if findID.ID == 0 {
		return err
	}

	return s.activitiesRepository.Delete(id)
}

func (s *ActivitiesServiceImpl) FindByEmail(email string) bool {

	isExist := s.activitiesRepository.FindByEmail(email)
	if !isExist {
		return false
	}
	return true
}

type ActivitiesServiceImpl struct {
	activitiesRepository repository.ActivitiesRepository
}

func NewActivitiesService(activitiesRepository repository.ActivitiesRepository) ActivitiesService {
	return &ActivitiesServiceImpl{activitiesRepository}
}
