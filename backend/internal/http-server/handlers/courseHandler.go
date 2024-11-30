package handlers

import (
	"Kursash/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      Add a new course
// @Description  Adds a new course to the storage.
// @Tags         Courses
// @Accept       json
// @Produce      json
// @Param        course  body      models.CourseModel  true  "Course to add"
// @Success      200     {object}  string    "status: ok"
// @Failure      400     {object}  string    "error: Failed to bind course"
// @Failure      500     {object}  string    "error: Failed to add course"
// @Router       /course [post]

func (h *Handler) AddCourse(c *gin.Context) {
	var course models.CourseModel
	if err := c.BindJSON(&course); err != nil {
		h.logger.Error("Failed to bind course", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind course"})
		return
	}
	if err := h.service.AddCourse(c, course); err != nil {
		h.logger.Error("Failed to add course", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add course"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @Summary      Get list of courses
// @Description  Retrieves the list of all courses in the storage.
// @Tags         Courses
// @Produce      json
// @Success      200  {array}   models.CourseModel  "List of courses"
// @Failure      500  {object}  string    "error: Failed to get courses list"
// @Router       /course [get]

func (h *Handler) GetCoursesList(c *gin.Context) {
	courses, err := h.service.GetCoursesList(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get courses list", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get courses list"})
		return
	}
	if courses == nil {
		c.JSON(http.StatusOK, gin.H{"status": "empty"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

// @Summary      Update a course
// @Description  Updates the details of an existing course.
// @Tags         Courses
// @Accept       json
// @Produce      json
// @Param        course  body      models.CourseModel  true  "Course to update"
// @Success      200     {object}  string    "status: ok"
// @Failure      400     {object}  string    "error: Failed to bind course"
// @Failure      500     {object}  string    "error: Failed to change course"
// @Router       /course [put]

func (h *Handler) ChangeCourse(c *gin.Context) {
	var course models.CourseModel
	if err := c.BindJSON(&course); err != nil {
		h.logger.Error("Failed to bind course", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind course"})
		return
	}
	if err := h.service.ChangeCourse(c, course.Title, course.Description); err != nil {
		h.logger.Error("Failed to change course", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change course"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @Summary      Delete a course
// @Description  Deletes a course from the storage by title.
// @Tags         Courses
// @Produce      json
// @Param        title  query     string  true  "Course title"
// @Success      200   {object}  string  "status: ok"
// @Failure      400   {object}  string  "error: Failed to get title"
// @Failure      500   {object}  string  "error: Failed to delete course"
// @Router       /course [delete]

func (h *Handler) DeleteCourse(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		h.logger.Error("Failed to get title", "error", "empty title")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get title"})
		return
	}
	if err := h.service.DeleteCourse(c, title); err != nil {
		h.logger.Error("Failed to delete course", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// @Summary      Get a course by title
// @Description  Retrieves a specific course from the storage by its title.
// @Tags         Courses
// @Produce      json
// @Param        title  query     string  true  "Course title"
// @Success      200    {object}  models.CourseModel  "Course details"
// @Failure      400    {object}  string    "error: Missing title"
// @Failure      404    {object}  string    "error: Course not found"
// @Failure      500    {object}  string    "error: Failed to get course"
// @Router       /course/{title} [get]

func (h *Handler) GetCourse(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		h.logger.Error("Failed to get title", "error", "empty title")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get title"})
		return
	}

	course, err := h.service.GetCourse(c.Request.Context(), title)
	if err != nil {
		h.logger.Error("Failed to get course", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get course"})
		return
	}

	if course == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}
