package object

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	rcqfby "WowjoyProject/DataAcceptanceSystem/pkg/RCQFBY"
)

// 获取第三方PACS提供的申请单数据和DICOM数据
func GetApplyAndDicomData(object global.ApplyDicomData) {
	global.Logger.Debug("开始获取申请单数据，请求参数object: ", object)
	// 1.通过HospitalID 获取医院相关数据库连接信息
	hospitalConfig, err := model.GetHospitalConfig(object.HospitalID)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("获取的医院相关连接信息：", hospitalConfig)

	// 获取申请单数据
	global.Logger.Debug("开始获取任城区妇保院申请单数据")
	rcqfby.GetApplyData(hospitalConfig, object)
	return
}

func GetApplyAndDicomDataTime(object global.ApplyDicomData) {
	global.Logger.Debug("开始获取申请单数据，请求参数object: ", object)
	// 1.通过HospitalID 获取医院相关数据库连接信息
	hospitalConfig, err := model.GetHospitalConfig(object.HospitalID)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("获取的医院相关连接信息：", hospitalConfig)

	// 获取申请单数据
	global.Logger.Debug("开始获取任城区妇保院申请单数据")
	rcqfby.GetApplyDataTime(hospitalConfig, object)
	return
}
