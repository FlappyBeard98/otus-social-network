package service

type SqlQuery interface {
	Sql() string
	Params() []interface{}
}

type Cached struct {
	SqlQuery
	Key string
}

type Invalidate struct {
	SqlQuery
	Key string
}
