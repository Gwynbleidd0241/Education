package repository

import (
	"Kursash/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CourseListPostgres struct {
	db *sqlx.DB
}

func NewCourseListPostgres(db *sqlx.DB) *CourseListPostgres {
	return &CourseListPostgres{db: db}
}

func (r *CourseListPostgres) Create(userId int, list models.CourseList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description, image_url, course_url) VALUES ($1, $2, $3, $4) RETURNING id", courseListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description, list.Image_url, list.Course_url)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CourseListPostgres) GetAll(userId int) ([]models.CourseList, error) {
	var lists []models.CourseList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.image_url, tl.course_url FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		courseListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}
