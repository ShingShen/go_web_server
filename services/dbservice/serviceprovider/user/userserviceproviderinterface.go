package userserviceprovider

import (
	"server/dto"

	sqlOperator "server/utils/sqloperator"
)

type IUserServiceFuncFactory interface {
	GetUserServiceFunc(db sqlOperator.ISqlDB, name string) (IUserServiceFunc, error)
}

type IUserServiceFunc interface {
	CreateUser(encodeUserPassword string, user dto.User) error
	UpdateUser(userId uint64, user dto.User) error
	UpdateUserProfileUrl(userId uint64, userProfileUrl string) error
	ResetUserPassword(userId uint64, newUserPassword string) error
	EnableUser(userId uint64) error
	GetUser(userId uint64) ([]byte, error)
	GetUserByAccount(userAccount string) ([]byte, error)
	GetAllUsers() ([]byte, error)
	GetSpecificRoles(role uint8) ([]byte, error)
	DeleteUser(userId uint64) error
	GetUserIdByAccountAndEncodePassword(authAccount string, encodeAuthPassword string) (uint64, error)
}
