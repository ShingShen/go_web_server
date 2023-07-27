package schedulecontroller

import (
	mock_sqloperator "server/tests/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetScheduleController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	scheduleControllerFactory := &ScheduleControllerFactory{}
	const name string = "scheduleController"
	const errName string = "errName"

	// scheduleController
	spirometerController, err := scheduleControllerFactory.GetScheduleController(mockDB, name)
	if err == nil {
		t.Logf("scheduleController passed: %v, %v", spirometerController, err)
	} else {
		t.Errorf("scheduleController failed: %v, %v", spirometerController, err)
	}

	// scheduleController Name Err
	scheduleControllerNameErr, err := scheduleControllerFactory.GetScheduleController(mockDB, errName)
	if err != nil {
		t.Logf("scheduleControllerNameErr passed: %v, %v", scheduleControllerNameErr, err)
	} else {
		t.Errorf("scheduleControllerNameErr failed: %v, %v", scheduleControllerNameErr, err)
	}
}

func TestController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	scheduleController := &ScheduleController{mockDB, &scheduleControllerFuncFactory{}}

	scheduleController.Create()
	scheduleController.Update()
	scheduleController.GetAnEvent()
	scheduleController.GetOneDayEvents()
	scheduleController.GetOneMonthEvents()
	scheduleController.Delete()
}
