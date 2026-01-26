package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				traceID := c.GetString("trace_id")

				// 1. 打日志（一定要）
				logger.Error("panic recovered",
					zap.Any("err", err),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
					zap.String("trace_id", traceID),
					zap.Stack("stack"),
				)

				// 2. 返回统一 JSON
				c.AbortWithStatusJSON(500, gin.H{
					"code":    "INTERNAL_ERROR",
					"message": "internal server error",
				})
			}
		}()

		c.Next()
	}
}
