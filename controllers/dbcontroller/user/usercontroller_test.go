package usercontroller

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestGetUserController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userControllerFactory := &UserControllerFactory{}
	const name string = "userController"
	const errName string = "errName"

	// userController
	userController, err := userControllerFactory.GetUserController(mockDB, name)
	if err == nil {
		t.Logf("userController passed: %v, %v", userController, err)
	} else {
		t.Errorf("userController failed: %v, %v", userController, err)
	}

	// userController Name Err
	userControllerrNameErr, err := userControllerFactory.GetUserController(mockDB, errName)
	if err != nil {
		t.Logf("userControllerrNameErr passed: %v, %v", userControllerrNameErr, err)
	} else {
		t.Errorf("userControllerrNameErr failed: %v, %v", userControllerrNameErr, err)
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userController := &UserController{mockDB, &userControllerFuncFactory{}}

	userController.Update()
}

func TestResetUserPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userController := &UserController{mockDB, &userControllerFuncFactory{}}

	userController.ResetUserPassword()
}

func TestEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userController := &UserController{mockDB, &userControllerFuncFactory{}}

	userController.Enable()
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userController := &UserController{mockDB, &userControllerFuncFactory{}}

	userController.Get()
}

func TestGetUserByAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userController := &UserController{mockDB, &userControllerFuncFactory{}}

	userController.GetUserByAccount()
}

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userController := &UserController{mockDB, &userControllerFuncFactory{}}

	userController.GetAllUsers()
}

func TestGetSpecificRoles(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userController := &UserController{mockDB, &userControllerFuncFactory{}}

	userController.GetSpecificRoles()
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userController := &UserController{mockDB, &userControllerFuncFactory{}}

	userController.Delete()
}
