package repository

import (
	"Kursash/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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

func (r *CourseListPostgres) GetById(userId, listId int) (models.CourseList, error) {
	var list models.CourseList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description, tl.image_url, tl.course_url FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		courseListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *CourseListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		courseListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *CourseListPostgres) Update(userId, listId int, input models.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Image_url != nil {
		setValues = append(setValues, fmt.Sprintf("image_url=$%d", argId))
		args = append(args, *input.Image_url)
		argId++
	}

	if input.Course_url != nil {
		setValues = append(setValues, fmt.Sprintf("course_url=$%d", argId))
		args = append(args, *input.Course_url)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		courseListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
