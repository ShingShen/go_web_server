package scheduleserviceprovider

import (
	"encoding/json"
	"fmt"
	"server/dto"
	"server/utils/operator"
	sqlOperator "server/utils/sqloperator"
)

type ScheduleServiceFuncFactory struct{}

func (s *ScheduleServiceFuncFactory) GetScheduleServiceFunc(db sqlOperator.ISqlDB, name string) (IScheduleServiceFunc, error) {
	if name == "scheduleServiceFunc" {
		return &ScheduleServiceFunc{db}, nil
	}
	return nil, fmt.Errorf("wrong schedule service func type passed")
}

type ScheduleServiceFunc struct {
	db sqlOperator.ISqlDB
}

func (s *ScheduleServiceFunc) CreateEvent(schedule dto.Schedule) error {
	const sql string = `INSERT INTO mysql_db.schedule(
		title,
		start_time,
		end_time,
		note,
		user_id
		) values(?,?,?,?,?);`
	res, err := s.db.Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.UserId,
	)
	if err != nil {
		fmt.Printf("Failed to insert schedule data, err: %v\n", err)
		return err
	}
	createUserRowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get schedule affected rows, err: %v\n", err)
		return err
	}

	fmt.Println("schedule affected rows: ", createUserRowsAffected)
	return nil
}

func (s *ScheduleServiceFunc) UpdateEvent(schedule dto.Schedule) error {
	const sql string = `UPDATE mysql_db.schedule SET 
	title=?,
	start_time=?,
	end_time=?,
	note=?
	WHERE schedule_id=?;`
	res, err := s.db.Exec(
		sql,
		schedule.Title,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Note,
		schedule.ScheduleId,
	)
	if err != nil {
		fmt.Printf("Failed to update schedule data, err: %v", err)
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v", err)
		return err
	}
	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}

func (s *ScheduleServiceFunc) GetAnEvent(schedule dto.Schedule) ([]byte, error) {
	const sql string = `SELECT 
	title,
	start_time,
	end_time,
	note,
	user_id
	FROM mysql_db.schedule WHERE schedule_id=?;`
	res := s.db.QueryRow(sql, schedule.ScheduleId)
	err := res.Scan(
		&schedule.Title,
		&schedule.StartTime,
		&schedule.EndTime,
		&schedule.Note,
		&schedule.UserId,
	)
	if err != nil {
		return nil, err
	} else {
		jsonData, _ := json.Marshal(map[string]interface{}{
			"title":      schedule.Title,
			"start_time": schedule.StartTime,
			"end_time":   schedule.EndTime,
			"note":       schedule.Note,
			"user_id":    schedule.UserId,
		})
		return jsonData, nil
	}
}

func (s *ScheduleServiceFunc) GetOneDayEvents(userId uint64, day string) ([]byte, error) {
	// day: yyyy-mm-dd
	var sql string = fmt.Sprintf(`SELECT * FROM mysql_db.schedule WHERE user_id=%d AND start_time <= '%s 23:59:59' AND end_time >= '%s 00:00:00';`, userId, day, day)
	res, err := s.db.Query(sql)
	if err != nil {
		fmt.Printf("Failed to query a day events, err: %v\n", err)
		return nil, err
	}

	list, err := operator.CreatingDataList(res)
	if err != nil {
		return nil, err
	} else {
		return list, nil
	}
}

func (s *ScheduleServiceFunc) GetOneMonthEvents(userId uint64, month string) ([]byte, error) {
	// month: yyyymm
	var sql string = fmt.Sprintf(`SELECT * FROM mysql_db.schedule WHERE user_id=%d AND (%s BETWEEN YEAR(start_time)*100 + MONTH(start_time) AND YEAR(end_time)*100 + MONTH(end_time));`, userId, month)
	res, err := s.db.Query(sql)
	if err != nil {
		fmt.Printf("Failed to query a month events, err: %v\n", err)
		return nil, err
	}

	list, err := operator.CreatingDataList(res)
	if err != nil {
		return nil, err
	} else {
		return list, nil
	}
}

func (s *ScheduleServiceFunc) DeleteEvent(scheduleId uint64) error {
	const sql string = `DELETE FROM mysql_db.schedule WHERE schedule_id=?;`
	res, err := s.db.Exec(sql, scheduleId)
	if err != nil {
		fmt.Printf("Failed to delete schedule data, err: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v", err)
		return err
	}
	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}
