package rdboperator

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestGetBuildDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	const dbName string = "dbName"
	const name string = "buildDB"
	const errName string = "errName"

	getBuildDB, err := GetBuildDB(mockDB, dbName, name)
	if err == nil {
		t.Logf("getBuildDB passed: %v, %v", getBuildDB, err)
	} else {
		t.Errorf("getBuildDB failed: %v, %v", getBuildDB, err)
	}

	getBuildDBNameErr, err := GetBuildDB(mockDB, dbName, errName)
	if err != nil {
		t.Logf("getBuildDBNameErr passed: %v, %v", getBuildDBNameErr, err)
	} else {
		t.Errorf("getBuildDBNameErr failed: %v, %v", getBuildDBNameErr, err)
	}
}

func TestCreateDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const dbName string = "dbName"
	sql := fmt.Sprintf("CREATE DATABASE %s\n", dbName)
	buildDB := &buildDB{mockDB, dbName}

	// failed to create db
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql).Return(nil, fmt.Errorf("result error"))
	buildDB.CreateDB()

	// db is created successfully
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	mockDB.EXPECT().Exec(sql).Return(mockSqlResult, nil)
	buildDB.CreateDB()
}

func TestCreateTable(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const dbName string = "dbName"
	const tableName string = "tableName"
	const tableId string = "tableId"
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (%s BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL);", dbName, tableName, tableId)
	buildDB := &buildDB{mockDB, dbName}

	// failed to create table
	mockDB.EXPECT().Exec(sql).Return(nil, fmt.Errorf("result error"))
	buildDB.CreateTable(tableName, tableId)

	// table is created successfully
	mockDB.EXPECT().Exec(sql).Return(mockSqlResult, nil)
	buildDB.CreateTable(tableName, tableId)
}

func TestAddColumn(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const dbName string = "dbName"
	const tableName string = "tableName"
	const columnName string = "columnName"
	const content string = "content"
	sql := fmt.Sprintf("ALTER TABLE %s.%s ADD %s %s;", dbName, tableName, columnName, content)
	buildDB := &buildDB{mockDB, dbName}

	// failed to add column
	mockDB.EXPECT().Exec(sql).Return(nil, fmt.Errorf("result error"))
	buildDB.AddColumn(tableName, columnName, content)

	// column is created successfully
	mockDB.EXPECT().Exec(sql).Return(mockSqlResult, nil)
	buildDB.AddColumn(tableName, columnName, content)
}

func TestAddIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const dbName string = "dbName"
	const tableName string = "tableName"
	const indexName string = "indexName"
	const columnName string = "columnName"
	sql := fmt.Sprintf("ALTER TABLE %s.%s ADD INDEX %s (%s);", dbName, tableName, indexName, columnName)
	buildDB := &buildDB{mockDB, dbName}

	// Failed to add Index on column
	mockDB.EXPECT().Exec(sql).Return(nil, fmt.Errorf("result error"))
	buildDB.AddIndex(tableName, indexName, columnName)

	// Index on column is created successfully
	mockDB.EXPECT().Exec(sql).Return(mockSqlResult, nil)
	buildDB.AddIndex(tableName, indexName, columnName)
}

func TestAddUnique(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const dbName string = "dbName"
	const tableName string = "tableName"
	const indexName string = "indexName"
	const columnName string = "columnName"
	sql := fmt.Sprintf("ALTER TABLE %s.%s ADD UNIQUE %s (%s);", dbName, tableName, indexName, columnName)
	buildDB := &buildDB{mockDB, dbName}

	// "Failed to add Unique Index on column
	mockDB.EXPECT().Exec(sql).Return(nil, fmt.Errorf("result error"))
	buildDB.AddUnique(tableName, indexName, columnName)

	// Unique Index on column is created successfully
	mockDB.EXPECT().Exec(sql).Return(mockSqlResult, nil)
	buildDB.AddUnique(tableName, indexName, columnName)
}
