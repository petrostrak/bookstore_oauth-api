package app

import (
	"bookstore_oauth-api/src/domain/accesstoken"
	"bookstore_oauth-api/src/http"
	"bookstore_oauth-api/src/repository/db"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

// StartApp will start the application with a new Repository and handler
func StartApp() {
	// atService := accesstoken.NewService(db.NewRepository())
	atHandler := http.NewHandler(accesstoken.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}
