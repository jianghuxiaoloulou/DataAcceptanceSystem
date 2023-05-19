package object

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	zlhis "WowjoyProject/DataAcceptanceSystem/pkg/ZLHIS"
)

// PACS信息推送
func ApplyFormStatusNotity(data global.ApplyFormStatusData) {
	// 具体业务
	global.Logger.Info("PACS推送的数据：", data)
	switch global.ObjectSetting.InterfaceSystemType {
	case global.InterfaceSystem_SLHZYY:
		switch data.PARAM.ParamType {
		case global.Apply_Status_Canceled:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已取消")
			// 具体业务员
		case global.Apply_Status_Registered:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已报到")
			// 具体业务员
		case global.Apply_Status_Checked:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已检查")
			// 具体业务员
		case global.Apply_Status_Drafted:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已起草")
			// 具体业务员
		case global.Apply_Status_WaitAudit:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 待审核")
			// 具体业务员
		case global.Apply_Status_Audited:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已审核")
			// 具体业务员
		case global.Apply_Status_Other:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 其它操作")
			// 具体业务员
		default:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单未知状态")
			// 具体业务员
		}
	case global.InterfaceSystem_YYTJZMZZZXZYY:
		switch data.PARAM.ParamType {
		case global.Apply_Status_Canceled:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已取消")
			// 具体业务员
			zlhis.CanceledWriteBack(data)
		case global.Apply_Status_Registered:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已报到")
			// 具体业务员
			zlhis.RegisteredWriteBack(data)
		case global.Apply_Status_Checked:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已检查")
			// 具体业务员
		case global.Apply_Status_Drafted:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已起草")
			// 具体业务员
		case global.Apply_Status_WaitAudit:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 待审核")
			// 具体业务员
		case global.Apply_Status_Audited:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已审核")
			// 具体业务员
			zlhis.AuditedWriteBack(data)
		case global.Apply_Status_Other:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单 其它操作")
			// 具体业务员
		default:
			global.Logger.Debug(data.PARAM.ParamValue, " 申请单未知状态")
			// 具体业务员
		}
	default:
	}
}
