package encryption

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"server/dto"
	"server/utils/operator"
	"strings"
)

func Base64Sha256(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	encodeAuthPassword := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return encodeAuthPassword
}

func HashPassword(password, salt string) string {
	hashedPassword := sha256.Sum256(([]byte(password + salt)))
	return hex.EncodeToString(hashedPassword[:])
}

func VerifyPassword(password, hashedPassword string) (bool, error) {
	hashedPasswordAndSalt := strings.Split(hashedPassword, ":")
	_, err := hex.DecodeString(hashedPasswordAndSalt[0])
	if err != nil {
		return false, err
	}
	salt, err := hex.DecodeString(hashedPasswordAndSalt[1])
	if err != nil {
		return false, err
	}
	inputHashedPassword := sha256.Sum256([]byte(password + hex.EncodeToString(salt)))

	return hex.EncodeToString(inputHashedPassword[:]) == hashedPasswordAndSalt[0], nil
}

func EncryptingUserPassword(password string) string {
	var config dto.Config
	configData, _ := operator.LoadJson(config, "config/env/config.json")
	userPasswordSalt := (configData.(map[string]interface{})["user"]).(map[string]interface{})["password_salt"].(string)
	encodeUserPassword := HashPassword(password, userPasswordSalt)

	return encodeUserPassword
}

// func EncryptingLoginToken(authAccount string, authPassword string) string {
// 	var config dto.Config

// 	authorizationContent := base64.StdEncoding.EncodeToString([]byte(authAccount + ":" + base64.StdEncoding.EncodeToString([]byte(authPassword))))

// 	// Creating login token
// 	configData, _ := operator.LoadJson(config, "config/env/config.json")
// 	loginTokenSalt := (configData.(map[string]interface{})["user"]).(map[string]interface{})["login_token_salt"].(string)
// 	loginToken, _ := HashPassword(authorizationContent, loginTokenSalt)

// 	return loginToken
// }

func EncryptingLoginToken(authAccount string, authPassword string) string {
	authorizationContent := base64.StdEncoding.EncodeToString([]byte(authAccount + ":" + base64.StdEncoding.EncodeToString([]byte(authPassword))))

	// Creating login token
	byteSalt := make([]byte, 16)
	rand.Read(byteSalt)
	loginTokenSalt := hex.EncodeToString(byteSalt)
	loginToken := HashPassword(authorizationContent, loginTokenSalt)

	return loginToken
}

func LoginAuthParser(authorization string) ([]map[string]interface{}, error) {
	// Parsing user account and password
	authorizationContent := strings.Split(authorization, "Basic ")[1]
	decodeAuthorization, _ := base64.StdEncoding.DecodeString(authorizationContent)
	authAccount := strings.Split(string(decodeAuthorization), ":")[0]
	decodeAuthPassword, _ := base64.StdEncoding.DecodeString(strings.Split(string(decodeAuthorization), ":")[1])
	authPassword := string(decodeAuthPassword)

	// Encrypting password
	encodeAuthPassword := EncryptingUserPassword(authPassword)

	return []map[string]interface{}{
		0: {"auth_account": authAccount, "encode_auth_password": encodeAuthPassword},
	}, nil
}
