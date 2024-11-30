package repository

import (
	"Kursash/internal/models"
	"context"
)

type courseStorage interface {
	AddCourse(ctx context.Context, course models.CourseModel) error
	GetCourse(ctx context.Context, title string) (*models.CourseModel, error)
	ChangeCourse(ctx context.Context, title string, newDescription string) error
	DeleteCourse(ctx context.Context, title string) error
	GetCourseList(ctx context.Context) ([]*models.CourseModel, error)
}
