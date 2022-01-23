package app

import (
	"moku-moku/domain/access_token"
	"moku-moku/http_at"
	"moku-moku/repository/db"
	"moku-moku/repository/rest"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http_at.NewHandler(
		access_token.NewService(
			rest.NewUsersRepository(),
			db.NewRepository()))

	// Add cors middleware
	router.Use(cors.Default())

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/login", atHandler.UserLogin)

	//Swagger Documentation
	opts := middleware.RedocOpts{SpecURL: "./swagger.yml", Title: "Moku-Moku-Auth"}
	swg := middleware.Redoc(opts, nil)
	router.GET("/docs", gin.WrapH(swg))
	router.GET("/swagger.yml", gin.WrapH(http.FileServer(http.Dir("./"))))

	router.Run(":9001")
}
