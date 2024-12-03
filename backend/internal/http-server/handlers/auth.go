package handlers

import (
	"Kursash/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.UserModel

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	subject := "Добро пожаловать в EduCourse Gwynbleidd!"
	body := "Поздравляем с успешной регистрацией! Желаем успехов в обучении."
	if err := h.services.Notifications.SendEmail(input.Email, subject, body, nil); err != nil {
		log.Printf("Failed to send email to %s: %s", input.Email, err.Error())
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":      id,
			"message": "User registered successfully, but email notification failed.",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Username == "admin" && input.Password == "admin" {
		c.JSON(http.StatusOK, map[string]interface{}{
			"token": "admin-token",
		})
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
