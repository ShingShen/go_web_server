package memoservice

import (
	"fmt"

	memoServiceProvider "server/services/dbservice/serviceprovider/memo"
	sqlOperator "server/utils/sqloperator"
)

type MemoServiceFactory struct{}

func (m *MemoServiceFactory) GetMemoService(db sqlOperator.ISqlDB, name string) (IMemoService, error) {
	if name == "memo" {
		return &MemoService{
			db:                     db,
			memoServiceFuncFactory: &memoServiceProvider.MemoServiceFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong memo type passed")
}

type MemoService struct {
	db                     sqlOperator.ISqlDB
	memoServiceFuncFactory memoServiceProvider.IMemoServiceFuncFactory
}

func (m *MemoService) CreateMemo(content string, userId uint64) error {
	tx, _ := m.db.Begin()
	defer tx.Rollback()

	memoServiceFunc, _ := m.memoServiceFuncFactory.GetMemoServiceFunc(m.db, "memoServiceFunc")
	return memoServiceFunc.CreateMemo(content, userId)
}

func (m *MemoService) UpdateMemo(memoId uint64, content string) error {
	tx, _ := m.db.Begin()
	defer tx.Rollback()

	memoServiceFunc, _ := m.memoServiceFuncFactory.GetMemoServiceFunc(m.db, "memoServiceFunc")
	return memoServiceFunc.UpdateMemo(memoId, content)
}

func (m *MemoService) GetMemo(memoId uint64) ([]byte, error) {
	tx, _ := m.db.Begin()
	defer tx.Rollback()

	memoServiceFunc, _ := m.memoServiceFuncFactory.GetMemoServiceFunc(m.db, "memoServiceFunc")
	return memoServiceFunc.GetMemo(memoId)
}

func (m *MemoService) GetMemosByUserId(userId uint64) ([]byte, error) {
	tx, _ := m.db.Begin()
	defer tx.Rollback()

	memoServiceFunc, _ := m.memoServiceFuncFactory.GetMemoServiceFunc(m.db, "memoServiceFunc")
	return memoServiceFunc.GetMemosByUserId(userId)
}

func (m *MemoService) DeleteMemo(memoId uint64) error {
	tx, _ := m.db.Begin()
	defer tx.Rollback()

	memoServiceFunc, _ := m.memoServiceFuncFactory.GetMemoServiceFunc(m.db, "memoServiceFunc")
	return memoServiceFunc.DeleteMemo(memoId)
}
