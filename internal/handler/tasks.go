package handler

import (
	"fmt"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	c.JSON(http.StatusOK, gin.H{"msg": "successfully delete task"})
}

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
