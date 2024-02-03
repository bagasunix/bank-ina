package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func RemoveTrailingSlash() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path != "/" && c.Request.URL.Path[len(c.Request.URL.Path)-1] == '/' {
			c.Redirect(http.StatusMovedPermanently, c.Request.URL.Path[:len(c.Request.URL.Path)-1])
			return
		}
		c.Next()
	}
}

func Secure() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "SAMEORIGIN")
		c.Header("Strict-Transport-Security", "max-age=3600")
		c.Header("Content-Security-Policy", "default-src 'self'")

		c.Next()
	}
}
