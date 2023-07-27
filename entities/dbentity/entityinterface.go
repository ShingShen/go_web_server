package dbentity

import sqlOperator "server/utils/sqloperator"

type IEntities interface {
	CreateTables(dbName string)
}

type IBuildDbContentFactory interface {
	getBuildDbContent(db sqlOperator.ISqlDB, dbName string, name string) (IBuildDbContent, error)
}

type IBuildDbContent interface {
	buildTables()
	buildColumns()
}
