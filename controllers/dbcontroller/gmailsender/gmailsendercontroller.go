package gmailsendercontroller

import (
	"fmt"
	"net/http"
	"server/middleware"

	sqlOperator "server/utils/sqloperator"
)

type GmailSenderControllerFactory struct{}

func (g *GmailSenderControllerFactory) GetGmailSenderController(db sqlOperator.ISqlDB, name string) (IGmailSenderController, error) {
	if name == "gmailSenderController" {
		return &GmailSenderController{
			db:                               db,
			gmailSenderControllerFuncFactory: &gmailSenderControllerFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong gmail sender controller type passed")
}

type GmailSenderController struct {
	db                               sqlOperator.ISqlDB
	gmailSenderControllerFuncFactory IGmailSenderControllerFuncFactory
}

func (g *GmailSenderController) GmailSender() http.HandlerFunc {
	gmailSenderControllerFunc, _ := g.gmailSenderControllerFuncFactory.getGmailSenderControllerFunc("gmailSenderControllerFunc")
	controller := gmailSenderControllerFunc.gmailSenderController()
	return http.HandlerFunc(
		middleware.RunHandler(
			middleware.Method(
				"POST", middleware.AllUserAuth(g.db, controller),
			),
		),
	)
}
