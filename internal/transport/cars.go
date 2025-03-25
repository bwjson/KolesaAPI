package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) Create(c *gin.Context) {
}

// @Summary      Get all cars
// @Description  Get all cars
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        limit query int false "Limit param"
// @Param        offset query int false "Offset param"
// @Success      200  {object}  successResponse
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /cars/extended [get]
func (h *Handler) GetAllCarsExtended(c *gin.Context) {
	ctx := c.Request.Context()

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid limit param")
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid offset param")
	}

	cars, total_count, err := h.services.Cars.GetAllExtended(ctx, limit, offset)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"cars":        cars,
		"total_count": total_count,
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully returned all the cars", data)
}

// @Summary      Main page cars
// @Description  Get basic info
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        limit query int false "Limit param"
// @Param        offset query int false "Offset param"
// @Success      200  {object}  successResponse
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /cars [get]
func (h *Handler) GetAllCars(c *gin.Context) {
	ctx := c.Request.Context()

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid limit param")
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid offset param")
	}

	cars, total_count, err := h.services.Cars.GetAll(ctx, limit, offset)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"cars":        cars,
		"total_count": total_count,
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully returned all the cars", data)
}

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
// @Router       /cars/{id} [get]
func (h *Handler) GetCarById(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid car id param")
	}

	car, err := h.services.Cars.GetById(ctx, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully returned a car", car)
}

func (h *Handler) UpdateById(c *gin.Context) {}

func (h *Handler) DeleteById(c *gin.Context) {

}
