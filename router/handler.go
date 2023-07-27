package router

import (
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"

	sqlOperator "server/utils/sqloperator"
)

type handlerFactory struct{}

func (h *handlerFactory) getHandler(db sqlOperator.ISqlDB, name string) (IHandler, error) {
	if name == "handler" {
		return &handler{
			db:                db,
			mux:               http.NewServeMux(),
			handleFuncFactory: &handleFuncFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong handler type passed ")
}

type handler struct {
	db                sqlOperator.ISqlDB
	mux               *http.ServeMux
	handleFuncFactory IHandleFuncFactory
}

func (h *handler) muxHandlers(rdb *redis.Client, cloudStorage *storage.Client) http.Handler {
	handleFunc, _ := h.handleFuncFactory.getHandleFunc(h.db, h.mux, "handleFunc")
	handleFunc.webHandleFunc()
	handleFunc.gmailSenderHandleFunc("/gmail")
	handleFunc.memoHandleFunc("/memo")
	handleFunc.scheduleHandleFunc("/schedule")
	handleFunc.storageFileHandleFunc(cloudStorage, "/storage_file")
	handleFunc.userHandleFunc(rdb, cloudStorage, "/user")

	return h.mux
}

func (h *handler) listen(mux http.Handler, port string) {
	httpServer := &http.Server{
		Handler:      mux,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Addr:         ":" + port,
	}

	fmt.Printf("Listening on port %s ......\n", port)

	listenAndServe := httpServer.ListenAndServe()
	if listenAndServe != nil {
		panic(listenAndServe)
	}
}
