package http

import (
	"github.com/gin-gonic/gin"
	"moku-moku/domain/access_token"
	"net/http"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "To be implemented")
}
