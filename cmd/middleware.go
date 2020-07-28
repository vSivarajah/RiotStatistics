package cmd

import (
	"github.com/gin-gonic/gin"
)

// add common middlewares here
func CustomHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
