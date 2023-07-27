package memocontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"server/dto"
	memoService "server/services/dbservice/service/memo"
	mock_sqloperator "server/tests/mocks"
)

func TestGetMemoControllerFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	memoControllerFuncFactory := &memoControllerFuncFactory{}
	const name string = "memoControllerFunc"
	const errName string = "errName"

	// memoControllerFunc
	memoControllerFunc, err := memoControllerFuncFactory.getMemoControllerFunc(mockDB, name)
	if err == nil {
		t.Logf("memoControllerFunc passed: %v, %v", memoControllerFunc, err)
	} else {
		t.Errorf("memoControllerFunc failed: %v, %v", memoControllerFunc, err)
	}

	// memoControllerFuncFunc Name Err
	memoControllerFuncFuncNameErr, err := memoControllerFuncFactory.getMemoControllerFunc(mockDB, errName)
	if err != nil {
		t.Logf("memoControllerFuncFuncNameErr passed: %v, %v", memoControllerFuncFuncNameErr, err)
	} else {
		t.Errorf("memoControllerFuncFuncNameErr failed: %v, %v", memoControllerFuncFuncNameErr, err)
	}
}

func TestCreateController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	controllerFunc := &memoControllerFunc{mockDB, &memoService.MemoServiceFactory{}}
	createController := controllerFunc.createController()
	const sql string = `INSERT INTO mysql_db.memo (content, has_read, user_id) values(?, ?, ?);`
	const content string = "content"
	const userId uint64 = 13
	payload := map[string]interface{}{
		"content": "content",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("Failed to marshal payload:", err)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql, content, 0, userId).Return(mockResult, nil)
	mockResult.EXPECT().LastInsertId().Return(int64(10), nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/memo/create?user_id=13", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	createController(responseRecorder, req)
	if responseRecorder.Code == 204 {
		t.Logf("createController passed: expected status code %d, and got %d", 204, responseRecorder.Code)
	} else {
		t.Errorf("createController failed: expected status code %d, but got %d", 204, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql, content, 0, userId).Return(nil, fmt.Errorf("no result"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("POST", "/memo/create?user_id=13", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	createController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("createControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("createControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestUpdateController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	controllerFunc := &memoControllerFunc{mockDB, &memoService.MemoServiceFactory{}}
	updateController := controllerFunc.updateController()
	const sql string = `UPDATE mysql_db.memo SET content=? WHERE memo_id=?;`
	const memoId uint64 = 12
	const content string = "content"
	payload := map[string]interface{}{
		"content": "content",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("Failed to marshal payload:", err)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql, content, memoId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/memo/update?memo_id=12", bytes.NewReader(body))
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
	mockDB.EXPECT().Exec(sql, content, memoId).Return(nil, fmt.Errorf("no result"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("PUT", "/memo/update?memo_id=12", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	updateController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("updateControllerErr passed: expected status code %d, and got %d", 200, responseRecorderErr.Code)
	} else {
		t.Errorf("updateControllerErr failed: Expected status code %d, but got %d", 200, responseRecorderErr.Code)
	}
}

func TestGetController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	controllerFunc := &memoControllerFunc{mockDB, &memoService.MemoServiceFactory{}}
	getController := controllerFunc.getController()
	const sql string = `SELECT * FROM mysql_db.memo WHERE memo_id=?;`
	var memo dto.Memo
	const memoId uint64 = 13

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().QueryRow(sql, memoId).Return(mockRow)
	mockRow.EXPECT().Scan(
		&memo.MemoId,
		&memo.Content,
		&memo.HasRead,
		&memo.CreatedBy,
		&memo.UpdatedBy,
		&memo.UserId,
		&memo.CreatedTime,
		&memo.UpdatedTime,
	).Return(nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/memo/get?memo_id=13", nil)
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
	mockDB.EXPECT().QueryRow(sql, memoId).Return(mockRow)
	mockRow.EXPECT().Scan(
		&memo.MemoId,
		&memo.Content,
		&memo.HasRead,
		&memo.CreatedBy,
		&memo.UpdatedBy,
		&memo.UserId,
		&memo.CreatedTime,
		&memo.UpdatedTime,
	).Return(fmt.Errorf("Scan error"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("GET", "/memo/get?memo_id=13", nil)
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

func TestGetMemosByUserIdController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	controllerFunc := &memoControllerFunc{mockDB, &memoService.MemoServiceFactory{}}
	getMemosByUserIdController := controllerFunc.getMemosByUserIdController()
	const sql string = `SELECT * FROM mysql_db.memo WHERE user_id=?;`
	const userId uint64 = 13

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Query(sql, userId).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return([]string{
		"memo_id",
		"content",
		"has_read",
		"created_by",
		"updated_by",
		"user_id",
		"created_time",
		"updated_time",
	}, nil)
	mockRows.EXPECT().Next().Return(false)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/memo/get_memos_by_user_id?user_id=13", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getMemosByUserIdController(responseRecorder, req)
	if responseRecorder.Code == 200 {
		t.Logf("getMemosByUserIdController passed: expected status code %d, and got %d", 200, responseRecorder.Code)
	} else {
		t.Errorf("getMemosByUserIdController failed: Expected status code %d, but got %d", 200, responseRecorder.Code)
	}

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Query(sql, userId).Return(nil, fmt.Errorf("Query Error"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("GET", "/memo/get_memos_by_user_id?user_id=13", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	getMemosByUserIdController(responseRecorderErr, reqErr)
	if responseRecorderErr.Code == 400 {
		t.Logf("getMemosByUserIdControllerErr passed: expected status code %d, and got %d", 400, responseRecorderErr.Code)
	} else {
		t.Errorf("getMemosByUserIdControllerErr failed: Expected status code %d, but got %d", 400, responseRecorderErr.Code)
	}
}

func TestDeleteController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	controllerFunc := &memoControllerFunc{mockDB, &memoService.MemoServiceFactory{}}
	deleteController := controllerFunc.deleteController()
	const sql string = `DELETE FROM mysql_db.memo WHERE memo_id=?;`
	const memoId uint64 = 17

	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql, memoId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/memo/delete?memo_id=17", nil)
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
	mockDB.EXPECT().Exec(sql, memoId).Return(nil, fmt.Errorf("res err"))
	responseRecorderErr := httptest.NewRecorder()
	reqErr, err := http.NewRequest("DELETE", "/memo/delete?memo_id=17", nil)
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
