package systemservice

import (
	"github.com/go-redis/redis"

	sqlOperator "server/utils/sqloperator"
)

type ISystemServiceFactory interface {
	GetSystemService(db sqlOperator.ISqlDB, name string) (ISystemService, error)
}

type ISystemService interface {
	RefreshLoginTokens(rdb *redis.Client) error
}
