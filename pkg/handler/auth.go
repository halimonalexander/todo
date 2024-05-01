package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/halimonalexander/todo"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id":       id,
		"name":     input.Name,
		"username": input.Username,
	})
}

func (h *Handler) singIn(c *gin.Context) {

}
