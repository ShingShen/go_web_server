package scheduleservice

import (
	"server/dto"
	sqlOperator "server/utils/sqloperator"
)

type IScheduleServiceFactory interface {
	GetScheduleService(db sqlOperator.ISqlDB, name string) (IScheduleService, error)
}

type IScheduleService interface {
	CreateEvent(schedule dto.Schedule) error
	UpdateEvent(schedule dto.Schedule) error
	GetAnEvent(schedule dto.Schedule) ([]byte, error)
	GetOnedayEvents(userId uint64, day string) ([]byte, error)
	GetOneMonthEvents(userId uint64, month string) ([]byte, error)
	DeleteEvent(scheduleId uint64) error
}
