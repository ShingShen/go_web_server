package router

import (
	mock_sqloperator "server/tests/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	handlerFactory := &handlerFactory{}
	const name string = "handler"
	const errName string = "errName"

	// getHandler
	getHandler, err := handlerFactory.getHandler(mockDB, name)
	if err == nil {
		t.Logf("getHandler passed: %v, %v", getHandler, err)
	} else {
		t.Errorf("getHandler failed: %v, %v", getHandler, err)
	}

	// getHandler Name Err
	getHandlerNameErr, err := handlerFactory.getHandler(mockDB, errName)
	if err != nil {
		t.Logf("getHandlerNameErr passed: %v, %v", getHandlerNameErr, err)
	} else {
		t.Errorf("getHandlerNameErr failed: %v, %v", getHandlerNameErr, err)
	}
}
