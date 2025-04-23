package transport

import (
	"github.com/bwjson/kolesa_api/internal/grpc"
	"github.com/bwjson/kolesa_api/internal/repository"
	"github.com/bwjson/kolesa_api/internal/service"
	"github.com/bwjson/kolesa_api/pkg"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

type Handler struct {
	services *service.Services
	repos    *repository.Repos
	s3       *pkg.S3Client
	gRPC     *grpc.Client
}

func NewHandler(services *service.Services, repo *repository.Repos, s3 *pkg.S3Client, gRPC *grpc.Client) *Handler {
	return &Handler{services: services, repos: repo, s3: s3, gRPC: gRPC}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "https://car-market-zeta.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	gin.SetMode(gin.ReleaseMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		cars.GET("/extended", h.GetAllCarsExtended)
		cars.GET("/:id", h.GetCarById)
		cars.PATCH("/:id", h.UpdateById)
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
