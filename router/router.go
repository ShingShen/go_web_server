package router

import (
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"

	sqlOperator "server/utils/sqloperator"
)

func NewServer(db sqlOperator.ISqlDB, name string) (IServer, error) {
	if name == "server" {
		return &server{db, &handlerFactory{}}, nil
	}
	return nil, fmt.Errorf("failed to create new server")
}

type server struct {
	db             sqlOperator.ISqlDB
	handlerFactory IHandlerFactory
}

func (s *server) RunServer(rdb *redis.Client, cloudStorage *storage.Client, port string) {
	tx, _ := s.db.Begin()
	defer tx.Rollback()

	// cronJobs.OneMin(db, rdb)

	handler, _ := s.handlerFactory.getHandler(s.db, "handler")
	mux := handler.muxHandlers(rdb, cloudStorage)
	handler.listen(mux, port)
}
