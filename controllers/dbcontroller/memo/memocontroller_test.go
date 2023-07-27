package memocontroller

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestGetMemoController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	memoControllerFactory := &MemoControllerFactory{}
	const name string = "memoController"
	const errName string = "errName"

	// memoController
	memoController, err := memoControllerFactory.GetMemoController(mockDB, name)
	if err == nil {
		t.Logf("memoController passed: %v, %v", memoController, err)
	} else {
		t.Errorf("memoController failed: %v, %v", memoController, err)
	}

	// memoController Name Err
	memoControllerNameErr, err := memoControllerFactory.GetMemoController(mockDB, errName)
	if err != nil {
		t.Logf("memoControllerNameErr passed: %v, %v", memoControllerNameErr, err)
	} else {
		t.Errorf("memoControllerNameErr failed: %v, %v", memoControllerNameErr, err)
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	memoController := &MemoController{mockDB, &memoControllerFuncFactory{}}

	memoController.Create()
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	memoController := &MemoController{mockDB, &memoControllerFuncFactory{}}

	memoController.Update()
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	memoController := &MemoController{mockDB, &memoControllerFuncFactory{}}

	memoController.Get()
}

func TestGetMemosByUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	memoController := &MemoController{mockDB, &memoControllerFuncFactory{}}

	memoController.GetMemosByUserId()
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	memoController := &MemoController{mockDB, &memoControllerFuncFactory{}}

	memoController.Delete()
}
