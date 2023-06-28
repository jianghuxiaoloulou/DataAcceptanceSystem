package object

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	zlhis "WowjoyProject/DataAcceptanceSystem/pkg/ZLHIS"
)

// 获取申请单数据
func GetApplyFormData(object global.ApplyFormInfoData) (count int, data []global.ApplyFormResultData) {
	global.Logger.Debug("开始获取申请单数据，请求参数object: ", object)
	// 1.通过HospitalID 获取医院相关数据库连接信息
	hospitalConfig, err := model.GetHospitalConfig(object.HospitalID)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("第一步：获取的医院相关连接信息：", hospitalConfig)
	// 2. 判断对接HIS厂商
	switch int(hospitalConfig.HISType.Int16) {
	case global.HIS_Type_ZLHIS:
		// 开始对接中联HIS厂商
		global.Logger.Debug("开始对接中联HIS厂商")
		count, data = zlhis.GetApplyData(hospitalConfig, object)
	default:
		// 未配置的数据库连接
		global.Logger.Error("请配置sys_dict_hospital_config表中his_type字段的连接方式")
		return
	}
	return
}
