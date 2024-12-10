package handlers

import (
	"Kursash/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// createCourse
// @Summary Создание образовательного курса
// @Description Создание нового курса пользователем
// @Tags courses
// @Accept json
// @Produce json
// @Param input body models.Course true "Данные курса"
// @Success 200 {object} map[string]interface{} "Курс успешно создан"
// @Failure 400 {object} errorResponse "Некорректный запрос"
// @Failure 500 {object} errorResponse "Ошибка сервера"
// @Security ApiKeyAuth
// @Router /api/courses/ [post]
func (h *Handler) createCourse(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Course
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Profession == "" {
		input.Profession = "Неизвестно"
	}

	id, err := h.services.Course.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []models.Course `json:"data"`
}

// getAllCourses
// @Summary Получение всех образовательных курсов
// @Description Возвращение списка всех курсов пользователя
// @Tags courses
// @Produce json
// @Success 200 {object} getAllListsResponse "Список курсов"
// @Failure 500 {object} errorResponse "Ошибка сервера"
// @Security ApiKeyAuth
// @Router /api/courses/ [get]
func (h *Handler) getAllCourses(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	courses, err := h.services.Course.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: courses,
	})
}

// getCoursesById
// @Summary Получение курса по ID
// @Description Возвращение данных курса по указанному ID
// @Tags courses
// @Produce json
// @Param id path int true "ID курса"
// @Success 200 {object} models.Course "Информация о курсе"
// @Failure 400 {object} errorResponse "Некорректный ID"
// @Failure 500 {object} errorResponse "Ошибка сервера"
// @Security ApiKeyAuth
// @Router /api/courses/{id} [get]
func (h *Handler) getCoursesById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.Course.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// updateCourses
// @Summary Обновление курса
// @Description Обновление данных курса
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID курса"
// @Param input body models.UpdateCourseInput true "Данные для обновления"
// @Success 200 {object} statusResponse "Курс успешно обновлен"
// @Failure 400 {object} errorResponse "Некорректный запрос"
// @Failure 500 {object} errorResponse "Ошибка сервера"
// @Security ApiKeyAuth
// @Router /api/courses/{id} [put]
func (h *Handler) updateCourses(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateCourseInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Profession != nil && *input.Profession == "" {
		*input.Profession = "Неизвестно"
	}

	if err := h.services.Course.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// deleteCourses
// @Summary Удаление курса
// @Description Удаление курса по указанному ID
// @Tags courses
// @Produce json
// @Param id path int true "ID курса"
// @Success 200 {object} statusResponse "Курс успешно удален"
// @Failure 400 {object} errorResponse "Некорректный ID"
// @Failure 500 {object} errorResponse "Ошибка сервера"
// @Security ApiKeyAuth
// @Router /api/courses/{id} [delete]
func (h *Handler) deleteCourses(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Course.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// getCoursesByProfession
// @Summary Получение курсов по направлению
// @Description Возвращение списка курсов по указанному направлению
// @Tags courses
// @Produce json
// @Param profession query string true "Направление курсов"
// @Success 200 {object} coursesResponse "Список курсов"
// @Failure 400 {object} errorResponse "Направление не указано"
// @Failure 500 {object} errorResponse "Ошибка сервера"
// @Security ApiKeyAuth
// @Router /api/courses/profession [get]
func (h *Handler) getCoursesByProfession(c *gin.Context) {
	profession := c.Query("profession")
	if profession == "" {
		newErrorResponse(c, http.StatusBadRequest, "profession is required")
		return
	}

	courses, err := h.services.Course.GetByProfession(profession)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, coursesResponse{Data: courses})
}

// addUserCourse
// @Summary Добавление курса пользователя
// @Description Добавление курса пользователя в список обучающегося
// @Tags user
// @Accept json
// @Produce json
// @Param input body map[string]int true "ID курса"
// @Success 200 {object} successResponse "Курс успешно добавлен"
// @Failure 400 {object} errorResponse "Некорректный запрос"
// @Failure 401 {object} errorResponse "Пользователь не авторизован"
// @Failure 500 {object} errorResponse "Ошибка сервера"
// @Security ApiKeyAuth
// @Router /api/user/courses [post]
func (h *Handler) addUserCourse(c *gin.Context) {
	var input struct {
		CourseID int `json:"course_id"`
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	err = h.services.Course.AddUserCourse(userId, input.CourseID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, successResponse{Message: "course added successfully"})
}

// getUserCourses
// @Summary Получение курсов пользователя
// @Description Возвращение списка курсов пользователя
// @Tags user
// @Produce json
// @Success 200 {object} getAllListsResponse "Список курсов"
// @Failure 401 {object} errorResponse "Пользователь не авторизован"
// @Failure 500 {object} errorResponse "Ошибка сервера"
// @Security ApiKeyAuth
// @Router /api/user/courses [get]
func (h *Handler) getUserCourses(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "user not authorized")
		return
	}

	courses, err := h.services.Course.GetUserCourses(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: courses,
	})
}