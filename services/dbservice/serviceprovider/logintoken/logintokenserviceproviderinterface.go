package logintokenserviceprovider

import (
	sqlOperator "server/utils/sqloperator"

	"github.com/go-redis/redis"
)

type ILoginTokenServiceFuncFactory interface {
	GetLoginTokenServiceFunc(db sqlOperator.ISqlDB, name string) (ILoginTokenServiceFunc, error)
}

type ILoginTokenServiceFunc interface {
	CreateLoginToken(userId uint64, loginToken string) error
	UpdateLoginToken(userId uint64, loginToken string) error
	GetLoginTokenByUserId(userId uint64) (*string, error)
	GetLoginTokenList() ([]byte, error)
	SetLoginTokenCache(rdb *redis.Client, userId uint64, loginToken string) error
	SetLoginTokenCacheList(rdb *redis.Client, loginTokenList []byte) error
	GetLoginTokenCacheByUserId(rdb *redis.Client, userId uint64) (*string, error)
}
