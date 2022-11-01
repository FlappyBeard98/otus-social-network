package database

import (
	"context"
	"database/sql"
	"github.com/georgysavva/scany/v2/sqlscan"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"social-network/common"
	"social-network/common/application"
)

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

// InitMysql initialize connection to mysql database
func InitMysql(databaseConnection string) *sql.DB {

	var db *sql.DB

	db, err := sql.Open("mysql", databaseConnection)
	if err != nil {
		panic(err.Error())
	}

	return db
}

// Migrate apply migration script from file
func Migrate(db *sql.DB, migrationFile string) error {
	bytes, err := os.ReadFile(migrationFile)
	sqlStmt := string(bytes)
	if err != nil {
		return err
	}

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer FixTx(tx, &err)

	_, err = tx.Exec(sqlStmt)

	return err
}

// FixTx commit or rollback transaction depending on the presence of an error
func FixTx(tx *sql.Tx, err *error) {
	if *err != nil {
		_ = tx.Rollback()
	} else {
		_ = tx.Commit()
	}
}

// SqlQuery interface mark structs as query to execute in the DbHandler
type SqlQuery interface {
	//Sql returns sql statement
	Sql() string
}

// SqlParametersProvider should be used for queries with duplicate parameters
type SqlParametersProvider interface {
	//GetParams returns params slice for query execution
	GetParams() []any
}


// DbHandler is common handler for structs that implements SqlQuery
type DbHandler[In SqlQuery, Out any] application.Handler[In, []Out]

// dbHandler implements DbHandler
type dbQuerierHandler[In SqlQuery, Out any] struct {
	sqlscan.Querier
}

// Handle implements Handle method of application.Handler interface for generic sql-query execution
func (receiver *dbQuerierHandler[In, Out]) Handle(ctx context.Context, arg In) (result []Out, err error) {

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

func NewDbQuerierHandler[In SqlQuery, Out any](connection sqlscan.Querier) DbHandler[In, Out] {
	return &dbQuerierHandler[In, Out]{connection}
}

// dbHandler implements DbHandler
type dbHandler[In SqlQuery, Out any] struct {
	*sql.DB
}

// Handle implements Handle method of application.Handler interface for generic sql-query execution
func (receiver *dbHandler[In, Out]) Handle(ctx context.Context, arg In) (result []Out, err error) {

	var args []any
	var a any = arg
	if sqlParametersProvider, ok := a.(SqlParametersProvider); ok {
		args = sqlParametersProvider.GetParams()
	} else {
		args = common.GetFieldsValuesAsSlice(arg)
	}

	statement := arg.Sql()

	_, err = receiver.Exec(statement,args...)
	
	return
}

func NewDbHandler[In SqlQuery](connection *sql.DB) DbHandler[In, any] {
	return &dbHandler[In, any]{connection}
}