package main

import (
	"net/http"

	"data-view/config"
	"data-view/handler"
	"data-view/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

type Router struct {
	DataSource *handler.DataSourceHandler
	Image      *handler.ImageHandler
	DataView   *handler.DataViewHandler
}

func (r *Router) NewRouter() (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	if config.Instance.App.AllowCrossDomain {
		router.Use(middleware.Cors())
	}
	router.Use(middleware.ErrHandler())

	router.StaticFS("/storage", http.Dir("storage"))
	v1 := router.Group("/api/v1/")
	{
		// 状态
		v1.GET("status", func(c *gin.Context) {
			c.String(http.StatusOK, "start success")
		})

		// 数据源管理
		dataSource := v1.Group("data-source/").Use(middleware.CheckUser())
		{
			dataSource.GET("/", r.DataSource.GetPage)
			dataSource.GET("/list", r.DataSource.GetList)
		}

		// 图片管理
		image := v1.Group("image/").Use(middleware.CheckUser())
		{
			image.GET("/", r.Image.GetPage)
			image.GET("/list", r.Image.GetList)
			image.GET("/upload.html", r.Image.UploadHtml)
			image.POST("/", r.Image.Upload)
			image.POST("/thumbnail", r.Image.UploadThumbnail)
		}

		// 数据可视化大屏管理
		dataView := v1.Group("data-view/").Use(middleware.CheckUser())
		{
			dataView.GET("/", r.DataView.GetPage)
			dataView.GET("/:id", r.DataView.Get)
			dataView.POST("/", r.DataView.Create)
			dataView.PUT("/", r.DataView.Update)
			dataView.GET("/chart-data", r.DataView.GetChartData)
			// 图表测试获取数据接口
			dataView.GET("/test-chart-data", r.DataView.GetTestChartData)
		}

		// 数据报表管理
		dataReport := v1.Group("data-report/").Use(middleware.CheckUser())
		{
			dataReport.GET("/", r.DataView.GetPage)
		}
	}

	return
}
