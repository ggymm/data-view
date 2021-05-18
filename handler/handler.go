package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/wire"
)

var SetHandler = wire.NewSet(
	DataSourceHandlerSet,
	ImageHandlerSet,
	DataViewHandlerSet,
)

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return err
	}
	return nil
}

// ParseQuery 解析Query参数
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return err
	}
	return nil
}

// ParseForm 解析Form请求
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return err
	}
	return nil
}

// 返回业务JSON数据
func returnJson(c *gin.Context, success bool, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"data":    data,
	})
}

func httpError(c *gin.Context, httpCode int, message interface{}) {
	c.JSON(httpCode, gin.H{"message": message})
}
