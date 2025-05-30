package handler

import (
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/metrics"
	"github.com/bwjson/kolesa_api/internal/grpc"
	"github.com/bwjson/kolesa_api/internal/service"
	"github.com/bwjson/kolesa_api/pkg/s3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
)

type Handler struct {
	log      *slog.Logger
	services *service.Services
	s3       *s3.S3Client
	gRPC     *grpc.Client
}

func NewHandler(log *slog.Logger, services *service.Services, s3 *s3.S3Client, gRPC *grpc.Client) *Handler {
	return &Handler{log: log, services: services, s3: s3, gRPC: gRPC}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(gin.Recovery())

	r.Use(metrics.PrometheusMiddleware())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://car-market-zeta.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8000/swagger/doc.json"),
	))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Your service is live 🎉",
		})
	})

	cars := r.Group("/api/cars")
	{
		cars.POST("/create", h.Create)
		cars.GET("/main", h.GetAllCars)
		cars.GET("/:id", h.GetCarById)
		cars.PATCH("/:id", h.UpdateById)
		cars.GET("/search", h.SearchCars)
	}

	details := r.Group("/api/details")
	{
		details.GET("/cities", h.GetAllCities)
		details.GET("/brands", h.GetAllBrands)
		details.GET("/models", h.GetAllModels)
		details.GET("/categories", h.GetAllCategories)
		details.GET("/bodies", h.GetAllBodies)
		details.GET("/generations", h.GetAllGenerations)
		details.GET("/colors", h.GetAllColors)
	}

	s3 := r.Group("/api/s3")
	{
		s3.GET("/auth_token", h.GetAuthToken)
		s3.POST("/upload_file", h.UploadFile)
	}

	auth := r.Group("/api/auth")
	{
		auth.POST("/request_code", h.RequestCode)
		auth.POST("/verify_code", h.VerifyCode)
		auth.POST("/refresh", h.RefreshAccessToken)
	}

	users := r.Group("/api/users")
	{
		users.POST("/create", h.CreateUser)
		users.GET("/get_all", h.GetUsers)
		users.GET("/:id", h.GetUserByID)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
	}

	return r
}
