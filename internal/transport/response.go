package transport

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type successResponse struct {
	Data interface{} `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.AbortWithStatusJSON(statusCode, successResponse{data})
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
