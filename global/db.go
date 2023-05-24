package global

// 数据库全局变量

import "database/sql"

// 服务编码
const (
	HisMysql  int = 1 // HIS数据库mysql
	HisOracle int = 2 // HIS数据库Oracle
)

var (
	ReadDBEngine   *sql.DB
	WriteDBEngine  *sql.DB
	OracleDBEngine *sql.DB
)
