package handler

import (
	"go-wallet/dto"
	errorhandler "go-wallet/errorHandler"
	"go-wallet/helper"
	"go-wallet/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.Register

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	helper.Response(dto.Params{
		StatusCode: http.StatusCreated,
		Message:    "Register Success, Please Login",
	})

	c.JSON(http.StatusCreated, "")
}
