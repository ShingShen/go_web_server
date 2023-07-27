package apipermission

import (
	"fmt"
	"net/http"

	sqlOperator "server/utils/sqloperator"
)

type IPermission interface {
	LoginAuth(w http.ResponseWriter, r *http.Request, db sqlOperator.ISqlDB, role uint8) (bool, error)
	AllUserAuth(w http.ResponseWriter, r *http.Request, db sqlOperator.ISqlDB) (bool, error)
	UserOnlyAuth(w http.ResponseWriter, r *http.Request, db sqlOperator.ISqlDB, role uint8) (bool, error)
}

func ApiPermission(name string) (IPermission, error) {
	if name == "authCheck" {
		return &AuthCheck{}, nil
	}
	return nil, fmt.Errorf("wrong permission type passed")
}
