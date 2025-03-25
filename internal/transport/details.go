package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      Get info about car
// @Description  Get info about one car by id
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Success      200  {object}  successResponse
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure     500  {object}  errorResponse
// @Router       /details/models [get]
func (h *Handler) GetAllCities(c *gin.Context) {
	ctx := c.Request.Context()
	cities, err := h.repos.Details.GetAllCities(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the cities", cities)
}

func (h *Handler) GetAllBrands(c *gin.Context) {
	ctx := c.Request.Context()
	brands, err := h.repos.Details.GetAllBrands(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the brands", brands)
}

func (h *Handler) GetAllModels(c *gin.Context) {
	ctx := c.Request.Context()
	models, err := h.repos.Details.GetAllModels(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the models", models)
}

func (h *Handler) GetAllCategories(c *gin.Context) {
	ctx := c.Request.Context()
	categories, err := h.repos.Details.GetAllCategories(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the categories", categories)
}

func (h *Handler) GetAllBodies(c *gin.Context) {
	ctx := c.Request.Context()
	bodies, err := h.repos.Details.GetAllBodies(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the bodies", bodies)
}

func (h *Handler) GetAllGenerations(c *gin.Context) {
	ctx := c.Request.Context()
	generations, err := h.repos.Details.GetAllGenerations(ctx)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully got the generations", generations)
}
