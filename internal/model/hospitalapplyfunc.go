package model

import "WowjoyProject/DataAcceptanceSystem/global"

// 获取医院申请单功能模块配置
func GetHospApplyFuncCfg(hospitalId string) (global.ApplyFuncConfig, error) {
	global.Logger.Info("医院申请单状态功能模块配置相关信息: ", hospitalId)
	sql := `SELECT hospital_id,his_type,apply_canceled,apply_registered,apply_checked,apply_drafted,apply_waitaudit,
	apply_audited,apply_diagnose,apply_viewremote,apply_auditeremote,apply_charging 
	FROM sys_hospital_apply_function 
	WHERE hospital_id = ?`
	var err error
	err = global.DBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.DBEngine.Close()
		global.DBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	row := global.DBEngine.QueryRow(sql, hospitalId)
	applyfunc := global.ApplyFuncConfig{}
	err = row.Scan(&applyfunc.HospitalId, &applyfunc.HISType, &applyfunc.ApplyCanceled, &applyfunc.ApplyRegistered, &applyfunc.ApplyChecked,
		&applyfunc.ApplyDrafted, &applyfunc.ApplyWaitaudit, &applyfunc.ApplyAudited, &applyfunc.ApplyDiagnose, &applyfunc.ApplyViewRemote,
		&applyfunc.ApplyAuditeRemote, &applyfunc.ApplyCharging)
	if err != nil {
		global.Logger.Error(err)
		return applyfunc, err
	}
	return applyfunc, nil
}
