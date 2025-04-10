package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) RequestCode(c *gin.Context) {
	// TODO

	// gRPC
	err := h.gRPC.SendVerificationCode("+77783784148")
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]interface{}{}

	NewSuccessResponse(c, http.StatusOK, "Successfully sent the message", data)
}

func (h *Handler) VerifyCode(c *gin.Context) {
	// TODO

	// gRPC
	h.gRPC.VerifyCode("+77783784148", "")

	// Create user if doesn't exist
}

func (h *Handler) RefreshAccessToken(c *gin.Context) {
	// TODO
}
