package handler

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUser добавляет нового пользователя
// @Summary Создать нового пользователя
// @Description Этот эндпоинт создает пользователя по данным, отправленным в теле запроса
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "Данные пользователя"
// @Success 201 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Router /api/users/ [post]
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
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

// GetUser получает информацию о пользователе по ID
// @Summary Получить пользователя
// @Description Этот эндпоинт возвращает информацию о пользователе по его ID
// @Tags Users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [get]
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
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetAllUsers возвращает список всех пользователей
// @Summary Получить всех пользователей
// @Description Этот эндпоинт возвращает список всех пользователей
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /api/users/ [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.services.Users.GetAll()
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUser обновляет данные пользователя
// @Summary Обновить пользователя
// @Description Этот эндпоинт обновляет информацию о пользователе по его ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param user body models.UserInput true "Обновленные данные пользователя"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [put]
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

// DeleteUser удаляет пользователя по ID
// @Summary Удалить пользователя
// @Description Этот эндпоинт удаляет пользователя из системы по его ID
// @Tags Users
// @Param id path int true "ID пользователя"
// @Success 204 "Пользователь удален"
// @Failure 400 {object} map[string]string
// @Router /api/users/{id} [delete]
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
	c.JSON(http.StatusNoContent, gin.H{"msg": "successfully delete user"})
}
