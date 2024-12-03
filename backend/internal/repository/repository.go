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
	GetById(userId, listId int) (models.CourseList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input models.UpdateListInput) error
}

type CourseItem interface {
	Create(listId int, item models.CourseItem) (int, error)
	GetAll(userId, listId int) ([]models.CourseItem, error)
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
		CourseItem:    NewCourseItemPostgres(db),
	}
}
