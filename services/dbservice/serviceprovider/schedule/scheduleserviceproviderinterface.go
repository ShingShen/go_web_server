package scheduleserviceprovider

import (
	"server/dto"
	sqlOperator "server/utils/sqloperator"
)

type IScheduleServiceFuncFactory interface {
	GetScheduleServiceFunc(db sqlOperator.ISqlDB, name string) (IScheduleServiceFunc, error)
}

type IScheduleServiceFunc interface {
	CreateEvent(schedule dto.Schedule) error
	UpdateEvent(schedule dto.Schedule) error
	GetAnEvent(schedule dto.Schedule) ([]byte, error)
	GetOneDayEvents(userId uint64, day string) ([]byte, error)
	GetOneMonthEvents(userId uint64, month string) ([]byte, error)
	DeleteEvent(scheduleId uint64) error
}
