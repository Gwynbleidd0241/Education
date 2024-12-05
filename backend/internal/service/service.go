package service

import (
	"Kursash/internal/models"
	"Kursash/internal/repository"
	"Kursash/notifications"
)

type Notifications struct {
	MailSender *notifications.MailSender
}

type Authorization interface {
	CreateUser(user models.UserModel) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Course interface {
	Create(userId int, list models.Course) (int, error)
	GetAll(userId int) ([]models.Course, error)
	GetById(userId, listId int) (models.Course, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input models.UpdateCourseInput) error
}

type Service struct {
	Authorization
	Course
	Notifications *Notifications
}

func NewService(repos *repository.Repository, notifications *Notifications) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Course:        NewCourseService(repos.Course),
		Notifications: notifications,
	}
}
