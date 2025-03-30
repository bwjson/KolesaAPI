package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      Get all cities
// @Tags         details
// @Produce      json
// @Success      200  {object}  successResponse
// @Failure      500  {object}  errorResponse
// @Router       /details/cities [get]
func (h *Handler) GetAllCities(c *gin.Context) {
	ctx := c.Request.Context()
	cities, err := h.repos.Details.GetAllCities(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the cities", cities)
}

// @Summary      Get all brands
// @Tags         details
// @Produce      json
// @Success      200  {object}  successResponse
// @Failure      500  {object}  errorResponse
// @Router       /details/brands [get]
func (h *Handler) GetAllBrands(c *gin.Context) {
	ctx := c.Request.Context()
	brands, err := h.repos.Details.GetAllBrands(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the brands", brands)
}

// @Summary      Get all models
// @Tags         details
// @Produce      json
// @Success      200  {object}  successResponse
// @Failure      500  {object}  errorResponse
// @Router       /details/models [get]
func (h *Handler) GetAllModels(c *gin.Context) {
	ctx := c.Request.Context()
	models, err := h.repos.Details.GetAllModels(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the models", models)
}

// @Summary      Get all categories
// @Tags         details
// @Produce      json
// @Success      200  {object}  successResponse
// @Failure      500  {object}  errorResponse
// @Router       /details/categories [get]
func (h *Handler) GetAllCategories(c *gin.Context) {
	ctx := c.Request.Context()
	categories, err := h.repos.Details.GetAllCategories(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the categories", categories)
}

// @Summary      Get all bodies
// @Tags         details
// @Produce      json
// @Success      200  {object}  successResponse
// @Failure      500  {object}  errorResponse
// @Router       /details/bodies [get]
func (h *Handler) GetAllBodies(c *gin.Context) {
	ctx := c.Request.Context()
	bodies, err := h.repos.Details.GetAllBodies(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the bodies", bodies)
}

// @Summary      Get all generations
// @Tags         details
// @Produce      json
// @Success      200  {object}  successResponse
// @Failure      500  {object}  errorResponse
// @Router       /details/generations [get]
func (h *Handler) GetAllGenerations(c *gin.Context) {
	ctx := c.Request.Context()
	generations, err := h.repos.Details.GetAllGenerations(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the generations", generations)
}

// @Summary      Get all colors
// @Tags         details
// @Produce      json
// @Success      200  {object}  successResponse
// @Failure      500  {object}  errorResponse
// @Router       /details/colors [get]
func (h *Handler) GetAllColors(c *gin.Context) {
	ctx := c.Request.Context()
	colors, err := h.repos.Details.GetAllColors(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the generations", colors)
}
