package logintokenentity

import (
	rdbOperator "server/utils/rdboperator"
	sqlOperator "server/utils/sqloperator"
)

func CreateLoginTokenColumn(db sqlOperator.ISqlDB, dbName string, loginToken string) {
	buildDB, _ := rdbOperator.GetBuildDB(db, dbName, "buildDB")

	// Adding column
	buildDB.AddColumn(loginToken, "login_token", "VARCHAR(512)")
	buildDB.AddColumn(loginToken, "user_id", "BIGINT(20) DEFAULT 0, ADD FOREIGN KEY(user_id) REFERENCES mysql_db.user(user_id)")
	buildDB.AddColumn(loginToken, "created_time", "TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP")
	buildDB.AddColumn(loginToken, "updated_time", "TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")

	// Adding index
	buildDB.AddUnique(loginToken, "u_user_id", "user_id")
}
