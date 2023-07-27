package memoentity

import (
	rdbOperator "server/utils/rdboperator"
	sqlOperator "server/utils/sqloperator"
)

func CreateMemoColumns(db sqlOperator.ISqlDB, dbName string, memo string) {
	buildDB, _ := rdbOperator.GetBuildDB(db, dbName, "buildDB")

	// Adding column
	buildDB.AddColumn(memo, "content", "VARCHAR(128) NOT NULL DEFAULT ''")
	buildDB.AddColumn(memo, "has_read", "TINYINT UNSIGNED NOT NULL DEFAULT 0")
	buildDB.AddColumn(memo, "created_by", "VARCHAR(64)")
	buildDB.AddColumn(memo, "updated_by", "VARCHAR(64)")
	buildDB.AddColumn(memo, "user_id", "BIGINT(20) DEFAULT 0, ADD FOREIGN KEY(user_id) REFERENCES mysql_db.user(user_id)")
	buildDB.AddColumn(memo, "created_time", "TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP")
	buildDB.AddColumn(memo, "updated_time", "TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")
}
