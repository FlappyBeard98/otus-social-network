package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/georgysavva/scany/v2/sqlscan"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/ilyakaznacheev/cleanenv"
)

type SqlQuery struct {
	sql    string
	params []any
}

func NewSqlQuery(sql string, params ...any) *SqlQuery {
	return &SqlQuery{sql: sql, params: params}
}

func (o *SqlQuery) QueryOne(ctx context.Context, db Db, dst any) error {
	rows, err := db.QueryContext(ctx, o.sql, o.params...)
	if err != nil {
		return fmt.Errorf("scany: query multiple result rows: %w", err)
	}
	if err := sqlscan.ScanOne(dst, rows); err != nil {
		return fmt.Errorf("scanning one: %w", err)
	}
	return nil
}

func (o *SqlQuery) Query(ctx context.Context, db Db, dst any) error {
	err := sqlscan.Select(ctx, db, dst, o.sql, o.params...)

	return err
}

func (o *SqlQuery) Exec(ctx context.Context, db Db) (ExecResult, error) {
	result := ExecResult{}
	r, err := db.Exec(o.sql, o.params...)

	if err != nil {
		return result, err
	}

	if n, err := r.LastInsertId(); err != nil {
		result.LastInsertId = -1
	} else {
		result.LastInsertId = n
	}

	if n, err := r.RowsAffected(); err != nil {
		result.RowsAffected = -1
	} else {
		result.RowsAffected = n
	}

	return result, nil
}

type ExecResult struct {
	LastInsertId int64
	RowsAffected int64
}

// Like prepares LIKE-statement for sql query
func Like(str *string, pre bool, post bool) *string {

	if str == nil {
		return nil
	}
	s := *str

	if pre {
		s = "%" + s
	}

	if post {
		s = s + "%"
	}

	return &s
}

func BeginTxFunc(ctx context.Context, opts *sql.TxOptions, db *sql.DB, fn func(context.Context, *sql.Tx) error) error {

	tx, err := db.BeginTx(ctx, opts)

	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	err = fn(ctx, tx)

	if err != nil {
		return err
	}

	return nil
}

type Db interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type DbConfig struct {
	ConnectionString string `json:"connectionString" env:"CONNECTION_STRING"`
}


func Connect(cfg DbConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.ConnectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
