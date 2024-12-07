package service

import (
	"Kursash/internal/models"
	"Kursash/internal/repository"
)

type CourseService struct {
	repo repository.Course
}

func NewCourseService(repo repository.Course) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) Create(userId int, list models.Course) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *CourseService) GetAll(userId int) ([]models.Course, error) {
	return s.repo.GetAll(userId)
}

func (s *CourseService) GetById(userId, listId int) (models.Course, error) {
	return s.repo.GetById(userId, listId)
}

func (s *CourseService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *CourseService) Update(userId, listId int, input models.UpdateCourseInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}

func (s *CourseService) GetByProfession(profession string) ([]models.Course, error) {
	return s.repo.GetCoursesByProfession(profession)
}

func (s *CourseService) AddUserCourse(userId, courseId int) error {
	return s.repo.AddUserCourse(userId, courseId)
}

func (s *CourseService) GetUserCourses(userId int) ([]models.Course, error) {
	return s.repo.GetUserCourses(userId)
}
