package app

import (
	"github.com/gin-gonic/gin"
	"moku-moku/domain/access_token"
	"moku-moku/http"
	"moku-moku/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":9001")
}
