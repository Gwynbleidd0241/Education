package models

type CourseModel struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Modules     []string `json:"modules" binding:"required"`
}
