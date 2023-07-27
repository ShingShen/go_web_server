package mysqldb

type ISqlOperator interface {
	Close() error
	Begin() (ISqlTx, error)
	Exec(query string, args ...any) (ISqlResult, error)
	Query(query string, args ...any) (ISqlRows, error)
	QueryRow(query string, args ...any) ISqlRow
}

type ISqlTx interface {
	Rollback() error
}

type ISqlResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type ISqlRows interface{}

type ISqlRow interface {
	Scan(dest ...any) error
}
