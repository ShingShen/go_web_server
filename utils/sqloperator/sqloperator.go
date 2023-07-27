package sqloperator

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewSqlDB(db IDb) ISqlDB {
	return &SqlDB{
		db: db,
	}
}

func Open(driverName, dataSourceName string) (IDb, error) {
	return sql.Open(driverName, dataSourceName)
}

type SqlDB struct {
	db IDb
	// db *sql.DB
}

func (s *SqlDB) Begin() (ISqlTx, error) {
	return s.db.Begin()
}

func (s *SqlDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (ISqlTx, error) {
	return s.db.BeginTx(ctx, opts)
}

func (s *SqlDB) Close() error {
	return s.db.Close()
}

func (s *SqlDB) Conn(ctx context.Context) (ISqlConn, error) {
	return s.db.Conn(ctx)
}

func (s *SqlDB) Driver() IDriver {
	return s.db.Driver()
}

func (s *SqlDB) Exec(query string, args ...any) (ISqlResult, error) {
	return s.db.Exec(query, args...)
}

func (s *SqlDB) ExecContext(ctx context.Context, query string, args ...any) (ISqlResult, error) {
	return s.db.ExecContext(ctx, query, args...)
}

func (s *SqlDB) Ping() error {
	return s.db.Ping()
}

func (s *SqlDB) PingContext(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *SqlDB) Prepare(query string) (ISqlStmt, error) {
	return s.db.Prepare(query)
}

func (s *SqlDB) Query(query string, args ...any) (ISqlRows, error) {
	return s.db.Query(query, args...)
}

func (s *SqlDB) QueryContext(ctx context.Context, query string, args ...any) (ISqlRows, error) {
	return s.db.QueryContext(ctx, query, args...)
}

func (s *SqlDB) QueryRow(query string, args ...any) ISqlRow {
	return s.db.QueryRow(query, args...)
}

func (s *SqlDB) QueryRowContext(ctx context.Context, query string, args ...any) ISqlRow {
	return s.db.QueryRowContext(ctx, query, args...)
}

func (s *SqlDB) SetConnMaxIdleTime(d time.Duration) {
	s.db.SetConnMaxIdleTime(d)
}

func (s *SqlDB) SetConnMaxLifetime(d time.Duration) {
	s.db.SetConnMaxLifetime(d)
}

func (s *SqlDB) SetMaxIdleConns(n int) {
	s.db.SetMaxIdleConns(n)
}

func (s *SqlDB) SetMaxOpenConns(n int) {
	s.db.SetMaxOpenConns(n)
}

func (s *SqlDB) Stats() ISqlDBStats {
	return s.db.Stats()
}
