package handlers

import (
	"Kursash/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(router *gin.Engine) *gin.Engine {

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		courses := api.Group("/courses")
		{
			courses.POST("/", h.createCourse)
			courses.GET("/", h.getAllCourses)
			courses.PUT("/:id", h.updateCourses)
			courses.DELETE("/:id", h.deleteCourses)
			courses.GET("/profession", h.getCoursesByProfession)
		}

		user := api.Group("/user")
		{
			user.POST("/courses", h.addUserCourse)
			user.GET("/courses", h.getUserCourses)
		}
	}
	return router
}
