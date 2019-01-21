package common

import "github.com/gin-gonic/gin"

func EnsureHTTPVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		major := c.Request.ProtoMajor
		minor := c.Request.ProtoMinor

		if (major == 1 && minor == 1) || (major == 2 && minor == 0) {
			c.Next()
		} else {
			c.Header("Set-Cookie", "")
			c.Header("Vary", "")
			c.AbortWithStatus(505)
		}
	}
}
