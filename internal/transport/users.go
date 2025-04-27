package transport

import (
	"fmt"
	"github.com/bwjson/kolesa_api/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user dto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	id, err := h.repos.Users.Create(c.Request.Context(), user)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusCreated, id)
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.repos.Users.GetAll(c.Request.Context())
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, users)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.repos.Users.GetByID(c.Request.Context(), id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var user dto.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	err = h.repos.Users.Update(c.Request.Context(), id, user)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, id)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	_, err = h.repos.Users.GetByID(c.Request.Context(), id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, fmt.Sprintf("There is no car with ID: %d", id))
		return
	}

	err = h.repos.Users.Delete(c.Request.Context(), id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, nil)
}
