package handler

import (
	"context"
	authv1 "github.com/EtoNeAnanasbI95/protos_auth/gen/go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// CheckAuth занимается проверкой авторизации пользователя
func (h *Handler) CheckAuth(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		h.log.Error("Authorization header is empty")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Authorization header is empty"})
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		h.log.Error("invalid Authorization filed format")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "invalid Authorization filed format"})
		return
	}
	uid, err := h.Sso.Api.Validate(context.Background(), &authv1.TokenRequest{Token: headerParts[1]})
	if err != nil {
		h.log.Error("Error while validating token", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Error while validating token"})
		return
	}
	c.Set("uid", uid.GetUid())
	c.Next()
}

func (h *Handler) CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
