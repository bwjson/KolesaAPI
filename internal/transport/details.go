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

	NewSuccessResponse(c, http.StatusOK, cities)
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

	NewSuccessResponse(c, http.StatusOK, brands)
}

// @Summary      Get all models
// @Tags         details
// @Produce      json
// @Success      200  {object}  successResponse
// @Failure      500  {object}  errorResponse
// @Router       /details/models [get]
func (h *Handler) GetAllModels(c *gin.Context) {
	ctx := c.Request.Context()

	brandSource := c.DefaultQuery("brand", "")
	models, err := h.repos.Details.GetAllModels(ctx, brandSource)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, models)
}

// @Summary      Get all generations
// @Tags         details
// @Produce      json
// @Success      200  {object}  successResponse
// @Failure      500  {object}  errorResponse
// @Router       /details/generations [get]
func (h *Handler) GetAllGenerations(c *gin.Context) {
	ctx := c.Request.Context()

	modelSource := c.DefaultQuery("model", "")
	generations, err := h.repos.Details.GetAllGenerations(ctx, modelSource)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, generations)
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

	NewSuccessResponse(c, http.StatusOK, categories)
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

	NewSuccessResponse(c, http.StatusOK, bodies)
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

	NewSuccessResponse(c, http.StatusOK, colors)
}
