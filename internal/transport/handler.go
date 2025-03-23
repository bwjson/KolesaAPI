package transport

import (
	"github.com/bwjson/api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	cars := r.Group("/api/cars")
	{
		// r.Use(AuthMiddleware)
		cars.POST("/", h.Create)
		cars.GET("/", h.GetAll)
		cars.GET("/{id}", h.GetById)
		cars.PATCH("/{id}", h.UpdateById)
		cars.DELETE("/{id}", h.DeleteById)
	}

	return r
}
