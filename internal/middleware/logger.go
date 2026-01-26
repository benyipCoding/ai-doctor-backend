package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// 先执行后续 handler
		c.Next()

		// 请求处理完成后
		latency := time.Since(start)
		status := c.Writer.Status()
		traceID := c.GetString("trace_id")

		fields := []zap.Field{
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.String("trace_id", traceID),
		}

		// 根据状态码决定日志级别
		switch {
		case status >= 500:
			logger.Error("http request failed", fields...)
		case status >= 400:
			logger.Warn("http request client error", fields...)
		default:
			logger.Info("http request completed", fields...)
		}
	}
}
