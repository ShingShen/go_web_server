package gmailsendercontroller

import (
	"net/http"
	sqlOperator "server/utils/sqloperator"
)

type IGmailSenderControllerFactory interface {
	GetGmailSenderController(db sqlOperator.ISqlDB, name string) (IGmailSenderController, error)
}

type IGmailSenderController interface {
	GmailSender() http.HandlerFunc
}

type IGmailSenderControllerFuncFactory interface {
	getGmailSenderControllerFunc(name string) (IGmailSenderControllerFunc, error)
}

type IGmailSenderControllerFunc interface {
	gmailSenderController() func(w http.ResponseWriter, r *http.Request)
}
