package handler

import (
	"mime/multipart"
)

var StorageHandlers = make(map[string]Storage)

func init() {
	StorageHandlers["go-fastdfs"] = FastdfsStorageHandle
	StorageHandlers["minio"] = MinioStorageHandle
	StorageHandlers["custom"] = CustomStorageHandle
}

type Storage interface {
	Upload(uploadFile *multipart.FileHeader) (string, error)
}

type StorageHandler struct {
	RunUpload func(uploadFile *multipart.FileHeader) (string, error)
}

func (handler *StorageHandler) Upload(uploadFile *multipart.FileHeader) (string, error) {
	return handler.RunUpload(uploadFile)
}

// FastdfsStorageHandle
// see: https://github.com/sjqzhang/go-fastdfs
var FastdfsStorageHandle = &StorageHandler{RunUpload: func(uploadFile *multipart.FileHeader) (string, error) {
	return "", nil
}}

// MinioStorageHandle
// see: https://github.com/sjqzhang/go-fastdfs
var MinioStorageHandle = &StorageHandler{RunUpload: func(uploadFile *multipart.FileHeader) (string, error) {
	return "", nil
}}

// CustomStorageHandle
// 简单的
var CustomStorageHandle = &StorageHandler{RunUpload: func(uploadFile *multipart.FileHeader) (string, error) {
	return "", nil
}}
