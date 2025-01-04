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
	users      CRUD
	tasks      CRUD
	usersTasks CRUD
}

func NewHandler(log *slog.Logger, s *service.Service) *Handler {
	return &Handler{
		users: NewUsersHandler(log, s),
	}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.users.Create)
			users.GET("/", h.users.GetAll)
			users.GET("/:id", h.users.Get)
			users.PUT("/:id", h.users.Update)
			users.DELETE("/:id", h.users.Delete)
		}
		// TODO: дописать роутинг для остальных ендпоинтов
		/*
			tasks := api.Group("/tasks")
			{
				tasks.POST("/", h.tasks.Create)
				tasks.GET("/:uid", h.tasks.GetAll)
				tasks.GET("/:uid/:id", h.tasks.Get)
				users.PUT("/:uid/:id", h.tasks.Update)
				tasks.DELETE("/:uid/:id", h.tasks.Delete)
			}
		*/
	}
	return router
}
