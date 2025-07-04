package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // tiếp tục xử lý route

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[GIN] %d | %v | %s %s\n", status, latency, c.Request.Method, c.Request.URL.Path)
	}
}
