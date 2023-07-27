package rdboperator

type IBuildDB interface {
	CreateDB()
	CreateTable(tableName string, id string)
	AddColumn(tableName string, columnName string, content string)
	AddIndex(tableName string, indexName string, columnName string)
	AddUnique(tableName string, indexName string, columnName string)
}
