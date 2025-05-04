package handler

import (
	"errors"
	"fmt"
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/dto"
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/response"
	"github.com/bwjson/kolesa_api/internal/service"
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
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        limit           query int     false "Limit param"
// @Param        offset          query int     false "Offset param"
// @Param        category        query string  false "Category filter"
// @Param        brand           query string  false "Brand filter"
// @Param        model           query string  false "Model filter"
// @Param        generation      query string  false "Generation filter"
// @Param        city            query string  false "City filter"
// @Param        color           query string  false "Color filter"
// @Param        body            query string  false "Body type filter"
// @Param        price_start     query int     false "Price start filter"
// @Param        price_end       query int     false "Price end filter"
// @Param        engine_start    query number  false "Engine volume start filter"
// @Param        engine_end      query number  false "Engine volume end filter"
// @Param        mileage_start   query int     false "Mileage start filter"
// @Param        mileage_end     query int     false "Mileage end filter"
// @Param        steering_wheel  query string  false "Steering wheel side filter"
// @Param        wheel_drive     query string  false "Wheel drive filter"
// @Success      200  {object}  response.successResponse
// @Failure      400  {object}  response.errorResponse
// @Failure      404  {object}  response.errorResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /cars/main [get]
func (h *Handler) GetAllCars(c *gin.Context) {
	ctx := c.Request.Context()
	filters := make(map[string]interface{})

	// pagination
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "9"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid limit param")
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid offset param")
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
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	priceEnd, err := parseQueryParam("price_end", c, 0)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	engineStart, err := parseQueryParam("engine_start", c, 0.0)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	engineEnd, err := parseQueryParam("engine_end", c, 0.0)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mileageStart, err := parseQueryParam("mileage_start", c, 0)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mileageEnd, err := parseQueryParam("mileage_end", c, 0)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
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
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]interface{}{
		"cars":        cars,
		"total_count": total_count,
	}

	response.NewSuccessResponse(c, http.StatusOK, data)
}

// @Summary      Get car by ID
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Success      200  {object}  response.successResponse
// @Failure      400  {object}  response.errorResponse
// @Failure      404  {object}  response.errorResponse
// @Failure     500  {object}  response.errorResponse
// @Router       /cars/{id} [get]
func (h *Handler) GetCarById(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid car id param")
		return
	}

	car, err := h.services.Cars.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, car)
}

// @Summary      Search cars
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        limit query int false "Limit param"
// @Param        offset query int false "Offset param"
// @Param        q query string false "Search query"
// @Success      200  {object}  response.successResponse
// @Failure      400  {object}  response.errorResponse
// @Failure      404  {object}  response.errorResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /cars/search [get]
func (h *Handler) SearchCars(c *gin.Context) {
	ctx := c.Request.Context()

	query := c.Query("q")

	// pagination
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "9"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid limit param")
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid offset param")
		return
	}

	cars, totalCount, err := h.services.Cars.SearchCars(ctx, query, limit, offset)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"cars":        cars,
		"total_count": totalCount,
	}

	response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *Handler) Create(c *gin.Context) {
	var car dto.CreateCarDTO

	ctx := c.Request.Context()

	if err := c.BindJSON(&car); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Cars.Create(ctx, car)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]interface{}{
		"id": id,
	}

	response.NewSuccessResponse(c, http.StatusOK, data)
}

func (h *Handler) UpdateById(c *gin.Context) {}

func (h *Handler) DeleteById(c *gin.Context) {}
