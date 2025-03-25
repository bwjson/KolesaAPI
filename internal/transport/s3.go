package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary      Get car avatar
// @Description  Get car avatar by fileId from S3
// @Tags         cars
// @Accept       json
// @Produce		 json
// @Param        car_id  path  string  true  "Car ID"
// @Success      200  {file}    successResponse
// @Failure      500  {object}  errorResponse
// @Router       /cars/avatar/{car_id} [get]
func (h *Handler) GetAvatarSource(c *gin.Context) {
	ctx := c.Request.Context()
	carId, err := strconv.Atoi(c.Param("car_id"))

	_, err = h.services.Cars.GetById(ctx, carId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	avatarSource, err := h.repos.Details.GetSourceById(ctx, carId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	NewSuccessResponse(c, http.StatusOK, "Avatar", avatarSource)
}

// @Summary      Get S3 auth token
// @Tags         s3
// @Produce		 json
// @Success      200  {object}  successResponse
// @Failure      500  {object}  errorResponse
// @Router       /s3/auth_token [get]
func (h *Handler) GetAuthToken(c *gin.Context) {
	authResponse, err := h.s3.GetS3Credentials()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	NewSuccessResponse(c, http.StatusOK, "Authorization token", authResponse.AuthToken)
}
