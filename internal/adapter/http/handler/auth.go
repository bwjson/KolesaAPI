package handler

import (
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/response"
	"github.com/bwjson/kolesa_api/internal/models"
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

// @Summary      Send request code to phone number
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body     CodeRequest true "Номер телефона"
// @Success      200     {object} response.successResponse
// @Failure      400     {object} response.errorResponse "Некорректный формат запроса"
// @Failure      500     {object} response.errorResponse "Ошибка при отправке кода"
// @Router       /auth/request_code [post]
func (h *Handler) RequestCode(c *gin.Context) {
	var request CodeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// gRPC
	err := h.gRPC.SendVerificationCode(request.PhoneNumber)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]interface{}{}

	response.NewSuccessResponse(c, http.StatusOK, data)
}

// @Summary      Verify code
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body     VerifyRequest true "Номер телефона и код"
// @Success      200     {object} map[string]string "access_token и refresh_token"
// @Failure      400     {object} response.errorResponse "Некорректный формат запроса"
// @Failure      500     {object} response.errorResponse "Ошибка верификации или создания пользователя"
// @Router       /auth/verify_code [post]
func (h *Handler) VerifyCode(c *gin.Context) {
	var request VerifyRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// gRPC
	accessToken, refreshToken, err := h.gRPC.VerifyCode(request.PhoneNumber, request.Code)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tokens := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	newUser := models.User{
		PhoneNumber: request.PhoneNumber,
		Email:       nil,
	}

	// check existing user
	var user *models.User
	user, err = h.services.Users.GetByPhoneNumber(c.Request.Context(), newUser.PhoneNumber)

	if user == nil {
		_, err := h.services.Users.Create(c.Request.Context(), newUser)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, "Cannot create new user")
			return
		}
	}

	response.NewSuccessResponse(c, http.StatusOK, tokens)
}

// @Summary      Refresh access token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body     RefreshRequest true "Refresh токен"
// @Success      200     {object} map[string]string "access_token"
// @Failure      400     {object} response.errorResponse "Некорректный формат запроса"
// @Failure      500     {object} response.errorResponse "Ошибка обновления токена"
// @Router       /auth/refresh [post]
func (h *Handler) RefreshAccessToken(c *gin.Context) {
	var refreshToken RefreshRequest

	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, err := h.gRPC.RefreshAccessToken(refreshToken.RefreshToken)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, accessToken)
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
