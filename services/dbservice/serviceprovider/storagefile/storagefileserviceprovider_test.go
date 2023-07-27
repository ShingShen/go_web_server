package storagefileserviceprovider

import (
	"testing"
)

func TestGetStorageFileServiceFunc(t *testing.T) {
	storageFileServiceFuncFactory := &StorageFileServiceFuncFactory{}
	const name string = "storageFileServiceFunc"
	const errName string = "errName"

	storageFileServiceFunc, err := storageFileServiceFuncFactory.GetStorageFileServiceFunc(name)
	if err == nil {
		t.Logf("storageFileServiceFunc passed: %v, %v", storageFileServiceFunc, err)
	} else {
		t.Errorf("storageFileServiceFunc failed: %v, %v", storageFileServiceFunc, err)
	}

	storageFileServiceFuncNameErr, err := storageFileServiceFuncFactory.GetStorageFileServiceFunc(errName)
	if err != nil {
		t.Logf("storageFileServiceFuncNameErr passed: %v, %v", storageFileServiceFuncNameErr, err)
	} else {
		t.Errorf("storageFileServiceFuncNameErr failed: %v, %v", storageFileServiceFuncNameErr, err)
	}
}
