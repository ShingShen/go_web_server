package schedulecontroller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/dto"
	scheduleService "server/services/dbservice/service/schedule"
	sqlOperator "server/utils/sqloperator"
	"strconv"
)

type scheduleControllerFuncFactory struct{}

func (s *scheduleControllerFuncFactory) getScheduleControllerFunc(db sqlOperator.ISqlDB, name string) (IScheduleControllerFunc, error) {
	if name == "scheduleControllerFunc" {
		return &scheduleControllerFunc{
			db:                     db,
			scheduleServiceFactory: &scheduleService.ScheduleServiceFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong schedule controller func type passed")
}

type scheduleControllerFunc struct {
	db                     sqlOperator.ISqlDB
	scheduleServiceFactory scheduleService.IScheduleServiceFactory
}

func (s *scheduleControllerFunc) createController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)

		getService, _ := s.scheduleServiceFactory.GetScheduleService(s.db, "schedule")
		service := getService.CreateEvent(
			dto.Schedule{
				Title:     res["title"].(string),
				Note:      res["note"].(string),
				StartTime: res["start_time"].(string),
				EndTime:   res["end_time"].(string),
				UserId:    uint64(userId),
			},
		)
		if service != nil {
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(201)
	}
}

func (s *scheduleControllerFunc) updateController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		scheduleId, _ := strconv.Atoi(r.URL.Query().Get("schedule_id"))
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)

		getService, _ := s.scheduleServiceFactory.GetScheduleService(s.db, "schedule")
		service := getService.UpdateEvent(
			dto.Schedule{
				Title:      res["title"].(string),
				Note:       res["note"].(string),
				StartTime:  res["start_time"].(string),
				EndTime:    res["end_time"].(string),
				ScheduleId: uint64(scheduleId),
			},
		)
		if service != nil {
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(204)
	}
}

func (s *scheduleControllerFunc) getAnEventController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		scheduleId, _ := strconv.Atoi(r.URL.Query().Get("schedule_id"))

		getService, _ := s.scheduleServiceFactory.GetScheduleService(s.db, "schedule")
		jsonData, err := getService.GetAnEvent(
			dto.Schedule{
				ScheduleId: uint64(scheduleId),
			},
		)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		w.Write(jsonData)
		w.WriteHeader(200)
	}
}

func (s *scheduleControllerFunc) getOneDayEventsController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		day := r.URL.Query().Get("day")

		getService, _ := s.scheduleServiceFactory.GetScheduleService(s.db, "schedule")
		list, err := getService.GetOnedayEvents(uint64(userId), day)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		w.Write(list)
		w.WriteHeader(200)
	}
}

func (s *scheduleControllerFunc) getOneMonthEventsController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		month := r.URL.Query().Get("month")

		getService, _ := s.scheduleServiceFactory.GetScheduleService(s.db, "schedule")
		list, err := getService.GetOneMonthEvents(uint64(userId), month)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		w.Write(list)
		w.WriteHeader(200)
	}
}

func (s *scheduleControllerFunc) deleteController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		scheduleId, _ := strconv.Atoi(r.URL.Query().Get("schedule_id"))

		getService, _ := s.scheduleServiceFactory.GetScheduleService(s.db, "schedule")
		service := getService.DeleteEvent(uint64(scheduleId))
		if service != nil {
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(204)
	}
}
