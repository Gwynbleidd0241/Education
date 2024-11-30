package service

import (
	"Kursash/internal/models"
	"Kursash/internal/repository"
	"context"
)

type Service struct {
	repo *repository.DB
	CourseService
	Authorization
}

type Authorization interface {
	CreateUser(user models.UserModel) (int, error)
}

type CourseService interface {
	AddCourse(ctx context.Context, course models.CourseModel) error
	GetCoursesList(ctx context.Context) ([]*models.CourseModel, error)
	ChangeCourse(ctx context.Context, id int, title, description, instructor string) error
	DeleteCourse(ctx context.Context, id int) error
	GetCourse(ctx context.Context, id int) (*models.CourseModel, error)
}

func NewService(repo *repository.DB) *Service {
	return &Service{
		repo: repo,
	}
}
