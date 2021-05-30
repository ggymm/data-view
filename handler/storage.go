package handler

import (
	"data-view/config"
	"errors"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"time"
)

var StorageHandlers = make(map[string]Storage)

func init() {
	StorageHandlers["go-fastdfs"] = FastdfsStorageHandle
	StorageHandlers["minio"] = MinioStorageHandle
	StorageHandlers["custom"] = CustomStorageHandle
}

type Storage interface {
	Upload(c *gin.Context) (string, error)
}

type StorageHandler struct {
	RunUpload func(c *gin.Context) (string, error)
}

func (handler *StorageHandler) Upload(c *gin.Context) (string, error) {
	return handler.RunUpload(c)
}

// FastdfsStorageHandle
// see: https://github.com/sjqzhang/go-fastdfs
var FastdfsStorageHandle = &StorageHandler{RunUpload: func(c *gin.Context) (string, error) {
	return "", errors.New("未实现")
}}

// MinioStorageHandle
// see: https://github.com/sjqzhang/go-fastdfs
var MinioStorageHandle = &StorageHandler{RunUpload: func(c *gin.Context) (string, error) {
	return "", errors.New("未实现")
}}

// CustomStorageHandle
// 简单的
var CustomStorageHandle = &StorageHandler{RunUpload: func(c *gin.Context) (string, error) {
	file, _ := c.FormFile("file")
	path := config.Instance.Storage.Path + "/" + time.Now().Format("2006/0102") + "/" + file.Filename
	_ = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err := c.SaveUploadedFile(file, path); err != nil {
		return "", err
	}
	return "/" + path, nil
}}
