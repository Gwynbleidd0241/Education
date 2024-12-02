package repository

import (
	"Kursash/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.UserModel) (int, error)
	GetUser(username, password string) (models.UserModel, error)
}

type CourseList interface {
	Create(userId int, list models.CourseList) (int, error)
	GetAll(userId int) ([]models.CourseList, error)
}

type CourseItem interface {
}

type Repository struct {
	Authorization
	CourseList
	CourseItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		CourseList:    NewCourseListPostgres(db),
	}
}
