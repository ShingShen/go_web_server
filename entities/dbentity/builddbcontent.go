package dbentity

import (
	"fmt"

	loginTokenEntity "server/entities/dbentity/logintoken"
	memoEntity "server/entities/dbentity/memo"
	scheduleentity "server/entities/dbentity/schedule"
	userEntity "server/entities/dbentity/user"
	rdbOperator "server/utils/rdboperator"
	sqlOperator "server/utils/sqloperator"
)

type buildDbContentFactory struct{}

func (b *buildDbContentFactory) getBuildDbContent(db sqlOperator.ISqlDB, dbName string, name string) (IBuildDbContent, error) {
	if name == "buildDbContent" {
		return &buildDbContent{
			db:     db,
			dbName: dbName,
		}, nil
	}
	return nil, fmt.Errorf("wrong build db content type passed")
}

type buildDbContent struct {
	db     sqlOperator.ISqlDB
	dbName string
}

func (b *buildDbContent) buildTables() {
	buildDB, _ := rdbOperator.GetBuildDB(b.db, b.dbName, "buildDB")
	buildDB.CreateTable("login_token", "login_token_id")
	buildDB.CreateTable("memo", "memo_id")
	buildDB.CreateTable("schedule", "schedule_id")
	buildDB.CreateTable("user", "user_id")
}

func (b *buildDbContent) buildColumns() {
	loginTokenEntity.CreateLoginTokenColumn(b.db, b.dbName, "login_token")
	memoEntity.CreateMemoColumns(b.db, b.dbName, "memo")
	scheduleentity.CreateScheduleColumns(b.db, b.dbName, "schedule")
	userEntity.CreateUserColumns(b.db, b.dbName, "user")
}
