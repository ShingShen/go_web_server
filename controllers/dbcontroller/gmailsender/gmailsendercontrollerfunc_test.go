package gmailsendercontroller

import (
	"testing"
)

func TestGmailSenderControllerFunc(t *testing.T) {
	gmailSenderControllerFuncFactory := &gmailSenderControllerFuncFactory{}
	const name string = "gmailSenderControllerFunc"
	const errName string = "errName"

	// getGmailSenderControllerFunc
	getGmailSenderControllerFunc, err := gmailSenderControllerFuncFactory.getGmailSenderControllerFunc(name)
	if err == nil {
		t.Logf("getGmailSenderControllerFunc passed: %v, %v", getGmailSenderControllerFunc, err)
	} else {
		t.Errorf("getGmailSenderControllerFunc failed: %v, %v", getGmailSenderControllerFunc, err)
	}

	// getGmailSenderControllerFunc Name Err
	getGmailSenderControllerFuncNameErr, err := gmailSenderControllerFuncFactory.getGmailSenderControllerFunc(errName)
	if err != nil {
		t.Logf("getGmailSenderControllerFuncNameErr passed: %v, %v", getGmailSenderControllerFuncNameErr, err)
	} else {
		t.Errorf("getGmailSenderControllerFuncNameErr failed: %v, %v", getGmailSenderControllerFuncNameErr, err)
	}
}
