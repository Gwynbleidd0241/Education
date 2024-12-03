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
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/courses")
			{
				items.POST("/", h.createCourse)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("courses")
		{
			items.GET("/:id", h.getCourseById)
			items.PUT("/:id", h.updateCourse)
			items.DELETE("/:id", h.deleteCourse)
		}
	}
	return router
}
