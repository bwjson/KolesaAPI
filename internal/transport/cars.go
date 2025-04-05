package transport

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func parseQueryParam(param string, c *gin.Context, defaultValue interface{}) (interface{}, error) {
	paramStr := c.DefaultQuery(param, "")
	if paramStr == "" {
		return defaultValue, nil
	}

	switch defaultValue.(type) {
	case int:
		value, err := strconv.Atoi(paramStr)
		if err != nil {
			return 0, fmt.Errorf("invalid %s param", param)
		}
		return value, nil
	case float64:
		value, err := strconv.ParseFloat(paramStr, 64)
		if err != nil {
			return 0.0, fmt.Errorf("invalid %s param", param)
		}
		return value, nil
	default:
		return nil, fmt.Errorf("unsupported default type for param %s", param)
	}
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
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid offset param")
		return
	}

	cars, total_count, err := h.services.Cars.GetAllExtended(ctx, limit, offset)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
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
// @Router       /cars/main [get]
func (h *Handler) GetAllCars(c *gin.Context) {
	ctx := c.Request.Context()
	filters := make(map[string]interface{})

	// pagination
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "9"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid limit param")
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid offset param")
		return
	}

	// join filters
	categorySource := c.DefaultQuery("category", "")
	brandSource := c.DefaultQuery("brand", "")
	modelSource := c.DefaultQuery("model", "")
	generationSource := c.DefaultQuery("generation", "")
	citySource := c.DefaultQuery("city", "")
	colorSource := c.DefaultQuery("color", "")
	bodySource := c.DefaultQuery("body", "")

	// in-model filters
	priceStart, err := parseQueryParam("price_start", c, 0)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	priceEnd, err := parseQueryParam("price_end", c, 0)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	engineStart, err := parseQueryParam("engine_start", c, 0.0)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	engineEnd, err := parseQueryParam("engine_end", c, 0.0)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mileageStart, err := parseQueryParam("mileage_start", c, 0)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mileageEnd, err := parseQueryParam("mileage_end", c, 0)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	steeringWheel := c.DefaultQuery("steering_wheel", "")
	wheelDrive := c.DefaultQuery("wheel_drive", "")

	filters["limit"] = limit
	filters["offset"] = offset
	filters["categorySource"] = categorySource
	filters["brandSource"] = brandSource
	filters["modelSource"] = modelSource
	filters["generationSource"] = generationSource
	filters["citySource"] = citySource
	filters["colorSource"] = colorSource
	filters["bodySource"] = bodySource
	filters["priceStart"] = priceStart
	filters["priceEnd"] = priceEnd
	filters["engineStart"] = engineStart
	filters["engineEnd"] = engineEnd
	filters["mileageStart"] = mileageStart
	filters["mileageEnd"] = mileageEnd
	filters["steeringWheel"] = steeringWheel
	filters["wheelDrive"] = wheelDrive

	cars, total_count, err := h.services.Cars.GetAll(ctx, filters) // add all others filters
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
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
		return
	}

	car, err := h.services.Cars.GetById(ctx, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, "Successfully returned a car", car)
}

func (h *Handler) UpdateById(c *gin.Context) {}

func (h *Handler) DeleteById(c *gin.Context) {}

func (h *Handler) Create(c *gin.Context) {}
