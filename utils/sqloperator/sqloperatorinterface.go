package sqloperator

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

// mockgen -source utils/sqloperator/sqloperatorinterface.go -destination tests/mocks/mocksqldbinterface.go

type ISqlDB interface {
	Begin() (ISqlTx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (ISqlTx, error)
	Close() error
	Conn(ctx context.Context) (ISqlConn, error)
	Driver() IDriver
	Exec(query string, args ...any) (ISqlResult, error)
	ExecContext(ctx context.Context, query string, args ...any) (ISqlResult, error)
	Ping() error
	PingContext(ctx context.Context) error
	Prepare(query string) (ISqlStmt, error)
	Query(query string, args ...any) (ISqlRows, error)
	QueryContext(ctx context.Context, query string, args ...any) (ISqlRows, error)
	QueryRow(query string, args ...any) ISqlRow
	QueryRowContext(ctx context.Context, query string, args ...any) ISqlRow
	SetConnMaxIdleTime(d time.Duration)
	SetConnMaxLifetime(d time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Stats() ISqlDBStats
}

type IDb interface {
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
	Conn(ctx context.Context) (*sql.Conn, error)
	Driver() driver.Driver
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Ping() error
	PingContext(ctx context.Context) error
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	SetConnMaxIdleTime(d time.Duration)
	SetConnMaxLifetime(d time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Stats() sql.DBStats
}

type ISqlTx interface {
	Commit() error
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Rollback() error
	Stmt(stmt *sql.Stmt) *sql.Stmt
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}

type ISqlConn interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PingContext(ctx context.Context) error
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Raw(f func(driverConn any) error) (err error)
}

type IDriver interface {
	Open(name string) (driver.Conn, error)
}

type ISqlStmt interface {
	Close() error
	Exec(args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, args ...any) (sql.Result, error)
	Query(args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, args ...any) (*sql.Rows, error)
	QueryRow(args ...any) *sql.Row
	QueryRowContext(ctx context.Context, args ...any) *sql.Row
}

type ISqlDBStats interface{}

type ISqlResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type ISqlRows interface {
	Close() error
	ColumnTypes() ([]*sql.ColumnType, error)
	Columns() ([]string, error)
	Err() error
	Next() bool
	NextResultSet() bool
	Scan(dest ...any) error
}

type ISqlRow interface {
	Err() error
	Scan(dest ...any) error
}
