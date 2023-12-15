package model

// 字典信息配置表

import "WowjoyProject/DataAcceptanceSystem/global"

func GetDictData() {
	global.Logger.Info("获取字典数据")
	sql := `SELECT dict_type,dict_code,dict_name,dict_his_code,dict_his_name FROM sys_dict_type`
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
		rows.Scan(&key.DictType, &key.DictCode, &key.DictName, &key.DictHisCode, &key.DictHisName)
		data := global.DictDataObject{
			Type:    int(key.DictType.Int16),
			Code:    int(key.DictCode.Int16),
			Name:    key.DictName.String,
			HisCode: key.DictHisCode.String,
			HisName: key.DictHisName.String,
		}
		global.DictDatas = append(global.DictDatas, data)
	}
	global.Logger.Debug(global.DictDatas)
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

// 通过字典Code获取HIS对应的Code
func ByCodeGetHisCode(code int) (hiscode string) {
	global.Logger.Info("获取字典HisCode通过字典编码: ", code)
	for _, dict := range global.DictDatas {
		if code == dict.Code {
			hiscode = dict.HisCode
			break
		}
	}
	return
}

// 通过HisCode获取字段Code
func ByHisCodeGetCode(hiscode string) (code int) {
	global.Logger.Info("获取字典Code通过HisCode字典编码: ", hiscode)
	for _, dict := range global.DictDatas {
		if hiscode == dict.HisCode {
			code = dict.Code
			break
		}
	}
	return
}
