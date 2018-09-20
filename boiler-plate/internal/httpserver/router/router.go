package router

import (
	"logger/internal/httpserver/controller"
	"logger/internal/httpserver/middleware"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

// NewV1Router ...
func NewV1Router() http.Handler {
	dummyController := controller.NewDummyController()

	r := gin.Default()

	// catch http Not Found, can't be done in middleware
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Cannot %s %s", c.Request.Method, c.Request.URL))
	})

	// catch any other error
	r.Use(middleware.ErrorHandler())

	v1 := r.Group("/v1")
	{
		v1.GET("/", dummyController.Index)
	}

	return r
}
