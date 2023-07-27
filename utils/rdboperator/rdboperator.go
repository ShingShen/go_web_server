package rdboperator

import (
	"fmt"

	sqlOperator "server/utils/sqloperator"
)

func GetBuildDB(db sqlOperator.ISqlDB, dbName string, name string) (IBuildDB, error) {
	if name == "buildDB" {
		return &buildDB{db, dbName}, nil
	}
	return nil, fmt.Errorf("wrong build db type passed")
}

type buildDB struct {
	db     sqlOperator.ISqlDB
	dbName string
}

func (b *buildDB) CreateDB() {
	tx, _ := b.db.Begin()
	defer tx.Rollback()

	sql := fmt.Sprintf("CREATE DATABASE %s\n", b.dbName)
	fmt.Printf("Creating %s...\n", b.dbName)

	_, err := b.db.Exec(sql)
	if err != nil {
		fmt.Printf("Failed to create %s: %v\n\n", b.dbName, err)
	} else {
		fmt.Printf("%s is created successfully!\n\n", b.dbName)
	}
}

func (b *buildDB) CreateTable(tableName string, id string) {
	createTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (%s BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL);", b.dbName, tableName, id)
	_, err := b.db.Exec(createTable)
	if err != nil {
		fmt.Printf("Failed to create %s table: %v\n", tableName, err)
	} else {
		fmt.Printf("%s table is created successfully!\n", tableName)
	}
}

func (b *buildDB) AddColumn(tableName string, columnName string, content string) {
	sql := fmt.Sprintf("ALTER TABLE %s.%s ADD %s %s;", b.dbName, tableName, columnName, content)
	_, err := b.db.Exec(sql)
	if err != nil {
		fmt.Printf("Failed to add %s %s column: %v\n", tableName, columnName, err)
	} else {
		fmt.Printf("%s %s column is created successfully!\n", tableName, columnName)
	}
}

func (b *buildDB) AddIndex(tableName string, indexName string, columnName string) {
	sql := fmt.Sprintf("ALTER TABLE %s.%s ADD INDEX %s (%s);", b.dbName, tableName, indexName, columnName)
	_, err := b.db.Exec(sql)
	if err != nil {
		fmt.Printf("Failed to add Index %s on column %s %s: %v\n", indexName, tableName, columnName, err)
	} else {
		fmt.Printf("Index %s on column %s %s is created successfully!\n", indexName, tableName, columnName)
	}
}

func (b *buildDB) AddUnique(tableName string, indexName string, columnName string) {
	sql := fmt.Sprintf("ALTER TABLE %s.%s ADD UNIQUE %s (%s);", b.dbName, tableName, indexName, columnName)
	_, err := b.db.Exec(sql)
	if err != nil {
		fmt.Printf("Failed to add Unique Index %s on column %s %s: %v\n", indexName, tableName, columnName, err)
	} else {
		fmt.Printf("Unique Index %s on column %s %s is created successfully!\n", indexName, tableName, columnName)
	}
}
