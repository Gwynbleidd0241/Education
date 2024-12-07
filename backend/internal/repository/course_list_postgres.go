package repository

import (
	"Kursash/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type CoursePostgres struct {
	db *sqlx.DB
}

func NewCoursePostgres(db *sqlx.DB) *CoursePostgres {
	return &CoursePostgres{db: db}
}

func (r *CoursePostgres) Create(userId int, list models.Course) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description, image_url, course_url, profession) VALUES ($1, $2, $3, $4, $5) RETURNING id", courseTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description, list.Image_url, list.Course_url, list.Profession)
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

func (r *CoursePostgres) GetAll(userId int) ([]models.Course, error) {
	var lists []models.Course

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.image_url, tl.course_url, tl.profession FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		courseTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *CoursePostgres) GetById(userId, listId int) (models.Course, error) {
	var list models.Course

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description, tl.image_url, tl.course_url, tl.profession FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		courseTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *CoursePostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		courseTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *CoursePostgres) Update(userId, listId int, input models.UpdateCourseInput) error {
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

	if input.Profession != nil {
		setValues = append(setValues, fmt.Sprintf("profession=$%d", argId))
		args = append(args, *input.Profession)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		courseTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *CoursePostgres) GetCoursesByProfession(profession string) ([]models.Course, error) {
	var courses []models.Course
	query := `SELECT id, title, description, image_url, course_url
              FROM course WHERE profession = $1`
	err := r.db.Select(&courses, query, profession)
	return courses, err
}

func (r *CoursePostgres) AddUserCourse(userId, courseId int) error {
	query := `INSERT INTO user_courses (user_id, course_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := r.db.Exec(query, userId, courseId)
	return err
}

func (r *CoursePostgres) GetUserCourses(userId int) ([]models.Course, error) {
	var courses []models.Course
	query := `
        SELECT c.id, c.title, c.description, c.image_url, c.course_url
        FROM user_courses uc
        INNER JOIN course c ON c.id = uc.course_id
        WHERE uc.user_id = $1
    `

	err := r.db.Select(&courses, query, userId)
	return courses, err
}
