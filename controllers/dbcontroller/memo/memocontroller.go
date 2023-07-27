package memocontroller

import (
	"fmt"
	"net/http"
	"server/middleware"

	sqlOperator "server/utils/sqloperator"
)

type MemoControllerFactory struct{}

func (m *MemoControllerFactory) GetMemoController(db sqlOperator.ISqlDB, name string) (IMemoController, error) {
	if name == "memoController" {
		return &MemoController{
			db:                        db,
			memoControllerFuncFactory: &memoControllerFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong memo controller type passed")
}

type MemoController struct {
	db                        sqlOperator.ISqlDB
	memoControllerFuncFactory IMemoControllerFuncFactory
}

func (m *MemoController) Create() http.HandlerFunc {
	memoControllerFunc, _ := m.memoControllerFuncFactory.getMemoControllerFunc(m.db, "memoControllerFunc")
	controller := memoControllerFunc.createController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"POST", middleware.AllUserAuth(m.db, controller),
			),
		),
	)
}

func (m *MemoController) Update() http.HandlerFunc {
	memoControllerFunc, _ := m.memoControllerFuncFactory.getMemoControllerFunc(m.db, "memoControllerFunc")
	controller := memoControllerFunc.updateController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"UPDATE", middleware.AllUserAuth(m.db, controller),
			),
		),
	)
}

func (m *MemoController) Get() http.HandlerFunc {
	memoControllerFunc, _ := m.memoControllerFuncFactory.getMemoControllerFunc(m.db, "memoControllerFunc")
	controller := memoControllerFunc.getController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"GET", middleware.AllUserAuth(m.db, controller),
			),
		),
	)
}

func (m *MemoController) GetMemosByUserId() http.HandlerFunc {
	memoControllerFunc, _ := m.memoControllerFuncFactory.getMemoControllerFunc(m.db, "memoControllerFunc")
	controller := memoControllerFunc.getMemosByUserIdController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"GET", middleware.AllUserAuth(m.db, controller),
			),
		),
	)
}

func (m *MemoController) Delete() http.HandlerFunc {
	memoControllerFunc, _ := m.memoControllerFuncFactory.getMemoControllerFunc(m.db, "memoControllerFunc")
	controller := memoControllerFunc.deleteController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"DELETE", middleware.AllUserAuth(m.db, controller),
			),
		),
	)
}
