package handler

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type CRUD interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Handler struct {
	log      *slog.Logger
	services *service.Service
}

func NewHandler(log *slog.Logger, s *service.Service) *Handler {
	return &Handler{
		log:      log,
		services: s,
	}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.CreateUser)
			users.GET("/", h.GetAllUsers)
			users.GET("/:id", h.GetUser)
			users.PUT("/:id", h.UpdateUser)
			users.DELETE("/:id", h.DeleteUser)
		}
		tasks := api.Group("/tasks", h.CheckUserId)
		{
			tasks.POST("/", h.CreateTask)
			tasks.GET("/", h.GetAllTasks)
			tasks.GET("/:id", h.GetTask)
			users.PUT("/:id", h.UpdateTask)
			tasks.DELETE(":id", h.DeleteTask)
		}
	}
	return router
}
