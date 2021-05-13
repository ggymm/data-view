package middleware

import (
	"fmt"
	"net/http"

	"data-view/logger"

	"github.com/gin-gonic/gin"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Instance.Error(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%v", err)})
				return
			}
		}()
		c.Next()
	}
}
