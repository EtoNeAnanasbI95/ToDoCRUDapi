package handler

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/service"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
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

func (h *UsersHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := h.s.Users.Get(id)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UsersHandler) GetAll(c *gin.Context) {
	users, err := h.s.Users.GetAll()
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UsersHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	var User models.UserInput
	if err = c.BindJSON(&User); err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.s.Users.Update(id, &User)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "successfully update user"})
}

func (h *UsersHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	err = h.s.Users.Delete(id)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully delete user"})
}
