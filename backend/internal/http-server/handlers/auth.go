package handlers

import (
	"Kursash/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// SignUp
// @Summary      Регистрация пользователя
// @Description  Создание нового пользователя и отправление уведомления на почту
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body models.UserModel true "Данные пользователя"
// @Success      200 {object} map[string]interface{} "Пользователь успешно зарегистрирован"
// @Failure      400 {object} map[string]string "Некорректный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /auth/sign-up [post]
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

// SignIn
// @Summary      Авторизация пользователя
// @Description  Генерация токена для авторизованного пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body signInInput true "Данные для авторизации"
// @Success      200 {object} map[string]interface{} "Токен успешно сгенерирован"
// @Failure      400 {object} map[string]string "Некорректный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
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