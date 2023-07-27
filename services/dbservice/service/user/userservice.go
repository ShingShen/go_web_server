package userservice

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"server/dto"
	"server/utils/encryption"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"

	loginTokenServiceProvider "server/services/dbservice/serviceprovider/logintoken"
	userServiceProvider "server/services/dbservice/serviceprovider/user"
	sqlOperator "server/utils/sqloperator"
	storageFileOperator "server/utils/storagefileoperator"
)

type UserServiceFactory struct{}

func (u *UserServiceFactory) GetUserService(db sqlOperator.ISqlDB, name string) (IUserService, error) {
	if name == "user" {
		return &UserService{
			db:                           db,
			userServiceFuncFactory:       &userServiceProvider.UserServiceFuncFactory{},
			loginTokenServiceFuncFactory: &loginTokenServiceProvider.LoginTokenServiceFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong user type passed")
}

type UserService struct {
	db                           sqlOperator.ISqlDB
	userServiceFuncFactory       userServiceProvider.IUserServiceFuncFactory
	loginTokenServiceFuncFactory loginTokenServiceProvider.ILoginTokenServiceFuncFactory
}

func (u *UserService) CreateUser(rdb *redis.Client, user dto.User) error {
	encodeUserPassword := encryption.EncryptingUserPassword(user.UserPassword)
	loginToken := encryption.EncryptingLoginToken(user.UserAccount, user.UserPassword)

	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	loginTokenServiceFunc, _ := u.loginTokenServiceFuncFactory.GetLoginTokenServiceFunc(u.db, "loginTokenServiceFunc")

	createUserFunc := userServiceFunc.CreateUser(encodeUserPassword, user)

	userId, _ := userServiceFunc.GetUserIdByAccountAndEncodePassword(
		user.UserAccount,
		encodeUserPassword,
	)

	// MySQL
	createLoginToken := loginTokenServiceFunc.CreateLoginToken(userId, loginToken)
	if createLoginToken != nil {
		return createLoginToken
	}

	// Cache
	createLoginTokenCache := loginTokenServiceFunc.SetLoginTokenCache(rdb, userId, loginToken)
	if createLoginTokenCache != nil {
		return createLoginTokenCache
	}

	return createUserFunc
}

func (u *UserService) CreateLoginToken(rdb *redis.Client, user dto.User) error {
	encodeUserPassword := encryption.EncryptingUserPassword(user.UserPassword)
	loginToken := encryption.EncryptingLoginToken(user.UserAccount, user.UserPassword)

	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	loginTokenServiceFunc, _ := u.loginTokenServiceFuncFactory.GetLoginTokenServiceFunc(u.db, "loginTokenServiceFunc")

	userId, _ := userServiceFunc.GetUserIdByAccountAndEncodePassword(
		user.UserAccount,
		encodeUserPassword,
	)

	// MySQL
	createLoginToken := loginTokenServiceFunc.CreateLoginToken(userId, loginToken)
	if createLoginToken != nil {
		return createLoginToken
	}

	// Cache
	createLoginTokenCache := loginTokenServiceFunc.SetLoginTokenCache(rdb, userId, loginToken)
	if createLoginTokenCache != nil {
		return createLoginTokenCache
	}

	return createLoginToken
}

func (u *UserService) UpdateLoginToken(rdb *redis.Client, user dto.User) error {
	encodeUserPassword := encryption.EncryptingUserPassword(user.UserPassword)
	loginToken := encryption.EncryptingLoginToken(user.UserAccount, user.UserPassword)

	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	loginTokenServiceFunc, _ := u.loginTokenServiceFuncFactory.GetLoginTokenServiceFunc(u.db, "loginTokenServiceFunc")

	userId, _ := userServiceFunc.GetUserIdByAccountAndEncodePassword(
		user.UserAccount,
		encodeUserPassword,
	)

	// MySQL
	updateLoginToken := loginTokenServiceFunc.UpdateLoginToken(userId, loginToken)
	if updateLoginToken != nil {
		return updateLoginToken
	}

	// Cache
	createLoginTokenCache := loginTokenServiceFunc.SetLoginTokenCache(rdb, userId, loginToken)
	if createLoginTokenCache != nil {
		return createLoginTokenCache
	}

	return updateLoginToken
}

func (u *UserService) UpdateUser(userId uint64, user dto.User) error {
	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	return userServiceFunc.UpdateUser(userId, user)
}

func (u *UserService) UploadUserProfile(file multipart.File, userId uint64, profileName string, cloudStorage *storage.Client) error {
	bucket := "bucket"
	objectName := "users/" + strconv.FormatUint(userId, 10) + "/profile/" + profileName
	uploadUserProfile := storageFileOperator.StorageFileUpload(file, bucket, objectName, cloudStorage)
	if uploadUserProfile != nil {
		return uploadUserProfile
	}

	userProfileUrl, err := storageFileOperator.StorageFileUrlGet(bucket, objectName, cloudStorage)
	if err != nil {
		return err
	}

	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")

	updateUser := userServiceFunc.UpdateUserProfileUrl(userId, userProfileUrl)
	if updateUser != nil {
		return updateUser
	}

	return nil
}

func (u *UserService) ResetUserPassword(userId uint64, newUserPassword string) error {
	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")

	encodeUserPassword := encryption.EncryptingUserPassword(newUserPassword)
	return userServiceFunc.ResetUserPassword(userId, encodeUserPassword)
}

func (u *UserService) EnableUser(userId uint64) error {
	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	return userServiceFunc.EnableUser(userId)
}

func (u *UserService) GetUser(userId uint64) ([]byte, error) {
	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	return userServiceFunc.GetUser(userId)
}

func (u *UserService) GetUserByAccount(userAccount string) ([]byte, error) {
	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	return userServiceFunc.GetUserByAccount(userAccount)
}

func (u *UserService) GetLoginTokenByUserId(rdb *redis.Client, userId uint64) (*string, error) {
	tx, _ := u.db.Begin()
	defer tx.Rollback()

	loginTokenServiceFunc, _ := u.loginTokenServiceFuncFactory.GetLoginTokenServiceFunc(u.db, "loginTokenServiceFunc")

	loginTokenCache, err := loginTokenServiceFunc.GetLoginTokenCacheByUserId(rdb, userId)
	if err == nil {
		return loginTokenCache, nil
	}

	fmt.Println("Reading login token in MySQL...")
	loginToken, err := loginTokenServiceFunc.GetLoginTokenByUserId(userId)
	if err != nil {
		return nil, err
	}
	return loginToken, nil
}

func (u *UserService) GetAllUsers() ([]byte, error) {
	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	return userServiceFunc.GetAllUsers()
}

func (u *UserService) GetSpecificRoles(role uint8) ([]byte, error) {
	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	return userServiceFunc.GetSpecificRoles(role)
}

func (u *UserService) DeleteUser(userId uint64) error {
	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	return userServiceFunc.DeleteUser(userId)
}

func (u *UserService) Login(authorization string) ([]byte, error) {
	tx, _ := u.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := u.userServiceFuncFactory.GetUserServiceFunc(u.db, "userServiceFunc")
	loginTokenServiceFunc, _ := u.loginTokenServiceFuncFactory.GetLoginTokenServiceFunc(u.db, "loginTokenServiceFunc")

	loginAuthParser, _ := encryption.LoginAuthParser(authorization)

	userId, _ := userServiceFunc.GetUserIdByAccountAndEncodePassword(
		loginAuthParser[0]["auth_account"].(string),
		loginAuthParser[0]["encode_auth_password"].(string),
	)
	loginToken, _ := loginTokenServiceFunc.GetLoginTokenByUserId(userId)

	loginJsonData, _ := json.Marshal(map[string]interface{}{
		"login_token": loginToken,
		"user_id":     userId,
	})

	return loginJsonData, nil
}
