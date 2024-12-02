package models

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
