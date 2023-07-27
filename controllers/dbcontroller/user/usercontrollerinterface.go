package usercontroller

import (
	"net/http"
	sqlOperator "server/utils/sqloperator"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"
)

type IUserControllerFactory interface {
	GetUserController(db sqlOperator.ISqlDB, name string) (IUserController, error)
}

type IUserController interface {
	Create(rdb *redis.Client) http.HandlerFunc
	CreateLoginToken(rdb *redis.Client) http.HandlerFunc
	UpdateLoginToken(rdb *redis.Client) http.HandlerFunc
	Update() http.HandlerFunc
	UploadUserProfile(cloudStorage *storage.Client) http.HandlerFunc
	ResetUserPassword() http.HandlerFunc
	Enable() http.HandlerFunc
	Get() http.HandlerFunc
	GetUserByAccount() http.HandlerFunc
	GetAllUsers() http.HandlerFunc
	GetSpecificRoles() http.HandlerFunc
	GetLoginTokenByUserId(rdb *redis.Client) http.HandlerFunc
	Delete() http.HandlerFunc
	Login() http.HandlerFunc
}

type IUserControllerFuncFactory interface {
	getUserControllerFunc(db sqlOperator.ISqlDB, name string) (IUserControllerFunc, error)
}

type IUserControllerFunc interface {
	createController(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request)
	createLoginTokenController(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request)
	updateLoginTokenController(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request)
	updateController() func(w http.ResponseWriter, r *http.Request)
	uploadUserProfileController(cloudStorage *storage.Client) func(w http.ResponseWriter, r *http.Request)
	resetUserPasswordController() func(w http.ResponseWriter, r *http.Request)
	enableUserController() func(w http.ResponseWriter, r *http.Request)
	getController() func(w http.ResponseWriter, r *http.Request)
	getUserByAccountController() func(w http.ResponseWriter, r *http.Request)
	getAllUsersController() func(w http.ResponseWriter, r *http.Request)
	getSpecificRolesController() func(w http.ResponseWriter, r *http.Request)
	getLoginTokenByUserIdController(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request)
	deleteController() func(w http.ResponseWriter, r *http.Request)
	loginController() func(w http.ResponseWriter, r *http.Request)
}
