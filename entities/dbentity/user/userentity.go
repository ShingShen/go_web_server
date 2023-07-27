package userentity

import (
	rdbOperator "server/utils/rdboperator"
	sqlOperator "server/utils/sqloperator"
)

func CreateUserColumns(db sqlOperator.ISqlDB, dbName string, user string) {
	buildDB, _ := rdbOperator.GetBuildDB(db, dbName, "buildDB")

	// Adding column
	buildDB.AddColumn(user, "user_account", "VARCHAR(64)")
	buildDB.AddColumn(user, "user_password", "VARCHAR(512)")
	buildDB.AddColumn(user, "email", "VARCHAR(64)")
	buildDB.AddColumn(user, "phone", "VARCHAR(64) NOT NULL DEFAULT ''")
	buildDB.AddColumn(user, "user_profile", "VARCHAR(512) NOT NULL DEFAULT ''")
	buildDB.AddColumn(user, "first_name", "VARCHAR(64) NOT NULL DEFAULT ''")
	buildDB.AddColumn(user, "last_name", "VARCHAR(64) NOT NULL DEFAULT ''")
	buildDB.AddColumn(user, "gender", "TINYINT UNSIGNED NOT NULL DEFAULT 0")
	buildDB.AddColumn(user, "birthday", "DATE")
	buildDB.AddColumn(user, "height", "SMALLINT UNSIGNED NOT NULL DEFAULT 0")
	buildDB.AddColumn(user, "med_compliance", "TINYINT UNSIGNED NOT NULL DEFAULT 0")
	buildDB.AddColumn(user, "alert", "TINYINT(1) UNSIGNED NOT NULL DEFAULT 0")
	buildDB.AddColumn(user, "allergies", "VARCHAR(64) NOT NULL DEFAULT ''")
	buildDB.AddColumn(user, "role", "TINYINT UNSIGNED NOT NULL DEFAULT 1")
	buildDB.AddColumn(user, "enabled", "TINYINT UNSIGNED NOT NULL DEFAULT 0")
	buildDB.AddColumn(user, "created_time", "TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP")
	buildDB.AddColumn(user, "updated_time", "TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")

	// Adding index
	buildDB.AddUnique(user, "u_user_account", "user_account")
	buildDB.AddUnique(user, "u_email", "email")
}
