package usercontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"server/dto"
	userservice "server/services/dbservice/service/user"
	mock_sqloperator "server/tests/mocks"
)

func TestGetUserControllerFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userControllerFuncFactory := &userControllerFuncFactory{}
	const name string = "userControllerFunc"
	const errName string = "errName"

	// userControllerFunc
	userControllerFunc, err := userControllerFuncFactory.getUserControllerFunc(mockDB, name)
	if err == nil {
		t.Logf("userControllerFunc passed: %v, %v", userControllerFunc, err)
	} else {
		t.Errorf("userControllerFunc failed: %v, %v", userControllerFunc, err)
	}

	// userControllerFunc Name Err
	userControllerFuncNameErr, err := userControllerFuncFactory.getUserControllerFunc(mockDB, errName)
	if err != nil {
		t.Logf("userControllerFuncNameErr passed: %v, %v", userControllerFuncNameErr, err)
	} else {
		t.Errorf("userControllerFuncNameErr failed: %v, %v", userControllerFuncNameErr, err)
	}
}

func TestUpdateController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	controllerFunc := &userControllerFunc{mockDB, &userservice.UserServiceFactory{}}
	updateController := controllerFunc.updateController()
	const sql string = `UPDATE mysql_db.user SET first_name=?, last_name=?, email=?, phone=?, gender=?, birthday=? WHERE user_id=?;`
	const userId uint64 = 13
	var user = dto.User{
		FirstName: "Amy",
		LastName:  "Lee",
		Gender:    2,
		Birthday:  "2001-02-13",
		Email:     "amy0213haha@gmail.com",
		Phone:     "0987654321",
	}
	payload := map[string]interface{}{
		"first_name": "Amy",
		"last_name":  "Lee",
		"gender":     2,
		"birthday":   "2001-02-13",
		"email":      "amy0213haha@gmail.com",
		"phone":      "0987654321",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("Failed to marshal payload:", err)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(
		sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Gender,
		user.Birthday,
		userId,
	).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/user/update?user_id=13", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	updateController(responseRecorder, req)
	if responseRecorder.Code == 204 {
		t.Logf("updateController passed: expected status code %d, and got %d", 204, responseRecorder.Code)
	} else {
		t.Errorf("updateController failed: Expected status code %d, but got %d", 204, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(
		sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Gender,
		user.Birthday,
		userId,
	).Return(nil, fmt.Errorf("no result"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("PUT", "/user/update?user_id=13", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	updateController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("updateControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("updateControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestEnableUserController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	controllerFunc := &userControllerFunc{mockDB, &userservice.UserServiceFactory{}}
	enableUserController := controllerFunc.enableUserController()
	const sql string = `UPDATE mysql_db.user SET enabled=? WHERE user_id=?;`
	const userId uint64 = 13

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql, 1, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/user/enable?user_id=13", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	enableUserController(responseRecorder, req)
	if responseRecorder.Code == 204 {
		t.Logf("enableUserController passed: expected status code %d, and got %d", 204, responseRecorder.Code)
	} else {
		t.Errorf("enableUserController failed: Expected status code %d, but got %d", 204, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql, 1, userId).Return(nil, fmt.Errorf("no result"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("PUT", "/user/enable?user_id=13", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	enableUserController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("enableUserControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("enableUserControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestGetController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	controllerFunc := &userControllerFunc{mockDB, &userservice.UserServiceFactory{}}
	getController := controllerFunc.getController()
	const sql string = `SELECT user_id, user_account, first_name, last_name, email, phone, user_profile, gender, birthday, height, allergies, med_compliance FROM mysql_db.user WHERE user_id=?;`
	var user dto.User
	const userId uint64 = 13

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().QueryRow(sql, userId).Return(mockRow)
	mockRow.EXPECT().Scan(
		&user.UserId,
		&user.UserAccount,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.UserProfile,
		&user.Gender,
		&user.Birthday,
		&user.Height,
		&user.Allergies,
		&user.MedCompliance,
	).Return(nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/get?user_id=13", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getController(responseRecorder, req)
	if responseRecorder.Code == 200 {
		t.Logf("getController passed: expected status code %d, and got %d", 200, responseRecorder.Code)
	} else {
		t.Errorf("getController failed: Expected status code %d, but got %d", 200, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().QueryRow(sql, userId).Return(mockRow)
	mockRow.EXPECT().Scan(
		&user.UserId,
		&user.UserAccount,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.UserProfile,
		&user.Gender,
		&user.Birthday,
		&user.Height,
		&user.Allergies,
		&user.MedCompliance,
	).Return(fmt.Errorf("Scan error"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("GET", "/user/get?user_id=13", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("getControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("getControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestGetUserByAccountController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	controllerFunc := &userControllerFunc{mockDB, &userservice.UserServiceFactory{}}
	getUserByAccountController := controllerFunc.getUserByAccountController()
	const sql string = `SELECT 
	user_id,
	user_account,
	first_name,
	last_name,
	user_profile,
	gender
	FROM mysql_db.user WHERE user_account=?;`
	var user dto.User
	const userAccount string = "userAccount"

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().QueryRow(sql, userAccount).Return(mockRow)
	mockRow.EXPECT().Scan(
		&user.UserId,
		&user.UserAccount,
		&user.FirstName,
		&user.LastName,
		&user.UserProfile,
		&user.Gender,
	).Return(nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/get_by_account?user_account=userAccount", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getUserByAccountController(responseRecorder, req)
	if responseRecorder.Code == 200 {
		t.Logf("getController passed: expected status code %d, and got %d", 200, responseRecorder.Code)
	} else {
		t.Errorf("getController failed: Expected status code %d, but got %d", 200, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().QueryRow(sql, userAccount).Return(mockRow)
	mockRow.EXPECT().Scan(
		&user.UserId,
		&user.UserAccount,
		&user.FirstName,
		&user.LastName,
		&user.UserProfile,
		&user.Gender,
	).Return(fmt.Errorf("Scan error"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("GET", "/user/get_by_account?user_account=userAccount", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getUserByAccountController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("getControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("getControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestGetAllUsersController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	controllerFunc := &userControllerFunc{mockDB, &userservice.UserServiceFactory{}}
	getAllUsersController := controllerFunc.getAllUsersController()
	const sql string = `SELECT * FROM mysql_db.user WHERE user_id!=? ORDER BY created_time DESC;`

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Query(sql, 0).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return([]string{
		"user_id",
		"user_account",
		"user_password",
		"phone",
		"first_name",
		"last_name",
		"birthday",
		"created_time",
		"updated_time",
		"email",
		"allergies",
		"enabled",
		"role",
		"gender",
		"height",
		"user_profile",
		"med_compliance",
		"alert",
	}, nil)
	mockRows.EXPECT().Next().Return(false)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/get_all", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getAllUsersController(responseRecorder, req)
	if responseRecorder.Code == 200 {
		t.Logf("getAllUsersController passed: expected status code %d, and got %d", 200, responseRecorder.Code)
	} else {
		t.Errorf("getAllUsersController failed: Expected status code %d, but got %d", 200, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Query(sql, 0).Return(nil, fmt.Errorf("Query Error"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("GET", "/user/get_all", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getAllUsersController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("getAllUsersControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("getAllUsersControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestGetSpecificRolesController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	controllerFunc := &userControllerFunc{mockDB, &userservice.UserServiceFactory{}}
	getSpecificRolesController := controllerFunc.getSpecificRolesController()
	const sql string = `SELECT * FROM mysql_db.user WHERE role=?;`
	const role uint8 = 1

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Query(sql, role).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return([]string{
		"user_id",
		"user_account",
		"user_password",
		"phone",
		"first_name",
		"last_name",
		"birthday",
		"created_time",
		"updated_time",
		"email",
		"allergies",
		"enabled",
		"role",
		"gender",
		"height",
		"user_profile",
		"med_compliance",
		"alert",
	}, nil)
	mockRows.EXPECT().Next().Return(false)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/get_specific_roles?role=1", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getSpecificRolesController(responseRecorder, req)
	if responseRecorder.Code == 200 {
		t.Logf("getSpecificRolesController passed: expected status code %d, and got %d", 200, responseRecorder.Code)
	} else {
		t.Errorf("getSpecificRolesController failed: Expected status code %d, but got %d", 200, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Query(sql, role).Return(nil, fmt.Errorf("Query Error"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("GET", "/user/get_specific_roles?role=1", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getSpecificRolesController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("getSpecificRolesControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("getSpecificRolesControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestDeleteController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	controllerFunc := &userControllerFunc{mockDB, &userservice.UserServiceFactory{}}
	deleteController := controllerFunc.deleteController()
	const sql string = `DELETE FROM mysql_db.user WHERE user_id=?;`
	const userId uint64 = 7

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/user/delete?user_id=7", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	deleteController(responseRecorder, req)
	if responseRecorder.Code == 204 {
		t.Logf("deleteController passed: expected status code %d, and got %d", 204, responseRecorder.Code)
	} else {
		t.Errorf("deleteController failed: Expected status code %d, but got %d", 204, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql, userId).Return(nil, fmt.Errorf("res err"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("DELETE", "/user/delete?user_id=7", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	deleteController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("deleteControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("deleteControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}
