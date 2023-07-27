package schedulecontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"server/dto"
	scheduleservice "server/services/dbservice/service/schedule"
	mock_sqloperator "server/tests/mocks"
)

func TestGetScheduleControllerFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	scheduleControllerFuncFactory := &scheduleControllerFuncFactory{}
	const name string = "scheduleControllerFunc"
	const errName string = "errName"

	// scheduleControllerFunc
	scheduleControllerFunc, err := scheduleControllerFuncFactory.getScheduleControllerFunc(mockDB, name)
	if err == nil {
		t.Logf("scheduleControllerFunc passed: %v, %v", scheduleControllerFunc, err)
	} else {
		t.Errorf("scheduleControllerFunc failed: %v, %v", scheduleControllerFunc, err)
	}

	// scheduleControllerFunc Name Err
	scheduleControllerFuncNameErr, err := scheduleControllerFuncFactory.getScheduleControllerFunc(mockDB, errName)
	if err != nil {
		t.Logf("scheduleControllerFuncNameErr passed: %v, %v", scheduleControllerFuncNameErr, err)
	} else {
		t.Errorf("scheduleControllerFuncNameErr failed: %v, %v", scheduleControllerFuncNameErr, err)
	}
}

func TestCreateController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	controllerFunc := &scheduleControllerFunc{mockDB, &scheduleservice.ScheduleServiceFactory{}}
	createController := controllerFunc.createController()
	const sql string = `INSERT INTO mysql_db.schedule(
		title,
		start_time,
		end_time,
		note,
		user_id
		) values(?,?,?,?,?);`
	var schedule = dto.Schedule{
		Title:     "title",
		StartTime: "2023-06-06 09:00:00",
		EndTime:   "2023-07-10 15:30:00",
		Note:      "note",
		UserId:    17,
	}
	payload := map[string]interface{}{
		"title":      "title",
		"start_time": "2023-06-06 09:00:00",
		"end_time":   "2023-07-10 15:30:00",
		"note":       "note",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("Failed to marshal payload:", err)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.UserId,
	).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/schedule/create?user_id=17", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	createController(responseRecorder, req)
	if responseRecorder.Code == 201 {
		t.Logf("createController passed: expected status code %d, and got %d", 201, responseRecorder.Code)
	} else {
		t.Errorf("createController failed: expected status code %d, but got %d", 201, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.UserId,
	).Return(nil, fmt.Errorf("no result"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("POST", "/schedule/create?user_id=17", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	createController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("createControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("createControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestUpdateController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	controllerFunc := &scheduleControllerFunc{mockDB, &scheduleservice.ScheduleServiceFactory{}}
	updateController := controllerFunc.updateController()
	const sql string = `UPDATE mysql_db.schedule SET 
	title=?,
	start_time=?,
	end_time=?,
	note=?
	WHERE schedule_id=?;`
	var schedule = dto.Schedule{
		Title:      "title",
		StartTime:  "2023-06-06 09:00:00",
		EndTime:    "2023-07-10 15:30:00",
		Note:       "note",
		ScheduleId: 13,
	}
	payload := map[string]interface{}{
		"title":      "title",
		"start_time": "2023-06-06 09:00:00",
		"end_time":   "2023-07-10 15:30:00",
		"note":       "note",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("Failed to marshal payload:", err)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.ScheduleId,
	).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/schedule/update?schedule_id=13", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	updateController(responseRecorder, req)
	if responseRecorder.Code == 204 {
		t.Logf("updateController passed: expected status code %d, and got %d", 204, responseRecorder.Code)
	} else {
		t.Errorf("updateController failed: Expected status code %d, but got %d", 204, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.ScheduleId,
	).Return(nil, fmt.Errorf("no result"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("PUT", "/schedule/update?schedule_id=13", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	updateController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("updateControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("updateControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestGetAnEventController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	controllerFunc := &scheduleControllerFunc{mockDB, &scheduleservice.ScheduleServiceFactory{}}
	getAnEventController := controllerFunc.getAnEventController()
	const sql string = `SELECT 
	title,
	start_time,
	end_time,
	note,
	user_id
	FROM mysql_db.schedule WHERE schedule_id=?;`
	var schedule = dto.Schedule{ScheduleId: 13}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().QueryRow(sql, schedule.ScheduleId).Return(mockRow)
	mockRow.EXPECT().Scan(
		&schedule.Title,
		&schedule.StartTime,
		&schedule.EndTime,
		&schedule.Note,
		&schedule.UserId,
	).Return(nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/schedule/get?schedule_id=13", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getAnEventController(responseRecorder, req)
	if responseRecorder.Code == 200 {
		t.Logf("getAnEventController passed: expected status code %d, and got %d", 200, responseRecorder.Code)
	} else {
		t.Errorf("getAnEventController failed: Expected status code %d, but got %d", 200, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().QueryRow(sql, schedule.ScheduleId).Return(mockRow)
	mockRow.EXPECT().Scan(
		&schedule.Title,
		&schedule.StartTime,
		&schedule.EndTime,
		&schedule.Note,
		&schedule.UserId,
	).Return(fmt.Errorf("Scan error"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("GET", "/schedule/get?schedule_id=13", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getAnEventController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("getAnEventControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("getAnEventControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestGetOneDayEventsController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	controllerFunc := &scheduleControllerFunc{mockDB, &scheduleservice.ScheduleServiceFactory{}}
	getOneDayEventsController := controllerFunc.getOneDayEventsController()
	const userId uint64 = 13
	const day string = "2023-07-10"
	var sql string = fmt.Sprintf(`SELECT * FROM mysql_db.schedule WHERE user_id=%d AND start_time <= '%s 23:59:59' AND end_time >= '%s 00:00:00';`, userId, day, day)

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Query(sql).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return([]string{
		"schedule_id",
		"title",
		"note",
		"start_time",
		"end_time",
		"user_id",
		"created_time",
		"updated_time",
	}, nil)
	mockRows.EXPECT().Next().Return(false)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/schedule/get_day?user_id=13&day=2023-07-10", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getOneDayEventsController(responseRecorder, req)
	if responseRecorder.Code == 200 {
		t.Logf("getOneDayEventsController passed: expected status code %d, and got %d", 200, responseRecorder.Code)
	} else {
		t.Errorf("getOneDayEventsController failed: Expected status code %d, but got %d", 200, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Query(sql).Return(nil, fmt.Errorf("Query Error"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("GET", "/schedule/get_day?user_id=13&day=2023-07-10", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getOneDayEventsController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("getOneDayEventsControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("getOneDayEventsControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestGetOneMonthEventsController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	controllerFunc := &scheduleControllerFunc{mockDB, &scheduleservice.ScheduleServiceFactory{}}
	getOneMonthEventsController := controllerFunc.getOneMonthEventsController()
	const userId uint64 = 13
	const month string = "202307"
	var sql string = fmt.Sprintf(`SELECT * FROM mysql_db.schedule WHERE user_id=%d AND (%s BETWEEN YEAR(start_time)*100 + MONTH(start_time) AND YEAR(end_time)*100 + MONTH(end_time));`, userId, month)

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Query(sql).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return([]string{
		"schedule_id",
		"title",
		"note",
		"start_time",
		"end_time",
		"user_id",
		"created_time",
		"updated_time",
	}, nil)
	mockRows.EXPECT().Next().Return(false)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/schedule/get_month?user_id=13&month=202307", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getOneMonthEventsController(responseRecorder, req)
	if responseRecorder.Code == 200 {
		t.Logf("getOneMonthEventsController passed: expected status code %d, and got %d", 200, responseRecorder.Code)
	} else {
		t.Errorf("getOneMonthEventsController failed: Expected status code %d, but got %d", 200, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Query(sql).Return(nil, fmt.Errorf("Query Error"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("GET", "/schedule/get_month?user_id=13&month=202307", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getOneMonthEventsController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("getOneMonthEventsControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("getOneMonthEventsControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestDeleteController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	controllerFunc := &scheduleControllerFunc{mockDB, &scheduleservice.ScheduleServiceFactory{}}
	deleteController := controllerFunc.deleteController()
	const sql string = `DELETE FROM mysql_db.schedule WHERE schedule_id=?;`
	const scheduleId uint64 = 7

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql, scheduleId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/schedule/delete?schedule_id=7", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	deleteController(responseRecorder, req)
	if responseRecorder.Code == 204 {
		t.Logf("deleteController passed: expected status code %d, and got %d", 204, responseRecorder.Code)
	} else {
		t.Errorf("deleteController failed: Expected status code %d, but got %d", 204, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql, scheduleId).Return(nil, fmt.Errorf("res err"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("DELETE", "/spirometer/delete?schedule_id=7", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	deleteController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("deleteControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("deleteControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}
