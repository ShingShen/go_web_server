package userservice

import (
	"mime/multipart"
	"server/dto"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"

	sqlOperator "server/utils/sqloperator"
)

type IUserServiceFactory interface {
	GetUserService(db sqlOperator.ISqlDB, name string) (IUserService, error)
}

type IUserService interface {
	CreateUser(rdb *redis.Client, user dto.User) error
	CreateLoginToken(rdb *redis.Client, user dto.User) error
	UpdateLoginToken(rdb *redis.Client, user dto.User) error
	UpdateUser(userId uint64, user dto.User) error
	UploadUserProfile(file multipart.File, userId uint64, profileName string, cloudStorage *storage.Client) error
	ResetUserPassword(userId uint64, newUserPassword string) error
	EnableUser(userId uint64) error
	GetUser(userId uint64) ([]byte, error)
	GetUserByAccount(userAccount string) ([]byte, error)
	GetLoginTokenByUserId(rdb *redis.Client, userId uint64) (*string, error)
	GetAllUsers() ([]byte, error)
	GetSpecificRoles(role uint8) ([]byte, error)
	DeleteUser(userId uint64) error
	Login(authorization string) ([]byte, error)
}
