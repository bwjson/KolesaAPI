package transport

import (
	"github.com/bwjson/api/internal/service"
	"github.com/bwjson/api/pkg"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Services
	s3       *pkg.S3Client
}

func NewHandler(services *service.Services, s3 *pkg.S3Client) *Handler {
	return &Handler{services: services, s3: s3}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	cars := r.Group("/api/cars")
	{
		// r.Use(AuthMiddleware)
		cars.POST("/", h.Create)
		cars.GET("/", h.GetAll)
		cars.GET("/extended", h.GetAllExtended)
		cars.GET("/:id", h.GetById)
		cars.PATCH("/:id", h.UpdateById)
		//cars.DELETE("/:id", h.DeleteById)
		cars.GET("/photo/:file_id", h.GetAvatar)
	}

	return r
}
