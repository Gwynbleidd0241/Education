package repository

import (
	"Kursash/internal/models"
	"context"
	"fmt"
	"time"
)

func (d *DB) AddCourse(ctx context.Context, course models.CourseModel) error {
	stmt := `
    INSERT INTO courses
        (title, description, modules)
    VALUES
        ($1, $2, $3)`

	cont, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := d.findCourse(cont, course.Title)
	if err == nil {
		return fmt.Errorf("CourseDB: AddCourse: Course already exists")
	}

	_, err = d.Client.Exec(cont, stmt, course.Title, course.Description)
	if err != nil {
		return fmt.Errorf("Error in CourseDB: AddCourse: %s", err.Error())
	}

	return nil
}

func (d *DB) GetCourse(ctx context.Context, title string) (*models.CourseModel, error) {
	stmt := `SELECT * FROM courses WHERE title = $1`
	cont, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := d.findCourse(cont, title)
	if err != nil {
		return nil, err
	}

	var courseItem models.CourseModel
	err = d.Client.QueryRow(cont, stmt, title).Scan(&courseItem.Title, &courseItem.Description)
	if err != nil {
		return nil, fmt.Errorf("Error in CourseDB: GetCourse: %s", err.Error())
	}
	return &courseItem, nil
}

func (d *DB) ChangeCourse(ctx context.Context, title string, newDescription string) error {
	stmt := `UPDATE courses SET description = $2 WHERE title = $1`
	cont, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := d.findCourse(cont, title)
	if err != nil {
		return err
	}

	_, err = d.Client.Exec(cont, stmt, title, newDescription)
	if err != nil {
		return fmt.Errorf("Error in CourseDB: ChangeCourse: %s", err.Error())
	}
	return nil
}

func (d *DB) DeleteCourse(ctx context.Context, title string) error {
	stmt := `DELETE FROM courses WHERE title = $1`
	cont, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := d.findCourse(cont, title)
	if err != nil {
		return err
	}

	_, err = d.Client.Exec(cont, stmt, title)
	if err != nil {
		return fmt.Errorf("Error in CourseDB: DeleteCourse: %s", err.Error())
	}
	return nil
}

func (d *DB) GetCourseList(ctx context.Context) ([]*models.CourseModel, error) {
	stmt := `SELECT * FROM courses`
	cont, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	rows, err := d.Client.Query(cont, stmt)
	if err != nil {
		return nil, fmt.Errorf("Error in CourseDB: GetCourseList: %s", err.Error())
	}
	defer rows.Close()

	var courses []*models.CourseModel
	for rows.Next() {
		var courseItem models.CourseModel
		err = rows.Scan(&courseItem.Title, &courseItem.Description)
		if err != nil {
			return nil, fmt.Errorf("Error in CourseDB: GetCourseList: %s", err.Error())
		}
		courses = append(courses, &courseItem)
	}
	return courses, nil
}

func (d *DB) findCourse(ctx context.Context, title string) error {
	stmt := `SELECT * FROM courses WHERE title = $1`
	var courseItem models.CourseModel
	cont, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	err := d.Client.QueryRow(cont, stmt, title).Scan(&courseItem.Title, &courseItem.Description)
	if err != nil {
		return fmt.Errorf("Error in CourseDB: findCourse: %s", err.Error())
	}
	return nil
}
