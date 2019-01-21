package common

import (
	"github.com/gin-gonic/gin"
)

func SecureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {

		// The HTTP X-XSS-Protection response header is a feature of the browsers that stops pages from loading when they detect
		// reflected cross-site scripting (XSS) attacks. Value "1; mode=block" enables XSS filtering - rather than sanitizing the page,
		// the browser will prevent rendering of the page if an attack is detected.
		c.Writer.Header().Add("X-XSS-Protection", "1; mode=block")

		// Prevents MIME Sniffing vulnerabilities which occur when a website allows users to upload data to the server.
		// TODO Enable when if we have file upload.
		// c.Writer.Header().Add("X-Content-Type-Options", "nosniff")

		// Prevents the page from being displayed in an iframe
		// TODO(vincents): Only allow frames for the gateway page
		// c.Writer.Header().Add("X-Frame-Options", "DENY")

		c.Next()
	}
}
