package gmailsendercontroller

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestGmailSenderController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	gmailSenderControllerFactory := &GmailSenderControllerFactory{}
	const name string = "gmailSenderController"
	const errName string = "errName"

	// getGmailSenderController
	getGmailSenderController, err := gmailSenderControllerFactory.GetGmailSenderController(mockDB, name)
	if err == nil {
		t.Logf("getGmailSenderController passed: %v, %v", getGmailSenderController, err)
	} else {
		t.Errorf("getGmailSenderController failed: %v, %v", getGmailSenderController, err)
	}

	// getGmailSenderController Name Err
	getGmailSenderControllerNameErr, err := gmailSenderControllerFactory.GetGmailSenderController(mockDB, errName)
	if err != nil {
		t.Logf("getGmailSenderControllerNameErr passed: %v, %v", getGmailSenderControllerNameErr, err)
	} else {
		t.Errorf("getGmailSenderControllerNameErr failed: %v, %v", getGmailSenderControllerNameErr, err)
	}
}
