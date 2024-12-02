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

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createCourse)
				items.GET("/", h.getAllCourses)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getCourseById)
			items.PUT("/:id", h.updateCourse)
			items.DELETE("/:id", h.deleteCourse)
		}
	}
	return router
}
