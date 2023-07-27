package scheduleentity

import (
	rdbOperator "server/utils/rdboperator"
	sqlOperator "server/utils/sqloperator"
)

func CreateScheduleColumns(db sqlOperator.ISqlDB, dbName string, schedule string) {
	buildDB, _ := rdbOperator.GetBuildDB(db, dbName, "buildDB")

	// Adding column
	buildDB.AddColumn(schedule, "title", "VARCHAR(128) NOT NULL DEFAULT ''")
	buildDB.AddColumn(schedule, "note", "VARCHAR(128) NOT NULL DEFAULT ''")
	buildDB.AddColumn(schedule, "start_time", "DATETIME NOT NULL")
	buildDB.AddColumn(schedule, "end_time", "DATETIME NOT NULL")
	buildDB.AddColumn(schedule, "user_id", "BIGINT(20) DEFAULT 0, ADD FOREIGN KEY(user_id) REFERENCES mysql_db.user(user_id)")
	buildDB.AddColumn(schedule, "created_time", "TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP")
	buildDB.AddColumn(schedule, "updated_time", "TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")
}
