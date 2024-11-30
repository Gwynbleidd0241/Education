package handlers

import (
	"Kursash/internal/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/go-ozzo/ozzo-log"
	"net/http"
)

type Handler struct {
	Router  *gin.Engine
	service *service.Service
	logger  *log.Logger
	authMW  *jwt.GinJWTMiddleware
	CourseHandler
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.Router.ServeHTTP(writer, request)
}

type CourseHandler interface {
	AddCourse(c *gin.Context)
	GetCoursesList(c *gin.Context)
	ChangeCourse(c *gin.Context)
	DeleteCourse(c *gin.Context)
	GetCourse(c *gin.Context)
}

func (h *Handler) registerRoutes() {
	auth := h.Router.Group("/auth")
	{
		auth.POST("/login", h.authMW.LoginHandler)
		auth.POST("/sign-up", h.SignUp)
	}

	authGroup := h.Router.Group("/course")
	authGroup.Use(h.authMW.MiddlewareFunc())
	{
		authGroup.POST("", h.AddCourse)
		authGroup.GET("", h.GetCoursesList)
		authGroup.PUT("/:title", h.ChangeCourse)
		authGroup.DELETE("/:title", h.DeleteCourse)
		authGroup.GET("/:title", h.GetCourse)
	}

	h.Router.OPTIONS("/course", h.Options)
}

func NewHandler(router *gin.Engine, serv *service.Service, logger *log.Logger, authMW *jwt.GinJWTMiddleware) *Handler {
	handler := &Handler{
		Router:  router,
		service: serv,
		logger:  logger,
		authMW:  authMW,
	}
	handler.registerRoutes()
	return handler
}

func (h *Handler) Options(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
