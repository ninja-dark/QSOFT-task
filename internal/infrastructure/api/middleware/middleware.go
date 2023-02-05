package middleware

import "github.com/gin-gonic/gin"

func SetPongHeader(c *gin.Context) {
	header := c.GetHeader("X-PING")
	if header == "ping" {
		c.Writer.Header().Set("X-PONG", "pong")
	}
}
