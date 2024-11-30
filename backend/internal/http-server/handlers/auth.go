package handlers

import (
	"Kursash/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.UserModel

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Error("Error binding JSON", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateUser(input)
	if err != nil {
		h.logger.Error("Ошибка создания пользователя", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Info("Пользователь успешно создан", "id", id)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
