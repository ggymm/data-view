package handler

import (
	"data-view/model"
	"data-view/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"strconv"
)

var DataViewHandlerSet = wire.NewSet(wire.Struct(new(DataViewHandler), "*"))

type DataViewHandler struct {
	DataViewModel   *model.DataViewModel
	DataSourceModel *model.DataSourceModel
}

func (h *DataViewHandler) GetPage(c *gin.Context) {
	var params schema.DataViewQueryParam
	if err := ParseQuery(c, &params); err != nil {
		httpError(c, http.StatusBadRequest, err)
		return
	}
	params.BusinessId = c.GetInt64("BusinessId")
	if list, count, err := h.DataViewModel.GetPage(params); err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	} else {
		returnJson(c, true, map[string]interface{}{"list": list, "count": count})
		return
	}
}

func (h *DataViewHandler) Get(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if id <= 0 || err != nil {
		httpError(c, http.StatusBadRequest, "ID格式不正确")
		return
	}
	if v, err := h.DataViewModel.Get(id, c.GetInt64("BusinessId")); err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	} else {
		returnJson(c, true, v)
		return
	}
}

func (h *DataViewHandler) GetChartData(c *gin.Context) {
	var params *schema.ChartDataParams
	if err := ParseQuery(c, &params); err != nil {
		httpError(c, http.StatusBadRequest, err)
		return
	}
	dataSource, err := h.DataSourceModel.Get(params.Database)
	if err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	}
	if result, err := h.DataViewModel.GetChartData(params, dataSource); err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	} else {
		returnJson(c, true, result)
		return
	}
}
