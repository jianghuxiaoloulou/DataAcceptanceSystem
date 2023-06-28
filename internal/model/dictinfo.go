package model

// 字典信息配置表

import "WowjoyProject/DataAcceptanceSystem/global"

// 获取字典NAME
// func GetDictName(code int) (string, error) {
// 	global.Logger.Info("获取字典NAME通过字典编码: ", code)
// 	sql := `SELECT dict_name FROM sys_dict_type WHERE dict_code = ?`
// 	var err error
// 	err = global.DBEngine.Ping()
// 	if err != nil {
// 		global.Logger.Error(err.Error())
// 		global.DBEngine.Close()
// 		global.DBEngine, _ = NewDBEngine(global.DatabaseSetting)
// 	}
// 	var name string
// 	row := global.DBEngine.QueryRow(sql, code)
// 	err = row.Scan(&name)
// 	if err != nil {
// 		global.Logger.Error(err)
// 		return name, err
// 	}
// 	return name, nil
// }

// // 获取字典CODE
// func GetDictCode(name string) (int, error) {
// 	global.Logger.Info("获取字典CODE通过字典名称: ", name)
// 	sql := `SELECT dict_code FROM sys_dict_type WHERE dict_name = ?`
// 	var err error
// 	err = global.DBEngine.Ping()
// 	if err != nil {
// 		global.Logger.Error(err.Error())
// 		global.DBEngine.Close()
// 		global.DBEngine, _ = NewDBEngine(global.DatabaseSetting)
// 	}
// 	var code int
// 	row := global.DBEngine.QueryRow(sql, name)
// 	err = row.Scan(&code)
// 	if err != nil {
// 		global.Logger.Error(err)
// 		return code, err
// 	}
// 	return code, nil
// }

func GetDictData() {
	global.Logger.Info("获取字典数据")
	sql := `SELECT dict_code,dict_name FROM sys_dict_type`
	var err error
	err = global.DBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.DBEngine.Close()
		global.DBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	rows, err := global.DBEngine.Query(sql)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		key := global.DictConfig{}
		rows.Scan(&key.DictCode, &key.DictName)
		data := global.DictDataObject{
			Code: int(key.DictCode.Int16),
			Name: key.DictName.String,
		}
		global.DictDatas = append(global.DictDatas, data)
	}
}
