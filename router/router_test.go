package router

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestNewServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	const name string = "server"
	const errName string = "errName"

	newServer, err := NewServer(mockDB, name)
	if err == nil {
		t.Logf("newServer passed: %v, %v", newServer, err)
	} else {
		t.Errorf("newServer failed: %v, %v", newServer, err)
	}

	newServerNameErr, err := NewServer(mockDB, errName)
	if err != nil {
		t.Logf("newServerNameErr passed: %v, %v", newServerNameErr, err)
	} else {
		t.Errorf("newServerNameErr failed: %v, %v", newServerNameErr, err)
	}
}
