package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// CheckUserId проверяет айди авторизации
func (h *Handler) CheckUserId(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		h.log.Error("Authorization header is empty")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Authorization header is empty"})
		return
	}
	id, err := strconv.Atoi(header)
	if err != nil {
		h.log.Error("Authorization header is invalid")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}
	_, err = h.services.Users.Get(id)
	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}
	c.Set("uid", id)
}
