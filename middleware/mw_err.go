package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"

	"data-view/logger"

	"github.com/gin-gonic/gin"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Instance.Error(printStackTrace(err))
				c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%v", err)})
				return
			}
		}()
		c.Next()
	}
}

func printStackTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	_, _ = fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		_, _ = fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}
