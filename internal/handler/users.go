package handler

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUser(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.Users.Create(user)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := h.services.Users.Get(id)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.services.Users.GetAll()
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) UpdateUser(c *gin.Context) {
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
	err = h.services.Users.Update(id, &User)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "successfully update user"})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	err = h.services.Users.Delete(id)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "successfully delete user"})
}
