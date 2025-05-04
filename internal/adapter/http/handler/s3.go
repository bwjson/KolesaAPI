package handler

import (
	"encoding/base64"
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
// @Accept       application/json
// @Produce      json
// @Param        file  body      string  true  "Base64 image"
// @Success      200   {object}  response.successResponse
// @Failure      400   {object}  response.errorResponse
// @Failure      500   {object}  response.errorResponse
// @Router       /s3/upload_file [post]
func (h *Handler) UploadFile(c *gin.Context) {
	var base64Image string

	if err := c.BindJSON(&base64Image); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fileBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fileName := "TEST1"

	_, err = h.s3.UploadFile(fileName, fileBytes)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, fileName)
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
//func (h *Handler) UploadFile(c *gin.Context) {
//	file, _, err := c.Request.FormFile("file")
//	if err != nil {
//		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
//		return
//	}
//	defer file.Close()
//
//	fileBytes, err := io.ReadAll(file)
//	if err != nil {
//		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	filename := "test4"
//
//	fileId, err := h.s3.UploadFile(filename, fileBytes)
//	if err != nil {
//		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	data := map[string]interface{}{
//		"fileId": fileId,
//	}
//
//	response.NewSuccessResponse(c, http.StatusOK, data)
//}
