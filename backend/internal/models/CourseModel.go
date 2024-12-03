package models

import "errors"

type CourseList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
	Image_url   string `json:"image_url" db:"image_url" binding:"required"`
	Course_url  string `json:"course_url" db:"course_url" binding:"required"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image_url   *string `json:"image_url"`
	Course_url  *string `json:"course_url"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Image_url == nil && i.Course_url == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type CourseItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Image_url   string `json:"image_url" db:"image_url"`
	Course_url  string `json:"course_url" db:"course_url"`
}
