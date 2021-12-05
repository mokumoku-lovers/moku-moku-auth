package app

import (
	"github.com/gin-gonic/gin"
	"moku-moku/domain/access_token"
	"moku-moku/http"
	"moku-moku/repository/db"
	"moku-moku/repository/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(
		access_token.NewService(
			rest.NewUsersRepository(),
			db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/login", atHandler.UserLogin)

	router.Run(":9001")
}
