package storagefileserviceprovider

import (
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type IStorageFileServiceFuncFactory interface {
	GetStorageFileServiceFunc(name string) (IStorageFileServiceFunc, error)
}

type IStorageFileServiceFunc interface {
	Upload(file multipart.File, objectName string, cloudStorage *storage.Client) error
}
