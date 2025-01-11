package handler

import (
	_ "github.com/EtoNeAnanasbI95/ToDoCRUD/docs"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/clients/sso/grpc"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/service"
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	Sso      *grpc.Client
}

func NewHandler(log *slog.Logger, s *service.Service, sso *grpc.Client) *Handler {
	return &Handler{
		log:      log,
		services: s,
		Sso:      sso,
	}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	//corsConfig := cors.DefaultConfig()
	//corsConfig.AllowAllOrigins = true
	//corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	//h.log.Info("Allow headers", corsConfig.AllowHeaders)
	//router.Use(cors.New(corsConfig))
	router.Use(h.CORSMiddleware)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("", h.CreateUser)
			users.GET("", h.GetAllUsers)
			users.GET(":id", h.GetUser)
			users.PUT(":id", h.UpdateUser)
			users.DELETE(":id", h.DeleteUser)
		}
		// TODO: там косяк в свагере с защищённым ендпоинтом, надо сделать так, чтоб просил заголовок
		tasks := api.Group("/tasks", h.CheckAuth)
		{
			tasks.POST("", h.CreateTask)
			tasks.GET("", h.GetAllTasks)
			tasks.GET(":id", h.GetTask)
			tasks.PUT(":id", h.UpdateTask)
			tasks.DELETE(":id", h.DeleteTask)
		}
	}
	return router
}
