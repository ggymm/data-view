package handler

import (
	"net/http"
	"time"

	"data-view/config"
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
		httpError(c, http.StatusBadRequest, err.Error())
		return
	}
	params.BusinessId = c.GetInt64("BusinessId")
	if list, count, err := h.ImageModel.GetPage(params); err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		returnJson(c, true, map[string]interface{}{"list": list, "count": count})
		return
	}
}

func (h *ImageHandler) GetList(c *gin.Context) {
	if list, err := h.ImageModel.GetList(c.GetInt64("BusinessId")); err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
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
<form action="/api/v1/image" method="post" enctype="multipart/form-data">
	<input type="file" name="file" />
	<input type="submit" value="上传" />
</form>
</body>
</html>`
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, uploadHtml)
	return
}

func (h *ImageHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		httpError(c, http.StatusBadRequest, err.Error())
		return
	}
	if path, err := StorageHandlers[config.Instance.Storage.Type].Upload(c); err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		image := new(model.Image)
		image.ImageName = file.Filename
		image.ImagePath = path
		image.ImageSize = file.Size
		image.BusinessId = c.GetInt64("BusinessId")
		image.CreateId = c.GetInt64("UserId")
		image.CreateTime = schema.JsonTime(time.Now())
		image.UpdateId = c.GetInt64("UserId")
		image.UpdateTime = schema.JsonTime(time.Now())
		if err := h.ImageModel.Save(image); err != nil {
			returnJson(c, false, err.Error())
			return
		} else {
			returnJson(c, true, path)
			return
		}
	}
}

func (h *ImageHandler) UploadThumbnail(c *gin.Context) {
	if path, err := StorageHandlers[config.Instance.Storage.Type].Upload(c); err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		returnJson(c, true, path)
		return
	}
}
