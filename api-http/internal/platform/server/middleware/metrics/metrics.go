package metrics

import (
	creatingmetric "example.com/gotraining/go-hexagonal_http_api-course/internal/metrics/creating"
	"example.com/gotraining/go-hexagonal_http_api-course/kit/bus"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware is a gin.HandlerFunc that logs some information
// of the incoming request and the consequent response.
func Middleware(bus bus.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		// Request URL path
		path := ctx.Request.URL.Path

		if ctx.Request.URL.RawQuery != "" {
			path = path + "?" + ctx.Request.URL.RawQuery
		}

		// Process request
		ctx.Next()
		// Results
		timestamp := time.Now()
		latency := timestamp.Sub(start)
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()

		fmt.Printf("[HTTP] %v | %3d | %13v | %15s | %-7s %#v\n",
			timestamp.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)

		err := bus.DispatchCommand(ctx, creatingmetric.NewMetricCommand(
			path,
			statusCode,
			20,
			20,
			latency.Milliseconds(),
		))

		if err != nil {
			fmt.Printf("[Metric middleware] %s panic recovered:\n%s\n", time.Now().Format("2006/01/02 - 15:04:05"), err)
		}


	}
}
