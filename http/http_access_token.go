package http

import (
	"github.com/gin-gonic/gin"
	"moku-moku/domain/access_token"
	"moku-moku/utils/errors"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
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
	at, err := h.service.GetByID(strings.TrimSpace(c.Param("access_token_id")))

	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, at)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var request access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	at, err := h.service.Create(request)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}
