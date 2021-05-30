package handler

import (
	"net/http"

	"data-view/model"
	"data-view/schema"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var DataSourceHandlerSet = wire.NewSet(wire.Struct(new(DataSourceHandler), "*"))

type DataSourceHandler struct {
	DataSourceModel *model.DataSourceModel
}

func (h *DataSourceHandler) GetPage(c *gin.Context) {
	var params schema.DataSourceQueryParam
	if err := ParseQuery(c, &params); err != nil {
		httpError(c, http.StatusBadRequest, err.Error())
		return
	}
	params.BusinessId = c.GetInt64("BusinessId")
	if list, count, err := h.DataSourceModel.GetPage(params); err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		returnJson(c, true, map[string]interface{}{"list": list, "count": count})
		return
	}
}

func (h *DataSourceHandler) GetList(c *gin.Context) {
	if list, err := h.DataSourceModel.GetList(c.GetInt64("BusinessId")); err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		returnJson(c, true, list)
		return
	}
}
