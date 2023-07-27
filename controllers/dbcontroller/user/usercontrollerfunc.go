package usercontroller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/dto"
	"server/utils/operator"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"

	userService "server/services/dbservice/service/user"
	sqlOperator "server/utils/sqloperator"
)

type userControllerFuncFactory struct{}

func (u *userControllerFuncFactory) getUserControllerFunc(db sqlOperator.ISqlDB, name string) (IUserControllerFunc, error) {
	if name == "userControllerFunc" {
		return &userControllerFunc{
			db:                 db,
			userServiceFactory: &userService.UserServiceFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong user controller func type passed")
}

type userControllerFunc struct {
	db                 sqlOperator.ISqlDB
	userServiceFactory userService.IUserServiceFactory
}

func (u *userControllerFunc) createController(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		service := getService.CreateUser(
			rdb,
			dto.User{
				UserAccount:  res["user_account"].(string),
				UserPassword: res["user_password"].(string),
				FirstName:    operator.IfElse(res["first_name"] != nil, res["first_name"], "").(string),
				LastName:     operator.IfElse(res["last_name"] != nil, res["last_name"], "").(string),
				Email:        res["email"].(string),
				Phone:        operator.IfElse(res["phone"] != nil, res["phone"], "").(string),
				Gender:       uint8(operator.IfElse(res["gender"] != nil, res["gender"], 0.0).(float64)),
				Birthday:     operator.IfElse(res["birthday"] != nil, res["birthday"], "1000-01-01").(string),
				Role:         uint8(res["role"].(float64)),
			},
		)
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(200)
	}
}

func (u *userControllerFunc) createLoginTokenController(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		service := getService.CreateLoginToken(
			rdb,
			dto.User{
				UserAccount:  res["user_account"].(string),
				UserPassword: res["user_password"].(string),
			},
		)
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(200)
	}
}

func (u *userControllerFunc) updateLoginTokenController(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		service := getService.UpdateLoginToken(
			rdb,
			dto.User{
				UserAccount:  res["user_account"].(string),
				UserPassword: res["user_password"].(string),
			},
		)
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(200)
	}
}

func (u *userControllerFunc) updateController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		service := getService.UpdateUser(
			uint64(userId),
			dto.User{
				FirstName: operator.IfElse(res["first_name"] != nil, res["first_name"], "").(string),
				LastName:  operator.IfElse(res["last_name"] != nil, res["last_name"], "").(string),
				Gender:    uint8(operator.IfElse(res["gender"] != nil, res["gender"], 0.0).(float64)),
				Birthday:  operator.IfElse(res["birthday"] != nil, res["birthday"], "1000-01-01").(string),
				Email:     operator.IfElse(res["email"] != nil, res["email"], "").(string),
				Phone:     operator.IfElse(res["phone"] != nil, res["phone"], "").(string),
			},
		)
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(204)
	}
}

func (u *userControllerFunc) uploadUserProfileController(cloudStorage *storage.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))

		r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		getService.UploadUserProfile(file, uint64(userId), "profileImage", cloudStorage)
	}
}

func (u *userControllerFunc) resetUserPasswordController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		service := getService.ResetUserPassword(
			uint64(userId),
			res["user_password"].(string),
		)
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(204)
	}
}

func (u *userControllerFunc) enableUserController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		service := getService.EnableUser(uint64(userId))
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(204)
	}
}

func (u *userControllerFunc) getController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		jsonData, err := getService.GetUser(uint64(userId))
		if err != nil {
			w.WriteHeader(400)
			return
		}

		w.Write(jsonData)
		w.WriteHeader(200)
	}
}

func (u *userControllerFunc) getUserByAccountController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userAccount := r.URL.Query().Get("user_account")

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		jsonData, err := getService.GetUserByAccount(userAccount)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		w.Write(jsonData)
		w.WriteHeader(200)
	}
}

func (u *userControllerFunc) getAllUsersController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		jsonData, err := getService.GetAllUsers()
		if err != nil {
			w.WriteHeader(400)
			return
		}

		w.Write(jsonData)
		w.WriteHeader(200)
	}
}

func (u *userControllerFunc) getSpecificRolesController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		role, _ := strconv.Atoi(r.URL.Query().Get("role"))

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		jsonData, err := getService.GetSpecificRoles(uint8(role))
		if err != nil {
			w.WriteHeader(400)
			return
		}

		w.Write(jsonData)
		w.WriteHeader(200)
	}
}

func (u *userControllerFunc) getLoginTokenByUserIdController(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		loginTokenData, err := getService.GetLoginTokenByUserId(rdb, uint64(userId))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		loginTokenJsonData, _ := json.Marshal(loginTokenData)

		w.Write([]byte(loginTokenJsonData))
		w.WriteHeader(200)
	}
}

func (u *userControllerFunc) deleteController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		service := getService.DeleteUser(uint64(userId))
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(204)
	}
}

func (u *userControllerFunc) loginController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		getService, _ := u.userServiceFactory.GetUserService(u.db, "user")
		loginJsonData, err := getService.Login(authorization)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		w.Write(loginJsonData)
		w.WriteHeader(200)
	}
}
