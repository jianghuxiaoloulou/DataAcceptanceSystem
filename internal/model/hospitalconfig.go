package model

// 医院配置信息表

import "WowjoyProject/DataAcceptanceSystem/global"

// 获取医院相关信息
func GetHospitalConfig(hospitalID string) (global.HospitalConfig, error) {
	global.Logger.Info("开始查询对接医院相关信息: ", hospitalID)
	sql := `SELECT hospital_id,hospital_name,pacs_db_type,pacs_db_conn,pacs_interface_url,his_type 
	FROM sys_dict_hospital_config WHERE hospital_id = ?`
	var err error
	err = global.DBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.DBEngine.Close()
		global.DBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}

	row := global.DBEngine.QueryRow(sql, hospitalID)
	hosdata := global.HospitalConfig{}
	err = row.Scan(&hosdata.HospitalId, &hosdata.HospitalName, &hosdata.PacsDBType, &hosdata.PacsDBConn, &hosdata.PacsInterfaceURL, &hosdata.HISType)
	if err != nil {
		global.Logger.Error(err)
		return hosdata, err
	}
	return hosdata, nil
}
