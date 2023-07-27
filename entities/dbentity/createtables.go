package dbentity

import (
	"fmt"

	sqlOperator "server/utils/sqloperator"
)

func GetCreateTables(db sqlOperator.ISqlDB, name string) (IEntities, error) {
	if name == "createTables" {
		return &entities{
			db:                    db,
			buildDbContentFactory: &buildDbContentFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong createTables type passed")
}

type entities struct {
	db                    sqlOperator.ISqlDB
	buildDbContentFactory IBuildDbContentFactory
}

func (e *entities) CreateTables(dbName string) {
	tx, _ := e.db.Begin()
	defer tx.Rollback()

	buildDbContent, _ := e.buildDbContentFactory.getBuildDbContent(e.db, dbName, "buildDbContent")
	buildDbContent.buildTables()
	buildDbContent.buildColumns()
}
