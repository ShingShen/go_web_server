package storagefilecontroller

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestGetStorageFileController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	storageFileControllerFactory := &StorageFileControllerFactory{}
	const name string = "storageFileController"
	const errName string = "errName"

	// storageFileController
	storageFileController, err := storageFileControllerFactory.GetStorageFileController(mockDB, name)
	if err == nil {
		t.Logf("storageFileController passed: %v, %v", storageFileController, err)
	} else {
		t.Errorf("storageFileController failed: %v, %v", storageFileController, err)
	}

	// storageFileController Name Err
	storageFileControllerNameErr, err := storageFileControllerFactory.GetStorageFileController(mockDB, errName)
	if err != nil {
		t.Logf("storageFileControllerNameErr passed: %v, %v", storageFileControllerNameErr, err)
	} else {
		t.Errorf("storageFileControllerNameErr failed: %v, %v", storageFileControllerNameErr, err)
	}
}
