package middleware

import (
	"fmt"
	errorfazz "logger/internal/error"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter"
	mgin "github.com/ulule/limiter/drivers/middleware/gin"
)

// ThrottleMiddleware ...
type ThrottleMiddleware struct {
	Limiter        *limiter.Limiter
	OnError        mgin.ErrorHandler
	OnLimitReached mgin.LimitReachedHandler
	KeyGetter      mgin.KeyGetter
}

func modifiedKeyGetter(c *gin.Context) string {
	return fmt.Sprintf("%s%s", c.ClientIP(), c.Request.URL)
}

// NewThrottleMiddleware ...
func NewThrottleMiddleware(limiter *limiter.Limiter) gin.HandlerFunc {
	middleware := &ThrottleMiddleware{
		Limiter:        limiter,
		OnError:        mgin.DefaultErrorHandler,
		OnLimitReached: mgin.DefaultLimitReachedHandler,
		KeyGetter:      modifiedKeyGetter,
	}

	return func(ctx *gin.Context) {
		middleware.Handle(ctx)
	}
}

// Handle ...
func (middleware *ThrottleMiddleware) Handle(c *gin.Context) {
	key := middleware.KeyGetter(c)
	context, err := middleware.Limiter.Get(c, key)
	if err != nil {
		c.Error(errorfazz.NewThrottleError("throttle error"))
		c.Abort()
		return
	}
	c.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
	c.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
	c.Header("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))

	if context.Reached {
		c.Error(errorfazz.NewThrottleError("limit exceeded"))
		c.Abort()
		return
	}

	c.Next()
}
