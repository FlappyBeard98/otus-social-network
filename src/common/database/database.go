package database

import (
	"context"
	"database/sql"
	"github.com/georgysavva/scany/v2/sqlscan"
	_ "github.com/go-sql-driver/mysql"
	"social-network/common"
)

type SqlQuery interface {
	Sql() string
}

func InitMysql(databaseConnection string) *sql.DB {
	db, err := sql.Open("mysql", databaseConnection)
	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	return db
}

func ExecuteSqlQuery[Result any](ctx context.Context, querier sqlscan.Querier, q SqlQuery) (result Result, err error) {
	args := common.GetFieldsValuesAsSlice(q)
	err = sqlscan.Select(ctx, querier, &result, q.Sql(), args...)

	return
}
