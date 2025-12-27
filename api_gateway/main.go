package main

import (
	"api-gateway/config"
	"api-gateway/handler"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Api Gateway is up!")
	})

	r.LoadHTMLGlob("templates/")

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Api Gateway is up!")
	})

	api := r.Group("/auth")
	{
		api.POST("/register", handler.ReverseProxy)

		api.POST("/login", handler.ReverseProxy)
	}

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	r.Run(addr)
}
