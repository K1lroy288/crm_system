package main

import (
	"api-gateway/config"
	"api-gateway/handler"
	"api-gateway/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Api Gateway is up!")
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "authentication.html", nil)
	})

	r.GET("/table", func(ctx *gin.Context) {
		_, err := utils.ValidateJWT(ctx)
		if err != nil {
			log.Printf("Invalid token: %v", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}

		ctx.HTML(http.StatusOK, "table.html", nil)
	})

	api := r.Group("/auth")
	{
		api.POST("/register", func(ctx *gin.Context) {
			handler.ReverseProxy(ctx.Writer, ctx.Request, cfg.UserServiceHost, cfg.UserServicePort)
		})

		api.POST("/login", func(ctx *gin.Context) {
			handler.ReverseProxy(ctx.Writer, ctx.Request, cfg.UserServiceHost, cfg.UserServicePort)
		})
	}

	api2 := r.Group("/user")
	{
		api2.GET("/masters", func(ctx *gin.Context) {
			_, err := utils.ValidateJWT(ctx)
			if err != nil {
				log.Printf("Invalid token: %v", err)
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}

			handler.ReverseProxy(ctx.Writer, ctx.Request, cfg.UserServiceHost, cfg.UserServicePort)
		})
	}

	api3 := r.Group("/visit")
	{
		api3.GET("/visits", func(ctx *gin.Context) {
			_, err := utils.ValidateJWT(ctx)
			if err != nil {
				log.Printf("Invalid token: %v", err)
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}

			handler.ReverseProxy(ctx.Writer, ctx.Request, cfg.VisitServiceHost, cfg.VisitServicePort)
		})

		api3.POST("/visits", func(ctx *gin.Context) {
			_, err := utils.ValidateJWT(ctx)
			if err != nil {
				log.Printf("Invalid token: %v", err)
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}

			handler.ReverseProxy(ctx.Writer, ctx.Request, cfg.VisitServiceHost, cfg.VisitServicePort)
		})

		api3.DELETE("/visits/:id", func(ctx *gin.Context) {
			_, err := utils.ValidateJWT(ctx)
			if err != nil {
				log.Printf("Invalid token: %v", err)
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}

			handler.ReverseProxy(ctx.Writer, ctx.Request, cfg.VisitServiceHost, cfg.VisitServicePort)
		})

		api3.PUT("/visits/:id", func(ctx *gin.Context) {
			_, err := utils.ValidateJWT(ctx)
			if err != nil {
				log.Printf("Invalid token: %v", err)
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}

			handler.ReverseProxy(ctx.Writer, ctx.Request, cfg.VisitServiceHost, cfg.VisitServicePort)
		})
	}

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	r.Run(addr)
}
