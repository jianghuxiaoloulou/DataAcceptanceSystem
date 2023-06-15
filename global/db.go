package global

// 数据库全局变量

import "database/sql"

// 服务编码
const (
	HisMysql  int = 1 // HIS数据库mysql
	HisOracle int = 2 // HIS数据库Oracle
)

var (
	PacsDBEngine    *sql.DB // PACS数据库
	MZApplyDBEngine *sql.DB // 门诊申请单数据库
	ZYApplyDBEngine *sql.DB // 住院申请单数据库
	TJApplyDBEngine *sql.DB // 住院申请单数据库
)
