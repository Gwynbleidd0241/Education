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
