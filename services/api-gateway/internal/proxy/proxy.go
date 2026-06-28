package proxy

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProxyToFixed(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fullURL := fmt.Sprintf("%s%s", target, c.Request.URL.Path)
		if c.Request.URL.RawQuery != "" {
			fullURL += "?" + c.Request.URL.RawQuery
		}

		req, err := http.NewRequest(c.Request.Method, fullURL, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "proxy error"})
			return
		}

		for key, values := range c.Request.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error":  "service unavailable",
				"target": target,
			})
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				c.Header(key, value)
			}
		}

		body, _ := io.ReadAll(resp.Body)
		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
	}
}

func ProxyTo(target string) gin.HandlerFunc {
	return ProxyToFixed(target)
}
