package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	apiPermission "server/utils/apipermission"
	sqlOperator "server/utils/sqloperator"
)

func AllUserAuth(db sqlOperator.ISqlDB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		permission, _ := apiPermission.ApiPermission("authCheck")
		auth, _ := permission.AllUserAuth(w, r, db)
		if !auth {
			return
		}
		next(w, r)
	}
}

func LoginAuth(db sqlOperator.ISqlDB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)
		permission, _ := apiPermission.ApiPermission("authCheck")
		auth, _ := permission.LoginAuth(w, r, db, uint8(res["role"].(float64)))
		if !auth {
			return
		}
		next(w, r)
	}
}

func UserOnlyAuth(role uint8, db sqlOperator.ISqlDB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		permission, _ := apiPermission.ApiPermission("authCheck")
		auth, _ := permission.UserOnlyAuth(w, r, db, role)
		if !auth {
			return
		}
		next(w, r)
	}
}
