package router

import (
	"fmt"
	"net/http"
	gmailSenderController "server/controllers/dbcontroller/gmailsender"
	memoController "server/controllers/dbcontroller/memo"
	scheduleController "server/controllers/dbcontroller/schedule"
	storageFileController "server/controllers/dbcontroller/storagefile"
	userController "server/controllers/dbcontroller/user"
	sqlOperator "server/utils/sqloperator"
	"server/web"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"
)

type handleFuncFactory struct{}

func (h *handleFuncFactory) getHandleFunc(db sqlOperator.ISqlDB, mux *http.ServeMux, name string) (IHandleFunc, error) {
	if name == "handleFunc" {
		return &handleFunc{
			db:                           db,
			mux:                          mux,
			gmailSenderControllerFactory: &gmailSenderController.GmailSenderControllerFactory{},
			memoControllerFactory:        &memoController.MemoControllerFactory{},
			scheduleControllerFactory:    &scheduleController.ScheduleControllerFactory{},
			storageFileControllerFactory: &storageFileController.StorageFileControllerFactory{},
			userControllerFactory:        &userController.UserControllerFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong handle func type passed")
}

type handleFunc struct {
	db                           sqlOperator.ISqlDB
	mux                          *http.ServeMux
	gmailSenderControllerFactory gmailSenderController.IGmailSenderControllerFactory
	memoControllerFactory        memoController.IMemoControllerFactory
	scheduleControllerFactory    scheduleController.IScheduleControllerFactory
	storageFileControllerFactory storageFileController.IStorageFileControllerFactory
	userControllerFactory        userController.IUserControllerFactory
}

func (h *handleFunc) webHandleFunc() {
	h.mux.HandleFunc("/web", web.HomePage)
}

func (h *handleFunc) gmailSenderHandleFunc(dir string) {
	controller, _ := h.gmailSenderControllerFactory.GetGmailSenderController(h.db, "gmailSenderController")
	h.mux.HandleFunc(dir+"/send", controller.GmailSender())
}

func (h *handleFunc) memoHandleFunc(dir string) {
	controller, _ := h.memoControllerFactory.GetMemoController(h.db, "memoController")
	h.mux.HandleFunc(dir+"/create", controller.Create())
	h.mux.HandleFunc(dir+"/update", controller.Update())
	h.mux.HandleFunc(dir+"/get", controller.Get())
	h.mux.HandleFunc(dir+"/get_memos_by_user_id", controller.GetMemosByUserId())
	h.mux.HandleFunc(dir+"/delete", controller.Delete())
}

func (h *handleFunc) scheduleHandleFunc(dir string) {
	controller, _ := h.scheduleControllerFactory.GetScheduleController(h.db, "scheduleController")
	h.mux.HandleFunc(dir+"/create", controller.Create())
	h.mux.HandleFunc(dir+"/update", controller.Update())
	h.mux.HandleFunc(dir+"/get", controller.GetAnEvent())
	h.mux.HandleFunc(dir+"/get_day", controller.GetOneDayEvents())
	h.mux.HandleFunc(dir+"/get_month", controller.GetOneMonthEvents())
	h.mux.HandleFunc(dir+"/delete", controller.Delete())
}

func (h *handleFunc) storageFileHandleFunc(cloudStorage *storage.Client, dir string) {
	controller, _ := h.storageFileControllerFactory.GetStorageFileController(h.db, "storageFileController")
	h.mux.HandleFunc(dir+"/upload", controller.Upload(cloudStorage))
}

func (h *handleFunc) userHandleFunc(rdb *redis.Client, cloudStorage *storage.Client, dir string) {
	controller, _ := h.userControllerFactory.GetUserController(h.db, "userController")
	h.mux.HandleFunc(dir+"/create", controller.Create(rdb))
	h.mux.HandleFunc(dir+"/create_login_token", controller.CreateLoginToken(rdb))
	h.mux.HandleFunc(dir+"/update_login_token", controller.UpdateLoginToken(rdb))
	h.mux.HandleFunc(dir+"/update", controller.Update())
	h.mux.HandleFunc(dir+"/upload_user_profile", controller.UploadUserProfile(cloudStorage))
	h.mux.HandleFunc(dir+"/reset_user_password", controller.ResetUserPassword())
	h.mux.HandleFunc(dir+"/get", controller.Get())
	h.mux.HandleFunc(dir+"/get_by_account", controller.GetUserByAccount())
	h.mux.HandleFunc(dir+"/get_all", controller.GetAllUsers())
	h.mux.HandleFunc(dir+"/get_specific_roles", controller.GetSpecificRoles())
	h.mux.HandleFunc(dir+"/get_login_token", controller.GetLoginTokenByUserId(rdb))
	h.mux.HandleFunc(dir+"/delete", controller.Delete())
	h.mux.HandleFunc(dir+"/login", controller.Login())
}
