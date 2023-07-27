package memoserviceprovider

import (
	"encoding/json"
	"fmt"
	"server/dto"
	"server/utils/operator"

	sqlOperator "server/utils/sqloperator"
)

type MemoServiceFuncFactory struct{}

func (m *MemoServiceFuncFactory) GetMemoServiceFunc(db sqlOperator.ISqlDB, name string) (IMemoServiceFunc, error) {
	if name == "memoServiceFunc" {
		return &MemoServiceFunc{db}, nil
	}
	return nil, fmt.Errorf("wrong memo service func type passed")
}

type MemoServiceFunc struct {
	db sqlOperator.ISqlDB
}

func (m *MemoServiceFunc) CreateMemo(content string, userId uint64) error {
	sql := `INSERT INTO mysql_db.memo (content, has_read, user_id) values(?, ?, ?);`
	res, err := m.db.Exec(sql, content, 0, userId)
	if err != nil {
		fmt.Printf("Failed to insert memo data, err: %v", err)
		return err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Failed to get inserted memo id, err: %v", err)
		return err
	}
	fmt.Println("Inserted memo id: ", lastInsertId)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v", err)
		return err
	}

	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}

func (m *MemoServiceFunc) UpdateMemo(memoId uint64, content string) error {
	sql := `UPDATE mysql_db.memo SET content=? WHERE memo_id=?;`
	res, err := m.db.Exec(sql, content, memoId)
	if err != nil {
		fmt.Printf("Failed to update memo data, err: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v", err)
		return err
	}
	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}

func (m *MemoServiceFunc) GetMemo(memoId uint64) ([]byte, error) {
	var memo dto.Memo
	sql := `SELECT * FROM mysql_db.memo WHERE memo_id=?;`
	res := m.db.QueryRow(sql, memoId)
	err := res.Scan(
		&memo.MemoId,
		&memo.Content,
		&memo.HasRead,
		&memo.CreatedBy,
		&memo.UpdatedBy,
		&memo.UserId,
		&memo.CreatedTime,
		&memo.UpdatedTime,
	)
	if err != nil {
		return nil, err
	} else {
		memoJsonData, _ := json.Marshal(&memo)
		return memoJsonData, nil
	}
}

func (m *MemoServiceFunc) GetMemosByUserId(userId uint64) ([]byte, error) {
	sql := `SELECT * FROM mysql_db.memo WHERE user_id=?;`
	res, err := m.db.Query(sql, userId)
	if err != nil {
		fmt.Printf("Failed to query memos, err: %v\n", err)
		return nil, err
	}

	memoList, err := operator.CreatingDataList(res)
	if err != nil {
		return nil, err
	} else {
		return memoList, nil
	}
}

func (m *MemoServiceFunc) DeleteMemo(memoId uint64) error {
	sql := `DELETE FROM mysql_db.memo WHERE memo_id=?;`
	res, err := m.db.Exec(sql, memoId)
	if err != nil {
		fmt.Printf("Failed to delete memo data, err: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v", err)
		return err
	}
	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}
