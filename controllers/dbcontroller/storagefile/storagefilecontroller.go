package storagefilecontroller

import (
	"fmt"
	"net/http"
	"server/middleware"

	"cloud.google.com/go/storage"

	sqlOperator "server/utils/sqloperator"
)

type StorageFileControllerFactory struct{}

func (s *StorageFileControllerFactory) GetStorageFileController(db sqlOperator.ISqlDB, name string) (IStorageFileController, error) {
	if name == "storageFileController" {
		return &StorageFileController{
			db:                               db,
			storageFileControllerFuncFactory: &storageFileControllerFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong file controller type passed")
}

type StorageFileController struct {
	db                               sqlOperator.ISqlDB
	storageFileControllerFuncFactory IStorageFileControllerFuncFactory
}

func (s *StorageFileController) Upload(cloudStorage *storage.Client) http.HandlerFunc {
	storageFileControllerFunc, _ := s.storageFileControllerFuncFactory.getStorageFileControllerFunc("storageFileControllerFunc")
	controller := storageFileControllerFunc.uploadController(cloudStorage)
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.UploadFileMethod(
				middleware.AllUserAuth(s.db, controller),
			),
		),
	)
}
