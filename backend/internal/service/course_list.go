package service

import (
	"Kursash/internal/models"
	"Kursash/internal/repository"
)

type CourseListService struct {
	repo repository.CourseList
}

func NewCourseListService(repo repository.CourseList) *CourseListService {
	return &CourseListService{repo: repo}
}

func (s *CourseListService) Create(userId int, list models.CourseList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *CourseListService) GetAll(userId int) ([]models.CourseList, error) {
	return s.repo.GetAll(userId)
}

func (s *CourseListService) GetById(userId, listId int) (models.CourseList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *CourseListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *CourseListService) Update(userId, listId int, input models.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}
