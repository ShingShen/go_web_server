package apipermission

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"server/dto"
	"server/utils/encryption"
	"strings"

	sqlOperator "server/utils/sqloperator"
)

type AuthCheck struct{}

func (a *AuthCheck) LoginAuth(w http.ResponseWriter, r *http.Request, db sqlOperator.ISqlDB, role uint8) (bool, error) {
	var user dto.User
	authorization := r.Header.Get("Authorization")

	loginAuthParser, _ := encryption.LoginAuthParser(authorization)
	authAccount := loginAuthParser[0]["auth_account"].(string)
	encodeAuthPassword := loginAuthParser[0]["encode_auth_password"].(string)

	checkingAccountAndPassword := `SELECT user_account FROM mysql_db.user WHERE user_account=? AND user_password=? AND role=?;`
	checkingAccountAndPasswordRes := db.QueryRow(checkingAccountAndPassword, authAccount, encodeAuthPassword, role)
	checkingAccountAndPasswordResScan := checkingAccountAndPasswordRes.Scan(&user.UserAccount)
	if checkingAccountAndPasswordResScan != nil {
		fmt.Println("Checked user account, Error: ", checkingAccountAndPasswordResScan)
		w.WriteHeader(401)
		return false, checkingAccountAndPasswordResScan
	}
	return true, checkingAccountAndPasswordResScan
}

func (a *AuthCheck) AllUserAuth(w http.ResponseWriter, r *http.Request, db sqlOperator.ISqlDB) (bool, error) {
	var user dto.User
	authorization := r.Header.Get("Authorization")

	authorizationContent := strings.Split(authorization, "Basic ")
	if len(authorizationContent) != 2 {
		w.WriteHeader(401)
		return false, fmt.Errorf("authorization format is incorrect")
	}
	decodeAuthorization, _ := base64.StdEncoding.DecodeString(authorizationContent[1])
	loginToken := string(decodeAuthorization)

	checkingAuth := `SELECT user_id FROM mysql_db.login_token WHERE login_token=?;`
	checkingAuthRes := db.QueryRow(checkingAuth, loginToken)
	checkingAuthResScan := checkingAuthRes.Scan(&user.UserAccount)
	fmt.Println("checkingAuthResScan: ", checkingAuthResScan)
	if checkingAuthResScan != nil {
		fmt.Println("Checked user account, Error: ", checkingAuthResScan)
		w.WriteHeader(401)
		return false, checkingAuthResScan
	}
	return true, checkingAuthResScan
}

func (a *AuthCheck) UserOnlyAuth(w http.ResponseWriter, r *http.Request, db sqlOperator.ISqlDB, role uint8) (bool, error) {
	var user dto.User
	authorization := r.Header.Get("Authorization")

	authorizationContent := strings.Split(authorization, "Basic ")
	if len(authorizationContent) != 2 {
		w.WriteHeader(401)
		return false, fmt.Errorf("authorization format is incorrect")
	}
	decodeAuthorization, _ := base64.StdEncoding.DecodeString(authorizationContent[1])
	loginToken := string(decodeAuthorization)

	checkingLoginToken, _ := a.checkingLoginToken(db, w, loginToken, user)
	if !checkingLoginToken {
		return false, fmt.Errorf("login token is incorrect")
	}

	checkingRole := `SELECT user_account FROM mysql_db.user
	INNER JOIN mysql_db.login_token ON mysql_db.user.user_id = mysql_db.login_token.user_id
	WHERE mysql_db.login_token.login_token=? AND mysql_db.user.role=?;`

	// checkingRole := `SELECT user_account FROM mysql_db.user WHERE login_token=? AND role=?;`
	checkingRoleRes := db.QueryRow(checkingRole, loginToken, role)
	checkingRoleResScan := checkingRoleRes.Scan(&user.UserAccount)
	fmt.Println("checkingAuthResScan: ", checkingRoleResScan)
	if checkingRoleResScan != nil {
		fmt.Println("Checked user account, Error: ", checkingRoleResScan)
		w.WriteHeader(403)
		return false, checkingRoleResScan
	}

	return true, checkingRoleResScan
}

func (a *AuthCheck) checkingLoginToken(db sqlOperator.ISqlDB, w http.ResponseWriter, loginToken string, user dto.User) (bool, error) {
	checkingLoginToken := `SELECT user_id FROM mysql_db.login_token WHERE login_token=?;`
	checkingLoginTokenRes := db.QueryRow(checkingLoginToken, loginToken)
	checkingLoginTokenResScan := checkingLoginTokenRes.Scan(&user.UserAccount)
	fmt.Println("checkingLoginTokenResScan: ", checkingLoginTokenResScan)
	if checkingLoginTokenResScan != nil {
		fmt.Println("Checked user account, Error: ", checkingLoginTokenResScan)
		w.WriteHeader(401)
		return false, checkingLoginTokenResScan
	}
	return true, checkingLoginTokenResScan
}
