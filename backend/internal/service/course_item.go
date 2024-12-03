package service

import (
	"Kursash/internal/models"
	"Kursash/internal/repository"
)

type CourseItemService struct {
	repo     repository.CourseItem
	listRepo repository.CourseList
}

func NewCourseItemService(repo repository.CourseItem, listRepo repository.CourseList) *CourseItemService {
	return &CourseItemService{repo: repo, listRepo: listRepo}
}

func (s *CourseItemService) Create(userId, listId int, item models.CourseItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *CourseItemService) GetAll(userId, listId int) ([]models.CourseItem, error) {
	return s.repo.GetAll(userId, listId)
}
