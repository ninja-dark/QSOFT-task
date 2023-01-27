package middleware

import "github.com/gin-gonic/gin"

func CheckHeader(c *gin.Context){
	header := c.Request.Header.Get("X-PING")
	if header == "ping"{
		c.Writer.Header().Set("X-PONG", "pong")
	}
	
}