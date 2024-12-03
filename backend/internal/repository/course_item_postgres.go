package repository

import (
	"Kursash/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CourseItemPostgres struct {
	db *sqlx.DB
}

func NewCourseItemPostgres(db *sqlx.DB) *CourseItemPostgres {
	return &CourseItemPostgres{db: db}
}

func (r *CourseItemPostgres) Create(listId int, item models.CourseItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description, image_url, course_url) values ($1, $2, $3, $4) RETURNING id", courseItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description, item.Image_url, item.Course_url)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *CourseItemPostgres) GetAll(userId, listId int) ([]models.CourseItem, error) {
	var items []models.CourseItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.umage_url, ti.course_url  FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		courseItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}
