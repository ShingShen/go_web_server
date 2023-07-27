package memoserviceprovider

import sqlOperator "server/utils/sqloperator"

type IMemoServiceFuncFactory interface {
	GetMemoServiceFunc(db sqlOperator.ISqlDB, name string) (IMemoServiceFunc, error)
}

type IMemoServiceFunc interface {
	CreateMemo(content string, userId uint64) error
	UpdateMemo(memoId uint64, content string) error
	GetMemo(memoId uint64) ([]byte, error)
	GetMemosByUserId(userId uint64) ([]byte, error)
	DeleteMemo(memoId uint64) error
}
