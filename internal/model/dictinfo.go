package model

// 字典信息配置表

import "WowjoyProject/DataAcceptanceSystem/global"

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

// 获取字典NAME
func GetDictName(code int) (name string) {
	global.Logger.Info("获取字典NAME通过字典编码: ", code)
	for _, dict := range global.DictDatas {
		if code == dict.Code {
			name = dict.Name
			break
		}
	}
	return
}

// 获取字典CODE
func GetDictCode(name string) (code int) {
	global.Logger.Info("获取字典CODE通过字典名称: ", name)
	for _, dict := range global.DictDatas {
		if name == dict.Name {
			code = dict.Code
			break
		}
	}
	return
}
