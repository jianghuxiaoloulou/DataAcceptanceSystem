package model

//HIS厂商信息配置表

import "WowjoyProject/DataAcceptanceSystem/global"

// 获取HIS厂商相关信息
func GetHisConfig(histype int) (global.HisConfig, error) {
	global.Logger.Info("开始查询对应HIS厂商相关信息: ", histype)
	sql := `SELECT his_type,his_type_name,apply_mz_db_type,apply_mz_db_conn,apply_mz_view_name,apply_zy_db_type,apply_zy_db_conn,
	apply_zy_view_name,apply_tj_db_type,apply_tj_db_conn,apply_tj_view_name,his_interface_url FROM sys_dict_his_config   
	WHERE his_type = ?`
	var err error
	err = global.DBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.DBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}

	row := global.DBEngine.QueryRow(sql, histype)
	hisdata := global.HisConfig{}
	err = row.Scan(&hisdata.HISType, &hisdata.HISTypeName, &hisdata.ApplyMZDBType, &hisdata.ApplyMZDBConn, &hisdata.ApplyMZViewName,
		&hisdata.ApplyZYDBType, &hisdata.ApplyZYDBConn, &hisdata.ApplyZYViewName, &hisdata.ApplyTJDBType, &hisdata.ApplyTJDBConn, &hisdata.ApplyTJViewName,
		&hisdata.HISInterfaceURL)
	if err != nil {
		global.Logger.Error(err)
		return hisdata, err
	}
	return hisdata, nil
}
