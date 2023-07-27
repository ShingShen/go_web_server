package gmailsendercontroller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	gmailSenderService "server/services/dbservice/service/gmailsender"
)

type gmailSenderControllerFuncFactory struct{}

func (g *gmailSenderControllerFuncFactory) getGmailSenderControllerFunc(name string) (IGmailSenderControllerFunc, error) {
	if name == "gmailSenderControllerFunc" {
		return &gmailSenderControllerFunc{
			gmailSenderServiceFactory: &gmailSenderService.GmailSenderServiceFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong gmail sender controller func type passed")
}

type gmailSenderControllerFunc struct {
	gmailSenderServiceFactory gmailSenderService.IGmailSenderServiceFactory
}

func (g *gmailSenderControllerFunc) gmailSenderController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)
		getService, _ := g.gmailSenderServiceFactory.GetGmailSenderService("gmailSender")
		service := getService.GmailSender(
			res["recipient"].(string),
			res["msg"].(string),
		)
		if service != nil {
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(204)
	}
}
