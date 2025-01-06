package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllTasks(c *gin.Context) {
	uid, ok := c.Get("uid")
	if !ok {
		h.log.Error("Get all tasks failed, uid is empty")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "uid is empty"})
		return
	}
	tasks, err := h.services.Tasks.GetAll(uid.(int))
	if err != nil {
		h.log.Error("Get all tasks failed", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Get all tasks failed"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTask(c *gin.Context) {
	panic("implement me")
}

func (h *Handler) UpdateTask(c *gin.Context) {
	panic("implement me")
}

func (h *Handler) DeleteTask(c *gin.Context) {
	panic("implement me")
}

func (h *Handler) CreateTask(c *gin.Context) {
	panic("implement me")
}
