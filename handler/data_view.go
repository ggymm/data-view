package handler

import (
	"data-view/chart"
	"encoding/json"
	"net/http"
	"strconv"

	"data-view/model"
	"data-view/schema"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var DataViewHandlerSet = wire.NewSet(wire.Struct(new(DataViewHandler), "*"))

type DataViewHandler struct {
	DataViewModel   *model.DataViewModel
	DataSourceModel *model.DataSourceModel
}

func (h *DataViewHandler) GetPage(c *gin.Context) {
	var params schema.DataViewQueryParam
	if err := ParseQuery(c, &params); err != nil {
		httpError(c, http.StatusBadRequest, err.Error())
		return
	}
	params.BusinessId = c.GetInt64("BusinessId")
	if list, count, err := h.DataViewModel.GetPage(params); err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
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
		httpError(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		returnJson(c, true, v)
		return
	}
}

func (h *DataViewHandler) Create(c *gin.Context) {
	var createReq schema.CreateDataViewReq
	if err := ParseJSON(c, &createReq); err != nil {
		httpError(c, http.StatusBadRequest, validatorErrorData(err))
		return
	}

	vip := new(model.ViewInstancePro)
	err := schema.Copy(&vip, &createReq)
	if err != nil {
		returnJson(c, false, err.Error())
		return
	}
	vip.BusinessId = c.GetInt64("BusinessId")
	vip.CreateId = c.GetInt64("UserId")
	vip.UpdateId = c.GetInt64("UserId")

	if id, err := h.DataViewModel.Create(vip); err != nil {
		returnJson(c, false, err.Error())
		return
	} else {
		returnJson(c, true, id)
		return
	}
}

func (h *DataViewHandler) Update(c *gin.Context) {
	var updateReq schema.UpdateDataViewReq
	if err := ParseJSON(c, &updateReq); err != nil {
		httpError(c, http.StatusBadRequest, validatorErrorData(err))
		return
	}

	vip := new(model.ViewInstancePro)
	err := schema.Copy(&vip, &updateReq)
	if err != nil {
		returnJson(c, false, err.Error())
		return
	}
	vip.BusinessId = c.GetInt64("BusinessId")
	vip.UpdateId = c.GetInt64("UserId")

	if err := h.DataViewModel.Update(vip); err != nil {
		returnJson(c, false, err.Error())
		return
	} else {
		returnJson(c, true, "success")
		return
	}
}

func (h *DataViewHandler) GetChartData(c *gin.Context) {
	var params schema.ChartDataParams
	if err := ParseJSON(c, &params); err != nil {
		httpError(c, http.StatusBadRequest, err.Error())
		return
	}
	if result, err := h.DataViewModel.GetChartData(params); err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		returnJson(c, true, result)
		return
	}
}

func (h *DataViewHandler) GetTestChartData(c *gin.Context) {
	query := c.Query("chartType")
	jsonData := chart.JsonDataMap[query]
	resultData := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonData), &resultData); err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
		return
	}
	returnJson(c, true, map[string]interface{}{"data": resultData})
	return
}
