package handler

import (
	"net/http"

	"data-view/model"
	"data-view/schema"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
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

func (h *ImageHandler) UploadHtml(c *gin.Context) {
	const uploadHtml string = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>测试上传文件</title>
</head>
<body>
<form action="/oss/api/v1/upload" method="post" enctype="multipart/form-data">
	<label>项目名称（必填）:
		<input type="text" name="project"/>
	</label>
	<label>模块名称（非必填）:
		<input type="text" name="module"/>
	</label>
	<input type="file" name="file"/>
	<input type="submit" value="上传"/>
</form>
</body>
</html>`
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, uploadHtml)
	return
}

func (h *ImageHandler) Upload(c *gin.Context) {

}
