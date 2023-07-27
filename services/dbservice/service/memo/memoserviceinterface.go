package memoservice

import sqlOperator "server/utils/sqloperator"

type IMemoServiceFactory interface {
	GetMemoService(db sqlOperator.ISqlDB, name string) (IMemoService, error)
}

type IMemoService interface {
	CreateMemo(content string, userId uint64) error
	UpdateMemo(memoId uint64, content string) error
	GetMemo(memoId uint64) ([]byte, error)
	GetMemosByUserId(userId uint64) ([]byte, error)
	DeleteMemo(memoId uint64) error
}
