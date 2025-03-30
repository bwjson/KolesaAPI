package transport

import (
	"github.com/bwjson/kolesa_api/internal/repository"
	"github.com/bwjson/kolesa_api/internal/service"
	"github.com/bwjson/kolesa_api/pkg"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Services
	repos    *repository.Repos
	s3       *pkg.S3Client
}

func NewHandler(services *service.Services, repo *repository.Repos, s3 *pkg.S3Client) *Handler {
	return &Handler{services: services, repos: repo, s3: s3}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//}))

	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	}

	return r
}
