package systemservice

import (
	"encoding/json"
	"fmt"
	"server/utils/encryption"
	"strconv"

	"github.com/go-redis/redis"

	loginTokenServiceProvider "server/services/dbservice/serviceprovider/logintoken"
	userServiceProvider "server/services/dbservice/serviceprovider/user"
	sqlOperator "server/utils/sqloperator"
)

type SystemServiceFactory struct{}

func (s *SystemServiceFactory) GetSystemService(db sqlOperator.ISqlDB, name string) (ISystemService, error) {
	if name == "system" {
		return &SystemService{
			db:                           db,
			userServiceFuncFactory:       &userServiceProvider.UserServiceFuncFactory{},
			loginTokenServiceFuncFactory: &loginTokenServiceProvider.LoginTokenServiceFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong system type passed")
}

type SystemService struct {
	db                           sqlOperator.ISqlDB
	userServiceFuncFactory       userServiceProvider.IUserServiceFuncFactory
	loginTokenServiceFuncFactory loginTokenServiceProvider.ILoginTokenServiceFuncFactory
}

func (s *SystemService) RefreshLoginTokens(rdb *redis.Client) error {
	tx, _ := s.db.Begin()
	defer tx.Rollback()

	userServiceFunc, _ := s.userServiceFuncFactory.GetUserServiceFunc(s.db, "userServiceFunc")
	loginTokenServiceFunc, _ := s.loginTokenServiceFuncFactory.GetLoginTokenServiceFunc(s.db, "loginTokenServiceFunc")

	userList, _ := userServiceFunc.GetAllUsers()
	var userlistJsonData []map[string]interface{}
	json.Unmarshal(userList, &userlistJsonData)
	for i := 0; i < len(userlistJsonData); i++ {
		userId, _ := strconv.ParseUint(userlistJsonData[i]["user_id"].(string), 10, 64)
		userAccount := userlistJsonData[i]["user_account"].(string)
		userPassword := userlistJsonData[i]["user_password"].(string)
		loginToken := encryption.EncryptingLoginToken(userAccount, userPassword)
		updateLoginToken := loginTokenServiceFunc.UpdateLoginToken(userId, loginToken)
		if updateLoginToken != nil {
			return updateLoginToken
		}
		fmt.Printf("Set user id %d, login token %s in MySQL DB\n", userId, loginToken)
	}

	loginTokenList, _ := loginTokenServiceFunc.GetLoginTokenList()
	var loginTokenListJsonData []map[string]interface{}
	json.Unmarshal(loginTokenList, &loginTokenListJsonData)
	setLoginTokenCacheList := loginTokenServiceFunc.SetLoginTokenCacheList(rdb, loginTokenList)
	if setLoginTokenCacheList != nil {
		return setLoginTokenCacheList
	}

	return nil
}
