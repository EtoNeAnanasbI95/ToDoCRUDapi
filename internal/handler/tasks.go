package handler

import (
	"fmt"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllTasks возвращает все задачи пользователя
// @Summary Получить задачи
// @Description Этот эндпоинт возвращает список всех задач, связанных с текущим пользователем
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Task
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 400 {object} map[string]string "Internal Server Error"
// @Router /api/tasks/ [get]
func (h *Handler) GetAllTasks(c *gin.Context) {
	uid, err := h.checkUserId(c)
	if err != nil {
		return
	}
	tasks, err := h.services.Tasks.GetAll(uid)
	if err != nil {
		h.log.Error("Get all tasks failed", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Get all tasks failed"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask получает информацию о задаче пользователя по ID
// @Summary Получить задаче пользователя
// @Description Этот эндпоинт возвращает информацию о задаче пользователя по его ID
// @Tags Tasks
// @Produce json
// @Param id path int true "ID задачи пользователя"
// @Security BearerAuth
// @Success 200 {object} models.Task
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string
// @Router /api/tasks/{id} [get]
func (h *Handler) GetTask(c *gin.Context) {
	uid, err := h.checkUserId(c)
	if err != nil {
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	task, err := h.services.Tasks.Get(uid, id)
	if err != nil {
		h.log.Error("Get task failed", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Get task failed"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// UpdateTask обновляет данные задачи пользователя
// @Summary Обновить задачу пользователя
// @Description Этот эндпоинт обновляет информацию о задаче пользователя по её ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID задачи пользователя"
// @Param task body models.TaskInput true "Обновленные данные задачи пользователя"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /api/tasks/{id} [put]
func (h *Handler) UpdateTask(c *gin.Context) {
	uid, err := h.checkUserId(c)
	if err != nil {
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}
	var task models.TaskInput
	if err = c.BindJSON(&task); err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.services.Tasks.Update(uid, id, &task)
	if err != nil {
		h.log.Error("Update task failed", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Update task failed"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Update task"})
}

// DeleteTask удаляет задачу пользователя по ID
// @Summary Удалить задачу пользователя
// @Description Этот эндпоинт удаляет задачу пользователя из системы по его ID
// @Tags Tasks
// @Param id path int true "ID задачи"
// @Security BearerAuth
// @Success 204 "Задача удалена"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /api/tasks/{id} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	uid, err := h.checkUserId(c)
	if err != nil {
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	err = h.services.Tasks.Delete(uid, id)
	if err != nil {
		h.log.Error("Delete task failed", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Delete task failed"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "successfully delete task"})
}

// CreateTask добавляет новую задачу пользователя
// @Summary Создать новую задачу пользователя
// @Description Этот эндпоинт создает задачу пользователя по данным, отправленным в теле запроса
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Данные задачи пользователя"
// @Security BearerAuth
// @Success 201 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /api/tasks/ [post]
func (h *Handler) CreateTask(c *gin.Context) {
	uid, err := h.checkUserId(c)
	if err != nil {
		return
	}
	task := &models.Task{}
	if err := c.ShouldBindJSON(task); err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.Tasks.Create(uid, task)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) checkUserId(c *gin.Context) (int, error) {
	uid, ok := c.Get("uid")
	if !ok {
		h.log.Error("Get all tasks failed, uid is empty")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "uid is empty"})
		return 0, fmt.Errorf("uid is empty")
	}
	return uid.(int), nil
}
