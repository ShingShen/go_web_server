package storagefilecontroller

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/storage"

	storageFileService "server/services/dbservice/service/storagefile"
)

type storageFileControllerFuncFactory struct{}

func (s *storageFileControllerFuncFactory) getStorageFileControllerFunc(name string) (IStorageFileControllerFunc, error) {
	if name == "storageFileControllerFunc" {
		return &storageFileControllerFunc{
			storageFileServiceFactory: &storageFileService.StorageFileServiceFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong storage file controller func type passed")
}

type storageFileControllerFunc struct {
	storageFileServiceFactory storageFileService.IStorageFileServiceFactory
}

func (s *storageFileControllerFunc) uploadController(cloudStorage *storage.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		objectName := r.FormValue("object_name")
		fmt.Printf("objectName: %+v\n", objectName)

		getService, _ := s.storageFileServiceFactory.GetStorageFileService("storageFile")
		getService.Upload(file, objectName, cloudStorage)
	}
}
