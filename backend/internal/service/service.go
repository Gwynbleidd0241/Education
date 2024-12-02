package service

import (
	"Kursash/internal/models"
	"Kursash/internal/repository"
)

type Authorization interface {
	CreateUser(user models.UserModel) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type CourseList interface {
	Create(userId int, list models.CourseList) (int, error)
	GetAll(userId int) ([]models.CourseList, error)
}

type CourseItem interface {
}

type Service struct {
	Authorization
	CourseList
	CourseItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		CourseList:    NewCourseListService(repos.CourseList),
	}
}
