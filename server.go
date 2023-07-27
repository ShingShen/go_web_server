package main

import (
	"context"
	"server/config/redis"
	"server/router"

	cloudStorage "server/config/firebase/cloudstorage"
	mySqlDB "server/config/mysqldb"
	dbEntity "server/entities/dbentity"
	rdbOperator "server/utils/rdboperator"
)

func main() {
	db := mySqlDB.Connect("mysql_ip", 3306, "user_name", "password")
	defer db.Close()

	rdb := redis.Connect("redis_ip", "")
	defer rdb.Close()

	ctx := context.Background()
	cloudStorage, _ := cloudStorage.Connect(ctx)
	defer cloudStorage.Close()

	buildDB, _ := rdbOperator.GetBuildDB(db, "mysql_db", "buildDB")
	createTables, _ := dbEntity.GetCreateTables(db, "createTables")
	router, _ := router.NewServer(db, "server")

	buildDB.CreateDB()
	createTables.CreateTables("mysql_db")
	router.RunServer(rdb, cloudStorage, "3301")
}
