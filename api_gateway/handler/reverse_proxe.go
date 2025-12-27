package handler

import (
	"api-gateway/config"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(ctx *gin.Context) {
	cfg := config.GetConfig()
	target := cfg.UserServiceHost + ":" + cfg.UserServicePort

	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = target
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)

}
