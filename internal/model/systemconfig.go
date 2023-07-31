package model

import "WowjoyProject/DataAcceptanceSystem/global"

// 系统配置
func GetSystemData() {
	global.Logger.Info("获取系统数据")
	sql := `SELECT qy_pacs_type,qy_pacs_conn, qy_pacs_interface_url FROM sys_config`
	err := global.DBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.DBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}

	row := global.DBEngine.QueryRow(sql)
	key := global.SystemCfg{}
	row.Scan(&key.QYPacsType, &key.QYPacsConn, &key.QYPacsInterfaceUrl)
	data := global.SystemDataObject{
		QYPacsType:         key.QYPacsType.String,
		QYPacsConn:         key.QYPacsConn.String,
		QYPacsInterfaceUrl: key.QYPacsInterfaceUrl.String,
	}
	global.SystemData = data
}
