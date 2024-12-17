package handler

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/service"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type UsersHandler struct {
	log *slog.Logger
	s   *service.Service
}

func NewUsersHandler(log *slog.Logger, s *service.Service) *UsersHandler {
	return &UsersHandler{
		log: log,
		s:   s,
	}
}

func (h *UsersHandler) Create(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.s.Users.Create(user)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// TODO: завершить имплентацию хендлера
func (h *UsersHandler) Get(c *gin.Context) {
	panic("implement me")
}

// TODO: завершить имплентацию хендлера
func (h *UsersHandler) GetAll(c *gin.Context) {
	panic("implement me")
}

// TODO: завершить имплентацию хендлера
func (h *UsersHandler) Update(c *gin.Context) {
	panic("implement me")
}

// TODO: завершить имплентацию хендлера
func (h *UsersHandler) Delete(c *gin.Context) {
	panic("implement me")
}
