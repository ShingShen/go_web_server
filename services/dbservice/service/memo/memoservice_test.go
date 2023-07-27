package memoservice

import (
	"testing"

	"github.com/golang/mock/gomock"

	"server/dto"
	memoserviceprovider "server/services/dbservice/serviceprovider/memo"
	mock_sqloperator "server/tests/mocks"
)

func TestGetMemoService(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	memoServiceFactory := &MemoServiceFactory{}
	const name string = "memo"
	const errName string = "errName"

	memoService, err := memoServiceFactory.GetMemoService(mockDB, name)
	if err == nil {
		t.Logf("memoService passed: %v, %v", memoService, err)
	} else {
		t.Errorf("memoService failed: %v, %v", memoService, err)
	}

	memoServiceNameErr, err := memoServiceFactory.GetMemoService(mockDB, errName)
	if err != nil {
		t.Logf("memoServiceNameErr passed: %v, %v", memoServiceNameErr, err)
	} else {
		t.Errorf("memoServiceNameErr failed: %v, %v", memoServiceNameErr, err)
	}
}

func TestCreateMemo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	const sql string = `INSERT INTO mysql_db.memo (content, has_read, user_id) values(?, ?, ?);`
	const content string = "content"
	const userId uint64 = 1
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	mockDB.EXPECT().Exec(sql, content, 0, userId).Return(mockResult, nil)
	mockResult.EXPECT().LastInsertId().Return(int64(10), nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)

	memoService := &MemoService{mockDB, &memoserviceprovider.MemoServiceFuncFactory{}}
	createMemo := memoService.CreateMemo(content, userId)
	if createMemo == nil {
		t.Logf("createMemo passed: %v", createMemo)
	} else {
		t.Errorf("createMemo failed: %v", createMemo)
	}
}

func TestUpdateMemo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	const sql string = `UPDATE mysql_db.memo SET content=? WHERE memo_id=?;`
	const memoId uint64 = 1
	const content string = "content"
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	mockDB.EXPECT().Exec(sql, content, memoId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	memoService := &MemoService{mockDB, &memoserviceprovider.MemoServiceFuncFactory{}}
	updateMemo := memoService.UpdateMemo(memoId, content)
	if updateMemo == nil {
		t.Logf("updateMemo passed: %v", updateMemo)
	} else {
		t.Errorf("updateMemo failed: %v", updateMemo)
	}
}

func TestGetMemo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	sql := `SELECT * FROM mysql_db.memo WHERE memo_id=?;`
	var memo dto.Memo
	const memoId uint64 = 1
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
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
	memoService := &MemoService{mockDB, &memoserviceprovider.MemoServiceFuncFactory{}}
	getMemo, err := memoService.GetMemo(memoId)
	if getMemo != nil {
		t.Logf("getMemo passed: %v, %v", getMemo, err)
	}
	if err != nil {
		t.Errorf("getMemo failed: %v, %v", getMemo, err)
	}
}

func TestGetMemosByUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	sql := `SELECT * FROM mysql_db.memo WHERE user_id=?;`
	const userId uint64 = 1

	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
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
	memoService := &MemoService{mockDB, &memoserviceprovider.MemoServiceFuncFactory{}}
	getMemosByUserId, err := memoService.GetMemosByUserId(userId)
	if getMemosByUserId != nil {
		t.Logf("getMemosByUserId passed: %v, %v", getMemosByUserId, err)
	}
	if err != nil {
		t.Errorf("getMemosByUserId failed: %v, %v", getMemosByUserId, err)
	}
}

func TestDeleteMemo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	sql := `DELETE FROM mysql_db.memo WHERE memo_id=?;`
	const pulseId uint64 = 1

	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	mockDB.EXPECT().Exec(sql, pulseId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	memoService := &MemoService{mockDB, &memoserviceprovider.MemoServiceFuncFactory{}}
	deleteMemo := memoService.DeleteMemo(pulseId)
	if deleteMemo == nil {
		t.Logf("deleteMemo passed: %v", deleteMemo)
	} else {
		t.Errorf("deleteMemo failed: %v", deleteMemo)
	}
}
