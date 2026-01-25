package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 1. 打日志（一定要）
				log.Printf(
					"[PANIC] path=%s method=%s err=%v",
					c.Request.URL.Path,
					c.Request.Method,
					err,
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
