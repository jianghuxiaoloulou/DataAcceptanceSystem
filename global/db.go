package global

// 数据库全局变量

import "database/sql"

var (
	ReadDBEngine   *sql.DB
	WriteDBEngine  *sql.DB
	OracleDBEngine *sql.DB
)
