package memocontroller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	memoService "server/services/dbservice/service/memo"
	sqlOperator "server/utils/sqloperator"
)

type memoControllerFuncFactory struct{}

func (m *memoControllerFuncFactory) getMemoControllerFunc(db sqlOperator.ISqlDB, name string) (IMemoControllerFunc, error) {
	if name == "memoControllerFunc" {
		return &memoControllerFunc{
			db:                 db,
			memoServiceFactory: &memoService.MemoServiceFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong memo controller func type passed")
}

type memoControllerFunc struct {
	db                 sqlOperator.ISqlDB
	memoServiceFactory memoService.IMemoServiceFactory
}

func (m *memoControllerFunc) createController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)

		getService, _ := m.memoServiceFactory.GetMemoService(m.db, "memo")
		service := getService.CreateMemo(
			res["content"].(string),
			uint64(userId),
		)
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(204)
	}
}

func (m *memoControllerFunc) updateController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		memoId, _ := strconv.Atoi(r.URL.Query().Get("memo_id"))
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)

		getService, _ := m.memoServiceFactory.GetMemoService(m.db, "memo")
		service := getService.UpdateMemo(
			uint64(memoId),
			res["content"].(string),
		)
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(204)
	}
}

func (m *memoControllerFunc) getController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		memoId, _ := strconv.Atoi(r.URL.Query().Get("memo_id"))

		getService, _ := m.memoServiceFactory.GetMemoService(m.db, "memo")
		jsonData, err := getService.GetMemo(uint64(memoId))
		if err != nil {
			w.WriteHeader(400)
			return
		}

		w.Write(jsonData)
		w.WriteHeader(200)
	}
}

func (m *memoControllerFunc) getMemosByUserIdController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))

		getService, _ := m.memoServiceFactory.GetMemoService(m.db, "memo")
		jsonData, err := getService.GetMemosByUserId(uint64(userId))
		if err != nil {
			w.WriteHeader(400)
			return
		}

		w.Write(jsonData)
		w.WriteHeader(200)
	}
}

func (m *memoControllerFunc) deleteController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		memoId, _ := strconv.Atoi(r.URL.Query().Get("memo_id"))

		getService, _ := m.memoServiceFactory.GetMemoService(m.db, "memo")
		service := getService.DeleteMemo(uint64(memoId))
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(204)
	}
}
