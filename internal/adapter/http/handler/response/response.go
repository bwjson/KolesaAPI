package response

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type successResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewSuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	var response successResponse

	if statusCode == 200 {
		response = successResponse{
			Status: "ok",
			Data:   data,
		}
	}
	c.AbortWithStatusJSON(statusCode, response)
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{
		Status:  "error",
		Message: message})
}
