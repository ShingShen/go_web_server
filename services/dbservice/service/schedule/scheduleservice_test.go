package scheduleservice

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"server/dto"
	scheduleserviceprovider "server/services/dbservice/serviceprovider/schedule"
	mock_sqloperator "server/tests/mocks"
)

func TestGetScheduleService(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	scheduleServiceFactory := &ScheduleServiceFactory{}
	const name string = "schedule"
	const errName string = "errName"

	scheduleService, err := scheduleServiceFactory.GetScheduleService(mockDB, name)
	if err == nil {
		t.Logf("scheduleService passed: %v, %v", scheduleService, err)
	} else {
		t.Errorf("scheduleService failed: %v, %v", scheduleService, err)
	}

	scheduleServiceNameErr, err := scheduleServiceFactory.GetScheduleService(mockDB, errName)
	if err != nil {
		t.Logf("scheduleServiceNameErr passed: %v, %v", scheduleServiceNameErr, err)
	} else {
		t.Errorf("scheduleServiceNameErr failed: %v, %v", scheduleServiceNameErr, err)
	}
}

func TestCreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	var schedule dto.Schedule
	const sql string = `INSERT INTO mysql_db.schedule(
		title,
		start_time,
		end_time,
		note,
		user_id
		) values(?,?,?,?,?);`
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.UserId,
	).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)

	scheduleService := &ScheduleService{mockDB, &scheduleserviceprovider.ScheduleServiceFuncFactory{}}
	createEvent := scheduleService.CreateEvent(schedule)
	if createEvent == nil {
		t.Logf("createEvent passed: %v", createEvent)
	} else {
		t.Errorf("createEvent failed: %v", createEvent)
	}
}

func TestUpdateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	var schedule dto.Schedule
	const sql string = `UPDATE mysql_db.schedule SET 
	title=?,
	start_time=?,
	end_time=?,
	note=?
	WHERE schedule_id=?;`
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.ScheduleId,
	).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)

	scheduleService := &ScheduleService{mockDB, &scheduleserviceprovider.ScheduleServiceFuncFactory{}}
	updateEvent := scheduleService.UpdateEvent(schedule)
	if updateEvent == nil {
		t.Logf("updateEvent passed: %v", updateEvent)
	} else {
		t.Errorf("updateEvent failed: %v", updateEvent)
	}
}

func TestGetAnEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	var schedule dto.Schedule
	const sql string = `SELECT 
	title,
	start_time,
	end_time,
	note,
	user_id
	FROM mysql_db.schedule WHERE schedule_id=?;`
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	mockDB.EXPECT().QueryRow(sql, schedule.ScheduleId).Return(mockRow)
	mockRow.EXPECT().Scan(
		&schedule.Title,
		&schedule.StartTime,
		&schedule.EndTime,
		&schedule.Note,
		&schedule.UserId,
	).Return(nil)
	scheduleService := &ScheduleService{mockDB, &scheduleserviceprovider.ScheduleServiceFuncFactory{}}
	anEvent, err := scheduleService.GetAnEvent(schedule)
	if anEvent != nil {
		t.Logf("anEvent passed: %v, %v", anEvent, err)
	}
	if err != nil {
		t.Errorf("anEvent failed: %v, %v", anEvent, err)
	}
}

func TestGetOneDayEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	const userId uint64 = 1
	const day string = "2023-07-10"
	var sql string = fmt.Sprintf(`SELECT * FROM mysql_db.schedule WHERE user_id=%d AND start_time <= '%s 23:59:59' AND end_time >= '%s 00:00:00';`, userId, day, day)

	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
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
	scheduleService := &ScheduleService{mockDB, &scheduleserviceprovider.ScheduleServiceFuncFactory{}}
	oneDayEvents, err := scheduleService.GetOnedayEvents(userId, day)
	if oneDayEvents != nil {
		t.Logf("oneDayEvents passed: %v, %v", oneDayEvents, err)
	}
	if err != nil {
		t.Errorf("oneDayEvents failed: %v, %v", oneDayEvents, err)
	}
}

func TestGetOneMonthEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	const userId uint64 = 1
	const month string = "202307"
	var sql string = fmt.Sprintf(`SELECT * FROM mysql_db.schedule WHERE user_id=%d AND (%s BETWEEN YEAR(start_time)*100 + MONTH(start_time) AND YEAR(end_time)*100 + MONTH(end_time));`, userId, month)

	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
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
	scheduleService := &ScheduleService{mockDB, &scheduleserviceprovider.ScheduleServiceFuncFactory{}}
	oneMonthEvents, err := scheduleService.GetOneMonthEvents(userId, month)
	if oneMonthEvents != nil {
		t.Logf("oneMonthEvents passed: %v, %v", oneMonthEvents, err)
	}
	if err != nil {
		t.Errorf("oneMonthEvents failed: %v, %v", oneMonthEvents, err)
	}
}

func TestDeleteEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	const sql string = `DELETE FROM mysql_db.schedule WHERE schedule_id=?;`
	const scheduleId uint64 = 1

	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	mockDB.EXPECT().Exec(sql, scheduleId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	scheduleService := &ScheduleService{mockDB, &scheduleserviceprovider.ScheduleServiceFuncFactory{}}
	deleteEvent := scheduleService.DeleteEvent(scheduleId)
	if deleteEvent == nil {
		t.Logf("deleteEvent passed: %v", deleteEvent)
	} else {
		t.Errorf("deleteEvent failed: %v", deleteEvent)
	}
}
