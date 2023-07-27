package storagefileserviceprovider

import (
	"fmt"
	"mime/multipart"

	"cloud.google.com/go/storage"

	storageFileOperator "server/utils/storagefileoperator"
)

type StorageFileServiceFuncFactory struct{}

func (s *StorageFileServiceFuncFactory) GetStorageFileServiceFunc(name string) (IStorageFileServiceFunc, error) {
	if name == "storageFileServiceFunc" {
		return &StorageFileServiceFunc{}, nil
	}
	return nil, fmt.Errorf("wrong storage file service func type passed")
}

type StorageFileServiceFunc struct{}

func (s *StorageFileServiceFunc) Upload(file multipart.File, objectName string, cloudStorage *storage.Client) error {
	bucket := "bucket"
	return storageFileOperator.StorageFileUpload(file, bucket, objectName, cloudStorage)
}
