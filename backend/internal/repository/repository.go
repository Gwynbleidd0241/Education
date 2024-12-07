package repository

import (
	"Kursash/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.UserModel) (int, error)
	GetUser(username, password string) (models.UserModel, error)
}

type Course interface {
	Create(userId int, list models.Course) (int, error)
	GetAll(userId int) ([]models.Course, error)
	GetById(userId, listId int) (models.Course, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input models.UpdateCourseInput) error
	GetCoursesByProfession(profession string) ([]models.Course, error)
	AddUserCourse(userId, courseId int) error
	GetUserCourses(userId int) ([]models.Course, error)
}

type Repository struct {
	Authorization
	Course
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Course:        NewCoursePostgres(db),
	}
}
