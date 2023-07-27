package mysqldb

import (
	"fmt"
	sqlOperator "server/utils/sqloperator"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(ip string, port int, userName string, password string) sqlOperator.ISqlDB {
	const driver string = "mysql"
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", userName, password, ip, port)
	db, _ := sqlOperator.Open(driver, conn)

	newSqlDB := sqlOperator.NewSqlDB(db)
	if err := db.Ping(); err != nil {
		fmt.Println("Failed to open MySQL DB: ", err)
		return nil
	}
	fmt.Println("MySQL DB connection is successful!")

	return newSqlDB
}
