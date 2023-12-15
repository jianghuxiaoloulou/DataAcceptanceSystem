package global

// 数据库全局变量

import "database/sql"

var (
	DBEngine *sql.DB // PACS集成平台数据库
)

func NullStringToString(str sql.NullString) string {
	if str.Valid {
		return str.String
	}
	return ""
}
