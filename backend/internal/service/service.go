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

type CourseList interface {
	Create(userId int, list models.CourseList) (int, error)
	GetAll(userId int) ([]models.CourseList, error)
	GetById(userId, listId int) (models.CourseList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input models.UpdateListInput) error
}

type CourseItem interface {
	Create(userId, listId int, item models.CourseItem) (int, error)
	GetAll(userId, listId int) ([]models.CourseItem, error)
}

type Service struct {
	Authorization
	CourseList
	CourseItem
	Notifications *Notifications
}

func NewService(repos *repository.Repository, notifications *Notifications) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		CourseList:    NewCourseListService(repos.CourseList),
		CourseItem:    NewCourseItemService(repos.CourseItem, repos.CourseList),
		Notifications: notifications,
	}
}
