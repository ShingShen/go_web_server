package router

import (
	"net/http"
	gmailSenderController "server/controllers/dbcontroller/gmailsender"
	memoController "server/controllers/dbcontroller/memo"
	scheduleController "server/controllers/dbcontroller/schedule"
	mock_sqloperator "server/tests/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetHandleFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	handleFuncFactory := &handleFuncFactory{}
	const name string = "handleFunc"
	const errName string = "errName"
	var mux *http.ServeMux = http.NewServeMux()

	// getHandleFunc
	getHandleFunc, err := handleFuncFactory.getHandleFunc(mockDB, mux, name)
	if err == nil {
		t.Logf("getHandleFunc passed: %v, %v", getHandleFunc, err)
	} else {
		t.Errorf("getHandleFunc failed: %v, %v", getHandleFunc, err)
	}

	// getHandleFunc Name Err
	getHandleFuncNameErr, err := handleFuncFactory.getHandleFunc(mockDB, mux, errName)
	if err != nil {
		t.Logf("getHandleFuncNameErr passed: %v, %v", getHandleFuncNameErr, err)
	} else {
		t.Errorf("getHandleFuncNameErr failed: %v, %v", getHandleFuncNameErr, err)
	}
}

func TestHandleFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	var mux *http.ServeMux = http.NewServeMux()
	handleFunc := &handleFunc{
		db:                           mockDB,
		mux:                          mux,
		gmailSenderControllerFactory: &gmailSenderController.GmailSenderControllerFactory{},
		memoControllerFactory:        &memoController.MemoControllerFactory{},
		scheduleControllerFactory:    &scheduleController.ScheduleControllerFactory{},
	}

	handleFunc.webHandleFunc()
	handleFunc.gmailSenderHandleFunc("/gmail")
	handleFunc.memoHandleFunc("/memo")
	handleFunc.scheduleHandleFunc("/schedule")
}
