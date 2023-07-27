package storagefileservice

import (
	"testing"
)

func TestGetStorageFileService(t *testing.T) {
	storageFileServiceFactory := &StorageFileServiceFactory{}
	const name string = "storageFile"
	const errName string = "errName"

	storageFileService, err := storageFileServiceFactory.GetStorageFileService(name)
	if err == nil {
		t.Logf("storageFileService passed: %v, %v", storageFileService, err)
	} else {
		t.Errorf("storageFileService failed: %v, %v", storageFileService, err)
	}

	storageFileServiceNameErr, err := storageFileServiceFactory.GetStorageFileService(errName)
	if err != nil {
		t.Logf("storageFileServiceNameErr passed: %v, %v", storageFileServiceNameErr, err)
	} else {
		t.Errorf("storageFileServiceNameErr failed: %v, %v", storageFileServiceNameErr, err)
	}
}
