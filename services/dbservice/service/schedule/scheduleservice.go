package scheduleservice

import (
	"fmt"
	"server/dto"
	scheduleServiceProvider "server/services/dbservice/serviceprovider/schedule"
	sqlOperator "server/utils/sqloperator"
)

type ScheduleServiceFactory struct{}

func (s *ScheduleServiceFactory) GetScheduleService(db sqlOperator.ISqlDB, name string) (IScheduleService, error) {
	if name == "schedule" {
		return &ScheduleService{
			db:                         db,
			scheduleServiceFuncFactory: &scheduleServiceProvider.ScheduleServiceFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong schedule type passed")
}

type ScheduleService struct {
	db                         sqlOperator.ISqlDB
	scheduleServiceFuncFactory scheduleServiceProvider.IScheduleServiceFuncFactory
}

func (s *ScheduleService) CreateEvent(schedule dto.Schedule) error {
	tx, _ := s.db.Begin()
	defer tx.Rollback()

	scheduleServiceFunc, _ := s.scheduleServiceFuncFactory.GetScheduleServiceFunc(s.db, "scheduleServiceFunc")
	return scheduleServiceFunc.CreateEvent(schedule)
}

func (s *ScheduleService) UpdateEvent(schedule dto.Schedule) error {
	tx, _ := s.db.Begin()
	defer tx.Rollback()

	scheduleServiceFunc, _ := s.scheduleServiceFuncFactory.GetScheduleServiceFunc(s.db, "scheduleServiceFunc")
	return scheduleServiceFunc.UpdateEvent(schedule)
}

func (s *ScheduleService) GetAnEvent(schedule dto.Schedule) ([]byte, error) {
	tx, _ := s.db.Begin()
	defer tx.Rollback()

	scheduleServiceFunc, _ := s.scheduleServiceFuncFactory.GetScheduleServiceFunc(s.db, "scheduleServiceFunc")
	return scheduleServiceFunc.GetAnEvent(schedule)
}

func (s *ScheduleService) GetOnedayEvents(userId uint64, day string) ([]byte, error) {
	tx, _ := s.db.Begin()
	defer tx.Rollback()

	scheduleServiceFunc, _ := s.scheduleServiceFuncFactory.GetScheduleServiceFunc(s.db, "scheduleServiceFunc")
	return scheduleServiceFunc.GetOneDayEvents(userId, day)
}

func (s *ScheduleService) GetOneMonthEvents(userId uint64, month string) ([]byte, error) {
	tx, _ := s.db.Begin()
	defer tx.Rollback()

	scheduleServiceFunc, _ := s.scheduleServiceFuncFactory.GetScheduleServiceFunc(s.db, "scheduleServiceFunc")
	return scheduleServiceFunc.GetOneMonthEvents(userId, month)
}

func (s *ScheduleService) DeleteEvent(scheduleId uint64) error {
	tx, _ := s.db.Begin()
	defer tx.Rollback()

	scheduleServiceFunc, _ := s.scheduleServiceFuncFactory.GetScheduleServiceFunc(s.db, "scheduleServiceFunc")
	return scheduleServiceFunc.DeleteEvent(scheduleId)
}
