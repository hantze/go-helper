package middleware

import (
	error "logger/internal/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

func respondWithError(code int, e interface{}, c *gin.Context) {
	c.JSON(code, e)
	c.Abort()
}

// ErrorHandler ...
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, e := range c.Errors {
			switch e.Err.(type) {
			case
				*error.AuthorizationError:
				respondWithError(http.StatusUnauthorized, e.Err, c)

			case *error.ThrottleError:
				respondWithError(http.StatusTooManyRequests, e.Err, c)

			case *error.ValidateError, *error.TestError,
				*error.HTTPError:
				respondWithError(http.StatusBadRequest, e.Err, c)

			default:
				respondWithError(http.StatusBadRequest, "Server Error", c)
			}
		}
	}
}
