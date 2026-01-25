package response

import (
	"net/http"

	"ai-doctor-backend/internal/errors"

	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context, data any, err error) {
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "OK",
			"data": data,
		})
		return
	}

	if bizErr, ok := err.(*errors.BizError); ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    bizErr.Code,
			"message": bizErr.Message,
		})
		return
	}

	// system error
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    "INTERNAL_ERROR",
		"message": "internal server error",
	})
}
