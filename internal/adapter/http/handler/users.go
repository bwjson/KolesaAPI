package handler

import (
	"fmt"
	"github.com/bwjson/kolesa_api/internal/adapter/http/handler/response"
	"github.com/bwjson/kolesa_api/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary      Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User data"
// @Success      201   {object}  response.successResponse
// @Failure      400   {object}  response.errorResponse
// @Failure      500   {object}  response.errorResponse
// @Router       /users/create [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	id, err := h.services.Users.Create(c.Request.Context(), user)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusCreated, id)
}

// @Summary      Get all users
// @Tags         users
// @Produce      json
// @Success      200  {object}  response.successResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /users/get_all [get]
func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.services.Users.GetAll(c.Request.Context())
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, users)
}

// @Summary      Get user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  response.successResponse
// @Failure      400  {object}  response.errorResponse
// @Failure      404  {object}  response.errorResponse
// @Router       /users/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.services.Users.GetByID(c.Request.Context(), id)
	if err != nil {
		response.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, user)
}

// @Summary      Update user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "User ID"
// @Param        user  body      models.User  true  "User data"
// @Success      200   {object}  response.successResponse
// @Failure      400   {object}  response.errorResponse
// @Failure      500   {object}  response.errorResponse
// @Router       /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	err = h.services.Users.Update(c.Request.Context(), id, user)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, id)
}

// @Summary      Delete user
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  response.successResponse
// @Failure      400  {object}  response.errorResponse
// @Failure      404  {object}  response.errorResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	_, err = h.services.Users.GetByID(c.Request.Context(), id)
	if err != nil {
		response.NewErrorResponse(c, http.StatusNotFound, fmt.Sprintf("There is no car with ID: %d", id))
		return
	}

	err = h.services.Users.Delete(c.Request.Context(), id)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusOK, nil)
}
