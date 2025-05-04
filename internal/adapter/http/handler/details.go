package handler

import (
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      Get all cities
// @Tags         details
// @Produce      json
// @Success      200  {object}  response.successResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /details/cities [get]
func (h *Handler) GetAllCities(c *gin.Context) {
	ctx := c.Request.Context()
	cities, err := h.services.Details.GetAllCities(ctx)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, cities)
}

// @Summary      Get all brands
// @Tags         details
// @Produce      json
// @Success      200  {object}  response.successResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /details/brands [get]
func (h *Handler) GetAllBrands(c *gin.Context) {
	ctx := c.Request.Context()

	brands, err := h.services.Details.GetAllBrands(ctx)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, brands)
}

// @Summary      Get all models
// @Tags         details
// @Produce      json
// @Param        brand  query string  false  "Brand filter"
// @Success      200  {object}  response.successResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /details/models [get]
func (h *Handler) GetAllModels(c *gin.Context) {
	ctx := c.Request.Context()

	brandSource := c.DefaultQuery("brand", "")
	models, err := h.services.Details.GetAllModels(ctx, brandSource)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, models)
}

// @Summary      Get all generations
// @Tags         details
// @Produce      json
// @Param        model  query string  false  "Model filter"
// @Success      200  {object}  response.successResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /details/generations [get]
func (h *Handler) GetAllGenerations(c *gin.Context) {
	ctx := c.Request.Context()

	modelSource := c.DefaultQuery("model", "")
	generations, err := h.services.Details.GetAllGenerations(ctx, modelSource)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, generations)
}

// @Summary      Get all categories
// @Tags         details
// @Produce      json
// @Success      200  {object}  response.successResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /details/categories [get]
func (h *Handler) GetAllCategories(c *gin.Context) {
	ctx := c.Request.Context()
	categories, err := h.services.Details.GetAllCategories(ctx)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, categories)
}

// @Summary      Get all bodies
// @Tags         details
// @Produce      json
// @Success      200  {object}  response.successResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /details/bodies [get]
func (h *Handler) GetAllBodies(c *gin.Context) {
	ctx := c.Request.Context()
	bodies, err := h.services.Details.GetAllBodies(ctx)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, bodies)
}

// @Summary      Get all colors
// @Tags         details
// @Produce      json
// @Success      200  {object}  response.successResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /details/colors [get]
func (h *Handler) GetAllColors(c *gin.Context) {
	ctx := c.Request.Context()
	colors, err := h.services.Details.GetAllColors(ctx)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, colors)
}
