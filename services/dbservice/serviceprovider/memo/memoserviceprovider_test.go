package memoserviceprovider

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"server/dto"
	mock_sqloperator "server/tests/mocks"
)

func TestGetSMemoServiceFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	memoServiceFuncFactory := &MemoServiceFuncFactory{}
	const name string = "memoServiceFunc"
	const errName string = "errName"

	memoServiceFunc, err := memoServiceFuncFactory.GetMemoServiceFunc(mockDB, name)
	if err == nil {
		t.Logf("memoServiceFunc passed: %v, %v", memoServiceFunc, err)
	} else {
		t.Errorf("memoServiceFunc failed: %v, %v", memoServiceFunc, err)
	}

	memoServiceFuncNameErr, err := memoServiceFuncFactory.GetMemoServiceFunc(mockDB, errName)
	if err != nil {
		t.Logf("memoServiceFuncNameErr passed: %v, %v", memoServiceFuncNameErr, err)
	} else {
		t.Errorf("memoServiceFuncNameErr failed: %v, %v", memoServiceFuncNameErr, err)
	}
}

func TestCreateMemo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `INSERT INTO mysql_db.memo (content, has_read, user_id) values(?, ?, ?);`
	const content string = "content"
	const userId uint64 = 1
	memoServiceFunc := &MemoServiceFunc{mockDB}

	// createMemo Res Err
	mockDB.EXPECT().Exec(sql, content, 0, userId).Return(nil, fmt.Errorf("no result"))
	createMemoResErr := memoServiceFunc.CreateMemo(content, userId)
	if createMemoResErr != nil {
		t.Logf("createMemoResErr passed: %v", createMemoResErr)
	} else {
		t.Errorf("createMemoResErr failed: %v", createMemoResErr)
	}

	// createMemo Insert ID Err
	mockDB.EXPECT().Exec(sql, content, 0, userId).Return(mockResult, nil)
	mockResult.EXPECT().LastInsertId().Return(int64(0), fmt.Errorf("Failed to insert ID"))
	createMemoInsertIDErr := memoServiceFunc.CreateMemo(content, userId)
	if createMemoInsertIDErr != nil {
		t.Logf("createMemoInsertIDErr passed: %v", createMemoInsertIDErr)
	} else {
		t.Errorf("createMemoInsertIDErr failed: %v", createMemoInsertIDErr)
	}

	// createMemo Rows Affected Err
	mockDB.EXPECT().Exec(sql, content, 0, userId).Return(mockResult, nil)
	mockResult.EXPECT().LastInsertId().Return(int64(10), nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	createMemoRowsAffectedErr := memoServiceFunc.CreateMemo(content, userId)
	if createMemoRowsAffectedErr != nil {
		t.Logf("createMemoRowsAffectedErr passed: %v", createMemoRowsAffectedErr)
	} else {
		t.Errorf("createMemoRowsAffectedErr failed: %v", createMemoRowsAffectedErr)
	}

	// createMemo
	mockDB.EXPECT().Exec(sql, content, 0, userId).Return(mockResult, nil)
	mockResult.EXPECT().LastInsertId().Return(int64(10), nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	createMemo := memoServiceFunc.CreateMemo(content, userId)
	if createMemo == nil {
		t.Logf("createMemo passed: %v", createMemo)
	} else {
		t.Errorf("createMemo failed: %v", createMemo)
	}
}

func TestUpdateMemo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `UPDATE mysql_db.memo SET content=? WHERE memo_id=?;`
	const memoId uint64 = 1
	const content string = "content"
	memoServiceFunc := &MemoServiceFunc{mockDB}

	// updateMemo Res Err
	mockDB.EXPECT().Exec(sql, content, memoId).Return(nil, fmt.Errorf("no result"))
	updateMemoResErr := memoServiceFunc.UpdateMemo(memoId, content)
	if updateMemoResErr != nil {
		t.Logf("updateMemoResErr passed: %v", updateMemoResErr)
	} else {
		t.Errorf("updateMemoResErr failed: %v", updateMemoResErr)
	}

	// updateMemo Rows Affected Err
	mockDB.EXPECT().Exec(sql, content, memoId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	updateMemoRowsAffectedErr := memoServiceFunc.UpdateMemo(memoId, content)
	if updateMemoRowsAffectedErr != nil {
		t.Logf("updateMemoRowsAffectedErr passed: %v", updateMemoRowsAffectedErr)
	} else {
		t.Errorf("updateMemoRowsAffectedErr failed: %v", updateMemoRowsAffectedErr)
	}

	// updateMemo
	mockDB.EXPECT().Exec(sql, content, memoId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	updateMemo := memoServiceFunc.UpdateMemo(memoId, content)
	if updateMemo == nil {
		t.Logf("updateMemo passed: %v", updateMemo)
	} else {
		t.Errorf("updateMemo failed: %v", updateMemo)
	}
}

func TestGetMemo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	sql := `SELECT * FROM mysql_db.memo WHERE memo_id=?;`
	var memo dto.Memo
	const memoId uint64 = 1
	memoServiceFunc := &MemoServiceFunc{mockDB}

	// getMemo Scan Err
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
	getMemoScanErr, err := memoServiceFunc.GetMemo(memoId)
	if err != nil {
		t.Logf("getMemoScanErr passed: %v, %v", getMemoScanErr, err)
	} else {
		t.Errorf("getMemoScanErr failed: %v, %v", getMemoScanErr, err)
	}

	// getMemo
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
	getMemo, err := memoServiceFunc.GetMemo(memoId)
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
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	sql := `SELECT * FROM mysql_db.memo WHERE user_id=?;`
	const userId uint64 = 1
	memoServiceFunc := &MemoServiceFunc{mockDB}

	// getMemosByUserId Query Err
	mockDB.EXPECT().Query(sql, userId).Return(nil, fmt.Errorf("Query Error"))
	getMemosByUserIdtQueryErr, err := memoServiceFunc.GetMemosByUserId(userId)
	if err != nil {
		t.Logf("getMemosByUserIdtQueryErr passed: %v, %v", getMemosByUserIdtQueryErr, err)
	} else {
		t.Errorf("getMemosByUserIdtQueryErr failed: %v, %v", getMemosByUserIdtQueryErr, err)
	}

	// getMemosByUserId Col Err
	mockDB.EXPECT().Query(sql, userId).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return(nil, fmt.Errorf("Columns error"))
	getMemosByUserIdColErr, err := memoServiceFunc.GetMemosByUserId(userId)
	if err != nil {
		t.Logf("getMemosByUserIdColErr passed: %v, %v", getMemosByUserIdColErr, err)
	} else {
		t.Errorf("getMemosByUserIdColErr failed: %v, %v", getMemosByUserIdColErr, err)
	}

	// getMemosByUserId
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
	getMemosByUserId, err := memoServiceFunc.GetMemosByUserId(userId)
	if err == nil {
		t.Logf("getMemosByUserId passed: %v, %v", getMemosByUserId, err)
	} else {
		t.Errorf("getMemosByUserId failed: %v, %v", getMemosByUserId, err)
	}
}

func TestDeleteMemo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	sql := `DELETE FROM mysql_db.memo WHERE memo_id=?;`
	const memoId uint64 = 1
	memoServiceFunc := &MemoServiceFunc{mockDB}

	// deleteMemo res err
	mockDB.EXPECT().Exec(sql, memoId).Return(nil, fmt.Errorf("res err"))
	deleteMemoResErr := memoServiceFunc.DeleteMemo(memoId)
	if deleteMemoResErr != nil {
		t.Logf("deleteMemoResErr passed: %v", deleteMemoResErr)
	} else {
		t.Errorf("deleteMemoResErr failed: %v", deleteMemoResErr)
	}

	// deleteMemo rows affected err
	mockDB.EXPECT().Exec(sql, memoId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("rows affected error"))
	deleteMemoRowsAffectedErr := memoServiceFunc.DeleteMemo(memoId)
	if deleteMemoRowsAffectedErr != nil {
		t.Logf("deleteMemoRowsAffectedErr passed: %v", deleteMemoRowsAffectedErr)
	} else {
		t.Errorf("deleteMemoRowsAffectedErr failed: %v", deleteMemoRowsAffectedErr)
	}

	// deleteMemo
	mockDB.EXPECT().Exec(sql, memoId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	deleteMemo := memoServiceFunc.DeleteMemo(memoId)
	if deleteMemo == nil {
		t.Logf("deleteMemo passed: %v", deleteMemo)
	} else {
		t.Errorf("deleteMemo failed: %v", deleteMemo)
	}
}
