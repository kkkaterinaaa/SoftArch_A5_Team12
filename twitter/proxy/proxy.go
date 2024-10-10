package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxy(target string) gin.HandlerFunc {
	targetURL, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()

	r.Any("/users/*any", proxy("http://localhost:8071"))

	r.Any("/users", proxy("http://localhost:8071"))

	r.Any("/messages/*any", proxy("http://localhost:8072"))

	r.Any("/messages", proxy("http://localhost:8072"))

	r.Any("/likes/*any", proxy("http://localhost:8073"))

	r.Any("/likes", proxy("http://localhost:8073"))

	httpServer := &http.Server{
		Addr:    ":8070",
		Handler: r,
	}

	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatal("Failed to start Proxy:", err)
		return
	}
}
