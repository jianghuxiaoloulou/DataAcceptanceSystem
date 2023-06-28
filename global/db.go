package global

// 数据库全局变量

import "database/sql"

var (
	DBEngine *sql.DB // PACS集成平台数据库
)
