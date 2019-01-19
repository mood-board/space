package mgo

import (
	"path"

	"github.com/gin-gonic/gin"
)

func DBConnectionMiddleware(db *Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		if path.Ext(c.Request.RequestURI) != "" {
			return
		}

		s := db.Session.Copy()
		defer s.Session.Close()

		c.Set(string(ContextKey), s)
		c.Next()
	}
}
