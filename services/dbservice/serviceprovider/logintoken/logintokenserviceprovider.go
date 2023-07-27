package logintokenserviceprovider

import (
	"encoding/json"
	"fmt"
	"server/dto"
	"server/utils/operator"
	"strconv"
	"time"

	"github.com/go-redis/redis"

	sqlOperator "server/utils/sqloperator"
)

type LoginTokenServiceFuncFactory struct{}

func (l *LoginTokenServiceFuncFactory) GetLoginTokenServiceFunc(db sqlOperator.ISqlDB, name string) (ILoginTokenServiceFunc, error) {
	if name == "loginTokenServiceFunc" {
		return &LoginTokenServiceFunc{db}, nil
	}
	return nil, fmt.Errorf("wrong login token service func type passed")
}

type LoginTokenServiceFunc struct {
	db sqlOperator.ISqlDB
}

func (l *LoginTokenServiceFunc) CreateLoginToken(userId uint64, loginToken string) error {
	sql := `INSERT INTO mysql_db.login_token(login_token,user_id) value(?,?);`
	res, err := l.db.Exec(sql, loginToken, userId)
	if err != nil {
		fmt.Printf("Failed to create login token, err: %v\n", err)
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v\n", err)
		return err
	}
	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}

func (l *LoginTokenServiceFunc) UpdateLoginToken(userId uint64, loginToken string) error {
	sql := `UPDATE mysql_db.login_token SET login_token=? WHERE user_id=?;`
	res, err := l.db.Exec(sql, loginToken, userId)
	if err != nil {
		fmt.Printf("Failed to update login token, err: %v", err)
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

func (l *LoginTokenServiceFunc) GetLoginTokenByUserId(userId uint64) (*string, error) {
	var user dto.User
	sql := `SELECT login_token FROM mysql_db.login_token WHERE user_id=?;`
	res := l.db.QueryRow(sql, userId)
	err := res.Scan(&user.LoginToken)
	if err != nil {
		return nil, err
	} else {
		return &user.LoginToken, nil
	}
}

func (l *LoginTokenServiceFunc) GetLoginTokenList() ([]byte, error) {
	sql := `SELECT user_id, login_token FROM mysql_db.login_token;`
	res, err := l.db.Query(sql)
	if err != nil {
		fmt.Printf("Failed to query login tokens, err: %v\n", err)
		return nil, err
	}

	loginTokenList, err := operator.CreatingDataList(res)
	if err != nil {
		return nil, err
	} else {
		return loginTokenList, nil
	}
}

func (l *LoginTokenServiceFunc) SetLoginTokenCache(rdb *redis.Client, userId uint64, loginToken string) error {
	var config dto.Config
	configData, _ := operator.LoadJson(config, "config/env/config.json")
	loginTokenCacheDB := int((configData.(map[string]interface{})["redis"]).(map[string]interface{})["login_token"].(float64))
	rdb.Options().DB = loginTokenCacheDB

	set, err := rdb.Set(strconv.FormatUint(userId, 10), loginToken, 2592000*time.Second).Result()
	if err != nil {
		fmt.Printf("Failed to set login token cache, err: %v", err)
		return err
	}
	fmt.Println("Set login token: ", set)
	return nil
}

func (l *LoginTokenServiceFunc) SetLoginTokenCacheList(rdb *redis.Client, loginTokenList []byte) error {
	var config dto.Config
	configData, _ := operator.LoadJson(config, "config/env/config.json")
	loginTokenCacheDB := int((configData.(map[string]interface{})["redis"]).(map[string]interface{})["login_token"].(float64))
	rdb.Options().DB = loginTokenCacheDB

	var list []map[string]interface{}
	json.Unmarshal(loginTokenList, &list)

	rdb.FlushDB()

	for i := 0; i < len(list); i++ {
		userId := list[i]["user_id"].(string)
		loginToken := list[i]["login_token"].(string)
		_, err := rdb.Set(userId, loginToken, 2592000*time.Second).Result()
		if err != nil {
			fmt.Printf("Failed to set login token cache in user id %s, err: %v", userId, err)
			return err
		}
		fmt.Printf("Set user id %s, login token cache %s\n", userId, loginToken)
	}

	return nil
}

func (l *LoginTokenServiceFunc) GetLoginTokenCacheByUserId(rdb *redis.Client, userId uint64) (*string, error) {
	var config dto.Config
	configData, _ := operator.LoadJson(config, "config/env/config.json")
	loginTokenCacheDB := int((configData.(map[string]interface{})["redis"]).(map[string]interface{})["login_token"].(float64))
	rdb.Options().DB = loginTokenCacheDB

	val, err := rdb.Get(strconv.FormatUint(userId, 10)).Result()
	if err == redis.Nil {
		fmt.Println("Failed to get login token cache, err: ", err)
		return nil, err
	} else if err != nil {
		fmt.Println("Failed to get login token cache, err: ", err)
		return nil, err
	}
	fmt.Println("Get login token: ", val)
	return &val, nil
}
