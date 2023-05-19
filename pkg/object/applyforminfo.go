package object

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	hzhis "WowjoyProject/DataAcceptanceSystem/pkg/HZHIS"
	zlhis "WowjoyProject/DataAcceptanceSystem/pkg/ZLHIS"
)

// 获取申请单数据
func GetApplyFormData(object global.ApplyFormInfoData) (data []global.ApplyFormResultData) {
	// 1. 通过存储过程获取数据
	// 2. 通过接口获取数据
	// 3. 通过接口推送数据（创建数据库接受数据）
	global.Logger.Debug("开始获取申请单数据，请求参数object: ", object)

	switch global.ObjectSetting.InterfaceSystemType {
	case global.InterfaceSystem_SLHZYY:
		global.Logger.Debug("树兰医院")
		data = hzhis.ByHZHisViewGetApply(object)
	case global.InterfaceSystem_YYTJZMZZZXZYY:
		global.Logger.Debug("酉阳土家族苗族自治县中医院")
		data = zlhis.ByZLHisViewGetApply(object)
	default:
	}
	return
}
