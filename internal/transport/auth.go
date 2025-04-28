package transport

import (
	"github.com/bwjson/kolesa_api/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VerifyRequest struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}

type CodeRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) RequestCode(c *gin.Context) {
	var request CodeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// gRPC
	err := h.gRPC.SendVerificationCode(request.PhoneNumber)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]interface{}{}

	NewSuccessResponse(c, http.StatusOK, data)
}

func (h *Handler) VerifyCode(c *gin.Context) {
	var request VerifyRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// gRPC
	accessToken, refreshToken, err := h.gRPC.VerifyCode(request.PhoneNumber, request.Code)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tokens := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	newUser := dto.User{
		PhoneNumber: request.PhoneNumber,
	}

	// check existing user
	var user *dto.User
	user, err = h.repos.Users.GetByPhoneNumber(c.Request.Context(), newUser.PhoneNumber)

	if &user == nil {
		_, err := h.repos.Users.Create(c.Request.Context(), newUser)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, "Cannot create new user")
			return
		}
	}

	NewSuccessResponse(c, http.StatusOK, tokens)
}

func (h *Handler) RefreshAccessToken(c *gin.Context) {
	var refreshToken RefreshRequest

	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, err := h.gRPC.RefreshAccessToken(refreshToken.RefreshToken)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, accessToken)
}

//func (h *Handler) AuthMiddleware(c *gin.Context) {
//	accessTokenString := c.Request.Header.Get("accessToken")
//	if accessTokenString == "" {
//		NewErrorResponse(c, http.StatusUnauthorized, "No access token")
//		return
//	}
//
//	// Bearer + AccessToken
//	accessTokenParts := strings.Split(accessTokenString, " ")
//	if len(accessTokenParts) != 2 {
//		NewErrorResponse(c, http.StatusUnauthorized, "Invalid access token")
//		return
//	}
//
//	_, err := jwt.Parse(accessTokenParts[1], func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//		// change this later
//		return []byte(os.Getenv("JWTSecret")), nil
//	})
//
//	//if err != nil {
//	//	return "", err
//	//}
//
//	// TODO: add signature validation
//}
