package storagefilecontroller

import (
	"testing"
)

func TestGetStorageFileControllerFunc(t *testing.T) {
	storageFileControllerFuncFactory := &storageFileControllerFuncFactory{}
	const name string = "storageFileControllerFunc"
	const errName string = "errName"

	// storageFileControllerFunc
	storageFileControllerFunc, err := storageFileControllerFuncFactory.getStorageFileControllerFunc(name)
	if err == nil {
		t.Logf("storageFileControllerFunc passed: %v, %v", storageFileControllerFunc, err)
	} else {
		t.Errorf("storageFileControllerFunc failed: %v, %v", storageFileControllerFunc, err)
	}

	// storageFileControllerFunc Name Err
	storageFileControllerFuncNameErr, err := storageFileControllerFuncFactory.getStorageFileControllerFunc(errName)
	if err != nil {
		t.Logf("storageFileControllerFuncNameErr passed: %v, %v", storageFileControllerFuncNameErr, err)
	} else {
		t.Errorf("storageFileControllerFuncNameErr failed: %v, %v", storageFileControllerFuncNameErr, err)
	}
}
