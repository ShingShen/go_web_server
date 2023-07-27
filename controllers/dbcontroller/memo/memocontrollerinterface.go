package memocontroller

import (
	"net/http"
	sqlOperator "server/utils/sqloperator"
)

type IMemoControllerFactory interface {
	GetMemoController(db sqlOperator.ISqlDB, name string) (IMemoController, error)
}

type IMemoController interface {
	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Get() http.HandlerFunc
	GetMemosByUserId() http.HandlerFunc
	Delete() http.HandlerFunc
}

type IMemoControllerFuncFactory interface {
	getMemoControllerFunc(db sqlOperator.ISqlDB, name string) (IMemoControllerFunc, error)
}

type IMemoControllerFunc interface {
	createController() func(w http.ResponseWriter, r *http.Request)
	updateController() func(w http.ResponseWriter, r *http.Request)
	getController() func(w http.ResponseWriter, r *http.Request)
	getMemosByUserIdController() func(w http.ResponseWriter, r *http.Request)
	deleteController() func(w http.ResponseWriter, r *http.Request)
}
