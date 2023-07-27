package router

import (
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"

	sqlOperator "server/utils/sqloperator"
)

type IServer interface {
	RunServer(rdb *redis.Client, cloudStorage *storage.Client, port string)
}

type IHandlerFactory interface {
	getHandler(db sqlOperator.ISqlDB, name string) (IHandler, error)
}

type IHandler interface {
	muxHandlers(rdb *redis.Client, cloudStorage *storage.Client) http.Handler
	listen(mux http.Handler, port string)
}

type IHandleFuncFactory interface {
	getHandleFunc(db sqlOperator.ISqlDB, mux *http.ServeMux, name string) (IHandleFunc, error)
}

type IHandleFunc interface {
	webHandleFunc()
	gmailSenderHandleFunc(dir string)
	memoHandleFunc(dir string)
	scheduleHandleFunc(dir string)
	storageFileHandleFunc(cloudStorage *storage.Client, dir string)
	userHandleFunc(rdb *redis.Client, cloudStorage *storage.Client, dir string)
}
