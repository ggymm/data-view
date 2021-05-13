package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("BusinessId", int64(1))
		c.Set("UserId", int64(1))
		c.Next()
		return
	}
}
