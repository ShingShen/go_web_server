package scheduleserviceprovider

import (
	"fmt"
	"testing"

	"server/dto"
	mock_sqloperator "server/tests/mocks"

	"github.com/golang/mock/gomock"
)

func TestGetScheduleServiceFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	scheduleServiceFuncFactory := &ScheduleServiceFuncFactory{}
	const name string = "scheduleServiceFunc"
	const errName string = "errName"

	// scheduleServiceFunc
	scheduleServiceFunc, err := scheduleServiceFuncFactory.GetScheduleServiceFunc(mockDB, name)
	if err == nil {
		t.Logf("scheduleServiceFunc passed: %v, %v", scheduleServiceFunc, err)
	} else {
		t.Errorf("scheduleServiceFunc failed: %v, %v", scheduleServiceFunc, err)
	}

	// scheduleServiceFunc Name Err
	scheduleServiceFuncNameErr, err := scheduleServiceFuncFactory.GetScheduleServiceFunc(mockDB, errName)
	if err != nil {
		t.Logf("scheduleServiceFuncNameErr passed: %v, %v", scheduleServiceFuncNameErr, err)
	} else {
		t.Errorf("scheduleServiceFuncNameErr failed: %v, %v", scheduleServiceFuncNameErr, err)
	}
}

func TestCreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `INSERT INTO mysql_db.schedule(
		title,
		start_time,
		end_time,
		note,
		user_id
		) values(?,?,?,?,?);`
	var schedule dto.Schedule
	scheduleServiceFunc := &ScheduleServiceFunc{mockDB}

	// createEvent Res Err
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.UserId,
	).Return(nil, fmt.Errorf("res err"))
	createEventResErr := scheduleServiceFunc.CreateEvent(schedule)
	if createEventResErr != nil {
		t.Logf("createEventResErr passed: %v", createEventResErr)
	} else {
		t.Errorf("createEventResErr failed: %v", createEventResErr)
	}

	// createEvent Rows Affected Err
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.UserId,
	).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	createEventRowsAffectedErr := scheduleServiceFunc.CreateEvent(schedule)
	if createEventRowsAffectedErr != nil {
		t.Logf("createEventRowsAffectedErr passed: %v", createEventRowsAffectedErr)
	} else {
		t.Errorf("createEventRowsAffectedErr failed: %v", createEventRowsAffectedErr)
	}

	// createEvent
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.UserId,
	).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	createEvent := scheduleServiceFunc.CreateEvent(schedule)
	if createEvent == nil {
		t.Logf("createEvent passed: %v", createEvent)
	} else {
		t.Errorf("createEvent failed: %v", createEvent)
	}
}

func TestUpdateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `UPDATE mysql_db.schedule SET 
	title=?,
	start_time=?,
	end_time=?,
	note=?
	WHERE schedule_id=?;`
	var schedule dto.Schedule
	scheduleServiceFunc := &ScheduleServiceFunc{mockDB}

	// updateEvent Res Err
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.ScheduleId,
	).Return(nil, fmt.Errorf("no result"))
	updateEventResErr := scheduleServiceFunc.UpdateEvent(schedule)
	if updateEventResErr != nil {
		t.Logf("updateEventResErr passed: %v", updateEventResErr)
	} else {
		t.Errorf("updateEventResErr failed: %v", updateEventResErr)
	}

	// updateEvent Rows Affected Err
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.ScheduleId,
	).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	updateEventRowsAffectedErr := scheduleServiceFunc.UpdateEvent(schedule)
	if updateEventRowsAffectedErr != nil {
		t.Logf("updateEventRowsAffectedErr passed: %v", updateEventRowsAffectedErr)
	} else {
		t.Errorf("updateEventRowsAffectedErr failed: %v", updateEventRowsAffectedErr)
	}

	// updateEvent
	mockDB.EXPECT().Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.ScheduleId,
	).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	updateEvent := scheduleServiceFunc.UpdateEvent(schedule)
	if updateEvent == nil {
		t.Logf("updateEvent passed: %v", updateEvent)
	} else {
		t.Errorf("updateEvent failed: %v", updateEvent)
	}
}

func TestGetAnEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	const sql string = `SELECT 
	title,
	start_time,
	end_time,
	note,
	user_id
	FROM mysql_db.schedule WHERE schedule_id=?;`
	var schedule dto.Schedule
	scheduleServiceFunc := &ScheduleServiceFunc{mockDB}

	// An event Scan Err
	mockDB.EXPECT().QueryRow(sql, schedule.ScheduleId).Return(mockRow)
	mockRow.EXPECT().Scan(
		&schedule.Title,
		&schedule.StartTime,
		&schedule.EndTime,
		&schedule.Note,
		&schedule.UserId,
	).Return(fmt.Errorf("Scan error"))
	anEventScanErr, err := scheduleServiceFunc.GetAnEvent(schedule)
	if err != nil {
		t.Logf("anEventScanErr passed: %v, %v", anEventScanErr, err)
	} else {
		t.Errorf("anEventScanErr failed: %v, %v", anEventScanErr, err)
	}

	// An event
	mockDB.EXPECT().QueryRow(sql, schedule.ScheduleId).Return(mockRow)
	mockRow.EXPECT().Scan(
		&schedule.Title,
		&schedule.StartTime,
		&schedule.EndTime,
		&schedule.Note,
		&schedule.UserId,
	).Return(nil)
	anEvent, err := scheduleServiceFunc.GetAnEvent(schedule)
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
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	const userId uint64 = 1
	const day string = "2023-07-10"
	var sql string = fmt.Sprintf(`SELECT * FROM mysql_db.schedule WHERE user_id=%d AND start_time <= '%s 23:59:59' AND end_time >= '%s 00:00:00';`, userId, day, day)
	scheduleServiceFunc := &ScheduleServiceFunc{mockDB}

	// oneDayEvents Query Err
	mockDB.EXPECT().Query(sql).Return(nil, fmt.Errorf("Query Error"))
	oneDayEventsQueryErr, err := scheduleServiceFunc.GetOneDayEvents(userId, day)
	if err != nil {
		t.Logf("oneDayEventsQueryErr passed: %v, %v", oneDayEventsQueryErr, err)
	} else {
		t.Errorf("oneDayEventsQueryErr failed: %v, %v", oneDayEventsQueryErr, err)
	}

	// oneDayEvents Col Err
	mockDB.EXPECT().Query(sql).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return(nil, fmt.Errorf("Columns error"))
	oneDayEventsColErr, err := scheduleServiceFunc.GetOneDayEvents(userId, day)
	if err != nil {
		t.Logf("oneDayEventsColErr passed: %v, %v", oneDayEventsColErr, err)
	} else {
		t.Errorf("oneDayEventsColErr failed: %v, %v", oneDayEventsColErr, err)
	}

	// oneDayEvents
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
	oneDayEvents, err := scheduleServiceFunc.GetOneDayEvents(userId, day)
	if err == nil {
		t.Logf("oneDayEvents passed: %v, %v", oneDayEvents, err)
	} else {
		t.Errorf("oneDayEvents failed: %v, %v", oneDayEvents, err)
	}
}

func TestGetOneMonthEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	const userId uint64 = 1
	const month string = "202307"
	var sql string = fmt.Sprintf(`SELECT * FROM mysql_db.schedule WHERE user_id=%d AND (%s BETWEEN YEAR(start_time)*100 + MONTH(start_time) AND YEAR(end_time)*100 + MONTH(end_time));`, userId, month)
	scheduleServiceFunc := &ScheduleServiceFunc{mockDB}

	// oneMonthEvents Query Err
	mockDB.EXPECT().Query(sql).Return(nil, fmt.Errorf("Query Error"))
	oneMonthEventsQueryErr, err := scheduleServiceFunc.GetOneMonthEvents(userId, month)
	if err != nil {
		t.Logf("oneMonthEventsQueryErr passed: %v, %v", oneMonthEventsQueryErr, err)
	} else {
		t.Errorf("oneMonthEventsQueryErr failed: %v, %v", oneMonthEventsQueryErr, err)
	}

	// oneMonthEvents Col Err
	mockDB.EXPECT().Query(sql).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return(nil, fmt.Errorf("Columns error"))
	oneMonthEventsColErr, err := scheduleServiceFunc.GetOneMonthEvents(userId, month)
	if err != nil {
		t.Logf("oneMonthEventsColErr passed: %v, %v", oneMonthEventsColErr, err)
	} else {
		t.Errorf("oneMonthEventsColErr failed: %v, %v", oneMonthEventsColErr, err)
	}

	// oneMonthEvents
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
	oneMonthEvents, err := scheduleServiceFunc.GetOneMonthEvents(userId, month)
	if err == nil {
		t.Logf("oneMonthEvents passed: %v, %v", oneMonthEvents, err)
	} else {
		t.Errorf("oneMonthEvents failed: %v, %v", oneMonthEvents, err)
	}
}

func TestDeleteEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `DELETE FROM mysql_db.schedule WHERE schedule_id=?;`
	const scheduleId uint64 = 1
	scheduleServiceFunc := &ScheduleServiceFunc{mockDB}

	// deleteEvent res err
	mockDB.EXPECT().Exec(sql, scheduleId).Return(nil, fmt.Errorf("res err"))
	deleteEventResErr := scheduleServiceFunc.DeleteEvent(scheduleId)
	if deleteEventResErr != nil {
		t.Logf("deleteEventResErr passed: %v", deleteEventResErr)
	} else {
		t.Errorf("deleteEventResErr failed: %v", deleteEventResErr)
	}

	// deleteEvent rows affected err
	mockDB.EXPECT().Exec(sql, scheduleId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("rows affected error"))
	deleteEventRowsAffectedErr := scheduleServiceFunc.DeleteEvent(scheduleId)
	if deleteEventRowsAffectedErr != nil {
		t.Logf("deleteEventRowsAffectedErr passed: %v", deleteEventRowsAffectedErr)
	} else {
		t.Errorf("deleteEventRowsAffectedErr failed: %v", deleteEventRowsAffectedErr)
	}

	// deleteEvent
	mockDB.EXPECT().Exec(sql, scheduleId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	deleteEvent := scheduleServiceFunc.DeleteEvent(scheduleId)
	if deleteEvent == nil {
		t.Logf("deleteEvent passed: %v", deleteEvent)
	} else {
		t.Errorf("deleteEvent failed: %v", deleteEvent)
	}
}
