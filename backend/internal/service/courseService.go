package service

import (
	"Kursash/internal/models"
	"context"
	"strings"
)

func (s *Service) AddCourse(ctx context.Context, course models.CourseModel) error {
	course.Title = strings.TrimSpace(course.Title)
	course.Description = strings.TrimSpace(course.Description)
	return s.repo.AddCourse(ctx, course)
}

func (s *Service) GetCourse(ctx context.Context, title string) (*models.CourseModel, error) {
	title = strings.TrimSpace(title)
	return s.repo.GetCourse(ctx, title)
}

func (s *Service) GetCoursesList(ctx context.Context) ([]*models.CourseModel, error) {
	return s.repo.GetCourseList(ctx)
}

func (s *Service) DeleteCourse(ctx context.Context, title string) error {
	title = strings.TrimSpace(title)
	return s.repo.DeleteCourse(ctx, title)
}

func (s *Service) ChangeCourse(ctx context.Context, title, description string) error {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)
	return s.repo.ChangeCourse(ctx, title, description)
}
