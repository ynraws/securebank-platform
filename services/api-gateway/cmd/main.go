package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/securebank/api-gateway/internal/middleware"
	"github.com/securebank/api-gateway/internal/proxy"
)

func main() {
	r := gin.Default()
	r.Use(middleware.RequestLogger())
	r.Use(middleware.CORS())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "api-gateway",
			"routes": gin.H{
				"account-svc": "http://account-svc:8081",
				"payment-svc": "http://payment-svc:8082",
				"fraud-svc":   "http://fraud-svc:8083",
				"notify-svc":  "http://notify-svc:8084",
			},
		})
	})

	r.Any("/api/v1/accounts", proxy.ProxyToFixed("http://account-svc:8081"))
	r.Any("/api/v1/accounts/*path", proxy.ProxyToFixed("http://account-svc:8081"))
	r.Any("/api/v1/payments", proxy.ProxyToFixed("http://payment-svc:8082"))
	r.Any("/api/v1/payments/*path", proxy.ProxyToFixed("http://payment-svc:8082"))
	r.Any("/api/v1/fraud", proxy.ProxyToFixed("http://fraud-svc:8083"))
	r.Any("/api/v1/fraud/*path", proxy.ProxyToFixed("http://fraud-svc:8083"))
	r.Any("/api/v1/notify", proxy.ProxyToFixed("http://notify-svc:8084"))
	r.Any("/api/v1/notify/*path", proxy.ProxyToFixed("http://notify-svc:8084"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("api-gateway starting on port %s", port)
	r.Run(":" + port)
}
