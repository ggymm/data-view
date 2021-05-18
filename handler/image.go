package handler

import (
	"data-view/model"
	"data-view/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var ImageHandlerSet = wire.NewSet(wire.Struct(new(ImageHandler), "*"))

type ImageHandler struct {
	ImageModel *model.ImageModel
}

func (h *ImageHandler) GetPage(c *gin.Context) {
	var params schema.ImageQueryParam
	if err := ParseQuery(c, &params); err != nil {
		httpError(c, http.StatusBadRequest, err)
		return
	}
	params.BusinessId = c.GetInt64("BusinessId")
	if list, count, err := h.ImageModel.GetPage(params); err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	} else {
		returnJson(c, true, map[string]interface{}{"list": list, "count": count})
		return
	}
}

func (h *ImageHandler) GetList(c *gin.Context) {
	if list, err := h.ImageModel.GetList(c.GetInt64("BusinessId")); err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	} else {
		returnJson(c, true, list)
		return
	}
}
