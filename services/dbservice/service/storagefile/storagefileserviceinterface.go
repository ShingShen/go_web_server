package storagefileservice

import (
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type IStorageFileServiceFactory interface {
	GetStorageFileService(name string) (IStorageFileService, error)
}

type IStorageFileService interface {
	Upload(file multipart.File, objectName string, cloudStorage *storage.Client) error
}
