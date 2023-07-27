package schedulecontroller

import (
	"net/http"
	sqlOperator "server/utils/sqloperator"
)

type IScheduleControllerFactory interface {
	GetScheduleController(db sqlOperator.ISqlDB, name string) (IScheduleController, error)
}

type IScheduleController interface {
	Create() http.HandlerFunc
	Update() http.HandlerFunc
	GetAnEvent() http.HandlerFunc
	GetOneDayEvents() http.HandlerFunc
	GetOneMonthEvents() http.HandlerFunc
	Delete() http.HandlerFunc
}

type IScheduleControllerFuncFactory interface {
	getScheduleControllerFunc(db sqlOperator.ISqlDB, name string) (IScheduleControllerFunc, error)
}

type IScheduleControllerFunc interface {
	createController() func(w http.ResponseWriter, r *http.Request)
	updateController() func(w http.ResponseWriter, r *http.Request)
	getAnEventController() func(w http.ResponseWriter, r *http.Request)
	getOneDayEventsController() func(w http.ResponseWriter, r *http.Request)
	getOneMonthEventsController() func(w http.ResponseWriter, r *http.Request)
	deleteController() func(w http.ResponseWriter, r *http.Request)
}
