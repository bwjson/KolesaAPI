package handler

import (
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/response"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

//func (h *Handler) GetAvatarSource(c *gin.Context) {
//	ctx := c.Request.Context()
//	carId, err := strconv.Atoi(c.Param("car_id"))
//
//	_, err = h.services.Cars.GetById(ctx, carId)
//	if err != nil {
//		NewErrorResponse(c, http.StatusBadRequest, err.Error())
//	}
//
//	avatarSource, err := h.repos.Details.GetSourceById(ctx, carId)
//	if err != nil {
//		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
//	}
//
//	NewSuccessResponse(c, http.StatusOK, "Avatar", avatarSource)
//}

// @Summary      Get S3 auth token
// @Tags         s3
// @Produce		 json
// @Success      200  {object}  response.successResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /s3/auth_token [get]
func (h *Handler) GetAuthToken(c *gin.Context) {
	authResponse, err := h.s3.GetS3Credentials()
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	response.NewSuccessResponse(c, http.StatusOK, authResponse.AuthToken)
}

// @Summary      Upload file
// @Tags         s3
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "File to upload"
// @Success      200   {object}  response.successResponse
// @Failure      400   {object}  response.errorResponse
// @Failure      500   {object}  response.errorResponse
// @Router       /s3/upload_file [post]
func (h *Handler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// change header.Filename to custom unique filename

	url, err := h.s3.UploadFile(header.Filename, fileBytes)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, url)
}
