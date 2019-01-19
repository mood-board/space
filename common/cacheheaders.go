package common

import (
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func hasImageExt(extension string) bool {
	switch strings.ToLower(extension) {
	case
		".gif",
		".jpeg",
		".jpg",
		".png",
		".svg",
		".webp":
		return true
	}
	return false
}

func CacheHeaders(staticPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request.URL.Path

		if strings.HasPrefix(r, staticPath) {
			if hasImageExt(path.Ext(r)) {
				// Assume that all image resources are fairly
				// stable and keep them for 3 minutes.
				c.Header("Cache-Control", "max-age=180")
			} else {
				// Use "no-cache" for assets that have or can
				// have interdependencies like css and js. UA
				// will still cache resources, but have to check
				// with server before using the asset. Server
				// responds with 304 Not Modified if it is ok to
				// use the cached version.
				c.Header("Cache-Control", "no-cache")
			}
		} else {
			// Prevent caching for all other resources as recommended by
			// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control
			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		}

		c.Next()
	}
}
