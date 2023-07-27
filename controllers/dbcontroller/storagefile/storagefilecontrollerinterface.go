package storagefilecontroller

import (
	"net/http"
	sqlOperator "server/utils/sqloperator"

	"cloud.google.com/go/storage"
)

type IStorageFileControllerFactory interface {
	GetStorageFileController(db sqlOperator.ISqlDB, name string) (IStorageFileController, error)
}

type IStorageFileController interface {
	Upload(cloudStorage *storage.Client) http.HandlerFunc
}

type IStorageFileControllerFuncFactory interface {
	getStorageFileControllerFunc(name string) (IStorageFileControllerFunc, error)
}

type IStorageFileControllerFunc interface {
	uploadController(loudStorage *storage.Client) func(w http.ResponseWriter, r *http.Request)
}
