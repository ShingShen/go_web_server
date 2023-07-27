package storagefileservice

import (
	"fmt"
	"mime/multipart"

	"cloud.google.com/go/storage"

	storageFileServiceProvider "server/services/dbservice/serviceprovider/storagefile"
)

type StorageFileServiceFactory struct{}

func (s *StorageFileServiceFactory) GetStorageFileService(name string) (IStorageFileService, error) {
	if name == "storageFile" {
		return &StorageFileService{
			storageFileServiceFuncFactory: &storageFileServiceProvider.StorageFileServiceFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong file type passed")
}

type StorageFileService struct {
	storageFileServiceFuncFactory storageFileServiceProvider.IStorageFileServiceFuncFactory
}

func (s *StorageFileService) Upload(file multipart.File, objectName string, cloudStorage *storage.Client) error {
	storageFileServiceFunc, _ := s.storageFileServiceFuncFactory.GetStorageFileServiceFunc("storageFileServiceFunc")
	return storageFileServiceFunc.Upload(file, objectName, cloudStorage)
}
