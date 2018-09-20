package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DummyController ...
type DummyController struct{}

// Ping ...
func (pc *DummyController) Index(c *gin.Context) {
	c.JSON(http.StatusOK,"ok")
}

// NewDummyController ...
func NewDummyController() *DummyController {
	return &DummyController{}
}
