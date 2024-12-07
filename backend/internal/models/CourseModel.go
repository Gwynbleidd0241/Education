package models

import (
	"errors"
)

type Course struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
	Image_url   string `json:"image_url" db:"image_url" binding:"required"`
	Course_url  string `json:"course_url" db:"course_url" binding:"required"`
	Profession  string `json:"profession"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type UpdateCourseInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image_url   *string `json:"image_url"`
	Course_url  *string `json:"course_url"`
	Profession  *string `json:"profession"`
}

func (i UpdateCourseInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Image_url == nil && i.Course_url == nil && i.Profession == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
