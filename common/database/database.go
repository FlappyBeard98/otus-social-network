package database

import (
	"context"
	"database/sql"
	"github.com/georgysavva/scany/v2/sqlscan"
	_ "github.com/go-sql-driver/mysql"
	"social-network/common"
	"social-network/common/application"
)

type SqlQuery interface {
	Sql() string
}

type SqlParametersProvider interface {
	GetParams() []any
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

func NilOrExprMysql(obj any,expr string) string{
	if obj == nil {
		return "true"
	} else {
		return expr
	}
}

func FixTx(tx *sql.Tx,err *error)  {
	if *err != nil {
		_ = tx.Rollback()
	} else {
		_ = tx.Commit()
	}
}

type DbHandler[In SqlQuery,Out any] application.Handler[In,[]Out]

type dbHandler[In SqlQuery,Out any] struct {
	sqlscan.Querier
}

func (receiver *dbHandler[In,Out]) Handle(ctx context.Context, arg In) (result []Out,err error) {

	var args []any
	var a any = arg
	if sqlParametersProvider, ok := a.(SqlParametersProvider); ok {
		args = sqlParametersProvider.GetParams()
	} else {
		args = common.GetFieldsValuesAsSlice(arg)
	}

	statement := arg.Sql()
	err = sqlscan.Select(ctx, receiver, &result, statement, args...)

	return
}

func NewDbHandler[In SqlQuery,Out any](connection sqlscan.Querier) DbHandler[In,Out]{
	return &dbHandler[In,Out]{connection}
}
