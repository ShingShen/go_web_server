package schedulecontroller

import (
	"fmt"
	"net/http"
	"server/middleware"
	sqlOperator "server/utils/sqloperator"
)

type ScheduleControllerFactory struct{}

func (s *ScheduleControllerFactory) GetScheduleController(db sqlOperator.ISqlDB, name string) (IScheduleController, error) {
	if name == "scheduleController" {
		return &ScheduleController{
			db:                            db,
			scheduleControllerFuncFactory: &scheduleControllerFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong schedule controller type passed")
}

type ScheduleController struct {
	db                            sqlOperator.ISqlDB
	scheduleControllerFuncFactory IScheduleControllerFuncFactory
}

func (s *ScheduleController) Create() http.HandlerFunc {
	controllerFunc, _ := s.scheduleControllerFuncFactory.getScheduleControllerFunc(s.db, "scheduleControllerFunc")
	controller := controllerFunc.createController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"POST", middleware.AllUserAuth(s.db, controller),
			),
		),
	)
}

func (s *ScheduleController) Update() http.HandlerFunc {
	controllerFunc, _ := s.scheduleControllerFuncFactory.getScheduleControllerFunc(s.db, "scheduleControllerFunc")
	controller := controllerFunc.updateController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"PUT", middleware.AllUserAuth(s.db, controller),
			),
		),
	)
}

func (s *ScheduleController) GetAnEvent() http.HandlerFunc {
	controllerFunc, _ := s.scheduleControllerFuncFactory.getScheduleControllerFunc(s.db, "scheduleControllerFunc")
	controller := controllerFunc.getAnEventController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"GET", middleware.AllUserAuth(s.db, controller),
			),
		),
	)
}

func (s *ScheduleController) GetOneDayEvents() http.HandlerFunc {
	controllerFunc, _ := s.scheduleControllerFuncFactory.getScheduleControllerFunc(s.db, "scheduleControllerFunc")
	controller := controllerFunc.getOneDayEventsController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"GET", middleware.AllUserAuth(s.db, controller),
			),
		),
	)
}

func (s *ScheduleController) GetOneMonthEvents() http.HandlerFunc {
	controllerFunc, _ := s.scheduleControllerFuncFactory.getScheduleControllerFunc(s.db, "scheduleControllerFunc")
	controller := controllerFunc.getOneMonthEventsController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"GET", middleware.AllUserAuth(s.db, controller),
			),
		),
	)
}

func (s *ScheduleController) Delete() http.HandlerFunc {
	controllerFunc, _ := s.scheduleControllerFuncFactory.getScheduleControllerFunc(s.db, "scheduleControllerFunc")
	controller := controllerFunc.deleteController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"DELETE", middleware.AllUserAuth(s.db, controller),
			),
		),
	)
}
