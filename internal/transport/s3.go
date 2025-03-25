package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      Get car avatar
// @Description  Get car avatar by fileId from S3
// @Tags         cars
// @Produce		 image/png, image/jpeg
// @Param        fileId  path  string  true  "ID файла"
// @Success      200  {file}  binary
// @Failure      500  {object}  errorResponse
// @Router       /cars/photo/{file_id} [get]
func (h *Handler) GetAvatar(c *gin.Context) {
	fileId := c.Param("file_id")

	avatar, err := h.s3.DownloadFile("kolesa", fileId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	contentType := http.DetectContentType(avatar)

	c.Data(http.StatusOK, contentType, avatar)
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
