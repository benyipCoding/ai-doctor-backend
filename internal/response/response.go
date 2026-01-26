package response

import (
	"ai-doctor-backend/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context, data any, err error) {
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "2000",
			"data":    data,
			"message": "success",
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
