package systemservice

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestGetSystemService(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	systemServiceFactory := &SystemServiceFactory{}
	const name string = "system"
	const errName string = "errName"

	systemService, err := systemServiceFactory.GetSystemService(mockDB, name)
	if err == nil {
		t.Logf("systemService passed: %v, %v", systemService, err)
	} else {
		t.Errorf("systemService failed: %v, %v", systemService, err)
	}

	systemServiceNameErr, err := systemServiceFactory.GetSystemService(mockDB, errName)
	if err != nil {
		t.Logf("systemServiceNameErr passed: %v, %v", systemServiceNameErr, err)
	} else {
		t.Errorf("systemServiceNameErr failed: %v, %v", systemServiceNameErr, err)
	}
}
