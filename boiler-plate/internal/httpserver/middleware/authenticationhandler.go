package middleware

import (
	errorfazz "logger/internal/error"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// AuthenticationHandler ...
type AuthenticationHandler struct {
}

func (ah *AuthenticationHandler) cancelOperation(c *gin.Context) {
	c.Error(errorfazz.NewAuthorizationError())
	c.Abort()
}

func (ah *AuthenticationHandler) getDummyAuthorization(c *gin.Context) {

}

// GetMiddleware ...
func (ah *AuthenticationHandler) GetMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//ah.getAuthorization(c)
		ah.getDummyAuthorization(c)
		c.Next()
	}
}

// NewAuthenticationHandler ...
func NewAuthenticationHandler(mdb *sqlx.DB) *AuthenticationHandler {
	return &AuthenticationHandler{}
}
