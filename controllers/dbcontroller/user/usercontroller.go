package usercontroller

import (
	"fmt"
	"net/http"
	"server/middleware"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"

	sqlOperator "server/utils/sqloperator"
)

type UserControllerFactory struct{}

func (u *UserControllerFactory) GetUserController(db sqlOperator.ISqlDB, name string) (IUserController, error) {
	if name == "userController" {
		return &UserController{
			db:                        db,
			userControllerFuncFactory: &userControllerFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong user controller type passed")
}

type UserController struct {
	db                        sqlOperator.ISqlDB
	userControllerFuncFactory IUserControllerFuncFactory
}

func (u *UserController) Create(rdb *redis.Client) http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.createController(rdb)
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"POST", controller,
			),
		),
	)
}

func (u *UserController) CreateLoginToken(rdb *redis.Client) http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.createLoginTokenController(rdb)
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"POST", middleware.AllUserAuth(u.db, controller),
			),
		),
	)
}

func (u *UserController) UpdateLoginToken(rdb *redis.Client) http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.updateLoginTokenController(rdb)
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"PUT", middleware.AllUserAuth(u.db, controller),
			),
		),
	)
}

func (u *UserController) Update() http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.updateController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"PUT", middleware.AllUserAuth(u.db, controller),
			),
		),
	)
}

func (u *UserController) UploadUserProfile(cloudStorage *storage.Client) http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.uploadUserProfileController(cloudStorage)
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.UploadFileMethod(
				middleware.AllUserAuth(u.db, controller),
			),
		),
	)
}

func (u *UserController) ResetUserPassword() http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.resetUserPasswordController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"PUT", middleware.AllUserAuth(u.db, controller),
			),
		),
	)
}

func (u *UserController) Enable() http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.enableUserController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"PUT", controller,
			),
		),
	)
}

func (u *UserController) Get() http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.getController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"GET", middleware.AllUserAuth(u.db, controller),
			),
		),
	)
}

func (u *UserController) GetUserByAccount() http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.getUserByAccountController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"GET", middleware.AllUserAuth(u.db, controller),
			),
		),
	)
}

func (u *UserController) GetAllUsers() http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.getAllUsersController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"GET", middleware.UserOnlyAuth(99, u.db, controller),
			),
		),
	)
}

func (u *UserController) GetSpecificRoles() http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.getSpecificRolesController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"GET", middleware.UserOnlyAuth(99, u.db, controller),
			),
		),
	)
}

func (u *UserController) GetLoginTokenByUserId(rdb *redis.Client) http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.getLoginTokenByUserIdController(rdb)
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"GET", middleware.AllUserAuth(u.db, controller),
			),
		),
	)
}

func (u *UserController) Delete() http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.deleteController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"DELETE", middleware.UserOnlyAuth(99, u.db, controller),
			),
		),
	)
}

func (u *UserController) Login() http.HandlerFunc {
	userControllerFunc, _ := u.userControllerFuncFactory.getUserControllerFunc(u.db, "userControllerFunc")
	controller := userControllerFunc.loginController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"POST", middleware.LoginAuth(u.db, controller),
			),
		),
	)
}
