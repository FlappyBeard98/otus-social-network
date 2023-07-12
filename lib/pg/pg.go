// Package mysql provides a set of functions to work with mysql database
package pg

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Db is a wrapper for sql.DB
type Db interface {
	// ExecContext executes a query without returning any rows.
	Query(context.Context,string,...any) (pgx.Rows, error)
	// QueryRow executes a query that is expected to return at most one row.
	QueryRow(context.Context, string, ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

// SqlQuery is a wrapper for sql query with params
type SqlQuery struct {
	sql    string
	params []any
}

// NewSqlQuery creates new SqlQuery
func NewSqlQuery(sql string, params ...any) *SqlQuery {
	return &SqlQuery{sql: sql, params: params}
}

// QueryOne executes query and scans one row into dst
func (o *SqlQuery) QueryOne(ctx context.Context,db Db, dst any) error {
	rows,err := db.Query(ctx,o.sql, o.params...)

	if err != nil {
		return fmt.Errorf("scany: query multiple result rows: %w", err)
	}

	if err := pgxscan.ScanOne(dst, rows); err != nil {
		return fmt.Errorf("scanning one: %w", err)
	}
	return nil
}

// Query executes query and scans all rows into dst
func (o *SqlQuery) Query(ctx context.Context, db Db, dst any) error {
	err := pgxscan.Select(ctx, db, dst, o.sql, o.params...)

	return err
}

// Exec executes query and returns result
func (o *SqlQuery) Exec(ctx context.Context, db Db) (int64, error) {
	
	r, err := db.Exec(ctx,o.sql, o.params...)

	if err != nil {
		return 0, err
	}

	return r.RowsAffected(), nil
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

// BeginTxFunc executes function in transaction
func BeginTxFunc(ctx context.Context, opts pgx.TxOptions, db *pgxpool.Pool, fn func(context.Context, Db) error) error {

	tx, err := db.BeginTx(ctx, opts)

	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()

	err = fn(ctx, tx)

	if err != nil {
		return err
	}

	return nil
}



// DbConfig is a configuration for mysql connection
type DbConfig struct {
	ConnectionString string `json:"connectionString" env:"CONNECTION_STRING"` // ConnectionString is a connection string for mysql
}

// Connect connects to mysql database
func Connect(cfg DbConfig) (*pgxpool.Pool, error) {
	ctx := context.Background()
	db, err := pgxpool.New(ctx,cfg.ConnectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
