package main

import (
	"go-wallet/config"
	"go-wallet/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api/v1")

	router.AuthRouter(api)

	r.Run("localhost:" + config.ENV.PORT)
}
