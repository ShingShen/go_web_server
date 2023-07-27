package logintokenserviceprovider

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"server/dto"
	mock_sqloperator "server/tests/mocks"
)

func TestGetLoginTokenServiceFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	loginTokenServiceFuncFactory := &LoginTokenServiceFuncFactory{}
	const name string = "loginTokenServiceFunc"
	const errName string = "errName"

	// loginTokenServiceFunc
	loginTokenServiceFunc, err := loginTokenServiceFuncFactory.GetLoginTokenServiceFunc(mockDB, name)
	if err == nil {
		t.Logf("loginTokenServiceFunc passed: %v, %v", loginTokenServiceFunc, err)
	} else {
		t.Errorf("loginTokenServiceFunc failed: %v, %v", loginTokenServiceFunc, err)
	}

	// loginTokenServiceFunc Name Err
	loginTokenServiceFuncNameErr, err := loginTokenServiceFuncFactory.GetLoginTokenServiceFunc(mockDB, errName)
	if err != nil {
		t.Logf("loginTokenServiceFuncNameErr passed: %v, %v", loginTokenServiceFuncNameErr, err)
	} else {
		t.Errorf("loginTokenServiceFuncNameErr failed: %v, %v", loginTokenServiceFuncNameErr, err)
	}
}

func TestCreateLoginToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `INSERT INTO mysql_db.login_token(login_token,user_id) value(?,?);`
	const loginToken string = "loginToken"
	const userId uint64 = 1
	loginTokenServiceFunc := &LoginTokenServiceFunc{mockDB}

	// createLoginToken Res Err
	mockDB.EXPECT().Exec(sql, loginToken, userId).Return(nil, fmt.Errorf("res err"))
	createLoginTokenResErr := loginTokenServiceFunc.CreateLoginToken(userId, loginToken)
	if createLoginTokenResErr != nil {
		t.Logf("createLoginTokenResErr passed: %v", createLoginTokenResErr)
	} else {
		t.Errorf("createLoginTokenResErr failed: %v", createLoginTokenResErr)
	}

	// createLoginToken Rows Affected Err
	mockDB.EXPECT().Exec(sql, loginToken, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	createLoginTokenRowsAffectedErr := loginTokenServiceFunc.CreateLoginToken(userId, loginToken)
	if createLoginTokenRowsAffectedErr != nil {
		t.Logf("createLoginTokenRowsAffectedErr passed: %v", createLoginTokenRowsAffectedErr)
	} else {
		t.Errorf("createLoginTokenRowsAffectedErr failed: %v", createLoginTokenRowsAffectedErr)
	}

	// createLoginToken
	mockDB.EXPECT().Exec(sql, loginToken, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	createLoginToken := loginTokenServiceFunc.CreateLoginToken(userId, loginToken)
	if createLoginToken == nil {
		t.Logf("createLoginToken passed: %v", createLoginToken)
	} else {
		t.Errorf("createLoginToken failed: %v", createLoginToken)
	}
}

func TestUpdateLoginToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `UPDATE mysql_db.login_token SET login_token=? WHERE user_id=?;`
	const loginToken string = "loginToken"
	const userId uint64 = 1
	loginTokenServiceFunc := &LoginTokenServiceFunc{mockDB}

	// updateLoginToken Res Err
	mockDB.EXPECT().Exec(sql, loginToken, userId).Return(nil, fmt.Errorf("no result"))
	updateLoginTokenResErr := loginTokenServiceFunc.UpdateLoginToken(userId, loginToken)
	if updateLoginTokenResErr != nil {
		t.Logf("updateLoginTokenResErr passed: %v", updateLoginTokenResErr)
	} else {
		t.Errorf("updateLoginTokenResErr failed: %v", updateLoginTokenResErr)
	}

	// updateLoginToken Rows Affected Err
	mockDB.EXPECT().Exec(sql, loginToken, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	updateLoginTokenRowsAffectedErr := loginTokenServiceFunc.UpdateLoginToken(userId, loginToken)
	if updateLoginTokenRowsAffectedErr != nil {
		t.Logf("updateLoginTokenRowsAffectedErr passed: %v", updateLoginTokenRowsAffectedErr)
	} else {
		t.Errorf("updateLoginTokenRowsAffectedErr failed: %v", updateLoginTokenRowsAffectedErr)
	}

	// updateLoginToken
	mockDB.EXPECT().Exec(sql, loginToken, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	updateLoginToken := loginTokenServiceFunc.UpdateLoginToken(userId, loginToken)
	if updateLoginToken == nil {
		t.Logf("updateLoginToken passed: %v", updateLoginToken)
	} else {
		t.Errorf("updateLoginToken failed: %v", updateLoginToken)
	}
}

func TestGetLoginTokenByUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	sql := `SELECT login_token FROM mysql_db.login_token WHERE user_id=?;`
	const userId uint64 = 1
	var user dto.User
	loginTokenServiceFunc := &LoginTokenServiceFunc{mockDB}

	// getLoginTokenByUserId Scan Err
	mockDB.EXPECT().QueryRow(sql, userId).Return(mockRow)
	mockRow.EXPECT().Scan(&user.LoginToken).Return(fmt.Errorf("Scan error"))
	getLoginTokenByUserIdScanErr, err := loginTokenServiceFunc.GetLoginTokenByUserId(userId)
	if err != nil {
		t.Logf("getLoginTokenByUserIdScanErr passed: %v, %v", getLoginTokenByUserIdScanErr, err)
	} else {
		t.Errorf("getLoginTokenByUserIdScanErr failed: %v, %v", getLoginTokenByUserIdScanErr, err)
	}

	// latestPulse
	mockDB.EXPECT().QueryRow(sql, userId).Return(mockRow)
	mockRow.EXPECT().Scan(&user.LoginToken).Return(nil)
	getLoginTokenByUserId, err := loginTokenServiceFunc.GetLoginTokenByUserId(userId)
	if getLoginTokenByUserId != nil {
		t.Logf("getLoginTokenByUserId passed: %v, %v", getLoginTokenByUserId, err)
	}
	if err != nil {
		t.Errorf("getLoginTokenByUserId failed: %v, %v", getLoginTokenByUserId, err)
	}
}

func TestGetLoginTokenList(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	sql := `SELECT user_id, login_token FROM mysql_db.login_token;`
	loginTokenServiceFunc := &LoginTokenServiceFunc{mockDB}

	// getLoginTokenList Query Err
	mockDB.EXPECT().Query(sql).Return(nil, fmt.Errorf("Query Error"))
	getLoginTokenListQueryErr, err := loginTokenServiceFunc.GetLoginTokenList()
	if err != nil {
		t.Logf("getLoginTokenListQueryErr passed: %v, %v", getLoginTokenListQueryErr, err)
	} else {
		t.Errorf("getLoginTokenListQueryErr failed: %v, %v", getLoginTokenListQueryErr, err)
	}

	// getLoginTokenList Col Err
	mockDB.EXPECT().Query(sql).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return(nil, fmt.Errorf("Columns error"))
	getLoginTokenListColErr, err := loginTokenServiceFunc.GetLoginTokenList()
	if err != nil {
		t.Logf("getLoginTokenListColErr passed: %v, %v", getLoginTokenListColErr, err)
	} else {
		t.Errorf("getLoginTokenListColErr failed: %v, %v", getLoginTokenListColErr, err)
	}

	// getLoginTokenList
	mockDB.EXPECT().Query(sql).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return([]string{
		"login_token_id",
		"login_token",
		"user_id",
		"created_time",
		"updated_time",
	}, nil)
	mockRows.EXPECT().Next().Return(false)
	getLoginTokenList, err := loginTokenServiceFunc.GetLoginTokenList()
	if err == nil {
		t.Logf("getLoginTokenList passed: %v, %v", getLoginTokenList, err)
	} else {
		t.Errorf("getLoginTokenList failed: %v, %v", getLoginTokenList, err)
	}
}
