package handler

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(w gin.ResponseWriter, r *http.Request, host string, port string) {
	target := host + ":" + port

	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = target
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}
