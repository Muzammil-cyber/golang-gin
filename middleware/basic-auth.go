package middleware

import "github.com/gin-gonic/gin"

// BasicAuthMiddleware provides HTTP Basic Authentication for Gin routes.
func BasicAuthMiddleware(username, password string) gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		username: password,
	})
}
