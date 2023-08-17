package object

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	rcqfby "WowjoyProject/DataAcceptanceSystem/pkg/RCQFBY"
	"strings"
)

// PACS信息推送
func ApplyFormStatusNotity(data global.ApplyFormStatusData) {
	// 具体业务
	global.Logger.Info("区域PACS请求数据：", data)
	// 获取申请单状态功能模块信息
	applyFuncCfg, err := model.GetHospApplyFuncCfg(data.HospitalID)
	if err != nil {
		global.Logger.Error("未查询到该医院的申请单功能配置，请在sys_hospital_apply_function表中配置相关信息 :", data.HospitalID)
		return
	}
	switch data.PARAM.ParamType {
	case global.Apply_Status_Canceled:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已取消")
		if applyFuncCfg.ApplyCanceled.String != "" {
			global.Logger.Debug("开始处理申请单已取消状态处理的功能")
			// 分隔符|*|
			applyCancel := strings.Split(applyFuncCfg.ApplyCanceled.String, "|*|")
			global.Logger.Debug("需要实现的功能：", applyCancel)
			for _, value := range applyCancel {
				switch value {
				case global.Apply_Func_A:
					global.Logger.Debug("实现功能A")
				case global.Apply_Func_B:
					global.Logger.Debug("实现功能B")
				case global.Apply_Func_C:
					global.Logger.Debug("实现功能C")
				case global.Apply_Func_D:
					global.Logger.Debug("实现功能D")
				case global.Apply_Func_E:
					global.Logger.Debug("实现功能E")
				default:
					global.Logger.Debug("未实现该功能")
				}
			}
		}
	case global.Apply_Status_Registered:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已报到")
		if applyFuncCfg.ApplyRegistered.String != "" {
			global.Logger.Debug("开始处理申请单已取消状态处理的功能")
			// 分隔符|*|
			applyRegis := strings.Split(applyFuncCfg.ApplyRegistered.String, "|*|")
			global.Logger.Debug("需要实现的功能：", applyRegis)
			for _, value := range applyRegis {
				switch value {
				case global.Apply_Func_A:
					global.Logger.Debug("实现功能A")
				case global.Apply_Func_B:
					global.Logger.Debug("实现功能B")
				case global.Apply_Func_C:
					global.Logger.Debug("实现功能C")
				case global.Apply_Func_D:
					global.Logger.Debug("实现功能D")
				case global.Apply_Func_E:
					global.Logger.Debug("实现功能E")
				default:
					global.Logger.Debug("未实现该功能")
				}
			}
		}
	case global.Apply_Status_Checked:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已检查")
		if applyFuncCfg.ApplyChecked.String != "" {
			global.Logger.Debug("开始处理申请单已取消状态处理的功能")
			// 分隔符|*|
			applyCheck := strings.Split(applyFuncCfg.ApplyChecked.String, "|*|")
			global.Logger.Debug("需要实现的功能：", applyCheck)
			for _, value := range applyCheck {
				switch value {
				case global.Apply_Func_A:
					global.Logger.Debug("实现功能A")
				case global.Apply_Func_B:
					global.Logger.Debug("实现功能B")
				case global.Apply_Func_C:
					global.Logger.Debug("实现功能C")
				case global.Apply_Func_D:
					global.Logger.Debug("实现功能D")
				case global.Apply_Func_E:
					global.Logger.Debug("实现功能E")
				default:
					global.Logger.Debug("未实现该功能")
				}
			}
		}
	case global.Apply_Status_Drafted:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已起草")
		if applyFuncCfg.ApplyDrafted.String != "" {
			global.Logger.Debug("开始处理申请单已取消状态处理的功能")
			// 分隔符|*|
			applyDraft := strings.Split(applyFuncCfg.ApplyDrafted.String, "|*|")
			global.Logger.Debug("需要实现的功能：", applyDraft)
			for _, value := range applyDraft {
				switch value {
				case global.Apply_Func_A:
					global.Logger.Debug("实现功能A")
				case global.Apply_Func_B:
					global.Logger.Debug("实现功能B")
				case global.Apply_Func_C:
					global.Logger.Debug("实现功能C")
				case global.Apply_Func_D:
					global.Logger.Debug("实现功能D")
				case global.Apply_Func_E:
					global.Logger.Debug("实现功能E")
				default:
					global.Logger.Debug("未实现该功能")
				}
			}
		}
	case global.Apply_Status_WaitAudit:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单 待审核")
		if applyFuncCfg.ApplyWaitaudit.String != "" {
			global.Logger.Debug("开始处理申请单已取消状态处理的功能")
			// 分隔符|*|
			applyWaitAudit := strings.Split(applyFuncCfg.ApplyWaitaudit.String, "|*|")
			global.Logger.Debug("需要实现的功能：", applyWaitAudit)
			for _, value := range applyWaitAudit {
				switch value {
				case global.Apply_Func_A:
					global.Logger.Debug("实现功能A")
				case global.Apply_Func_B:
					global.Logger.Debug("实现功能B")
				case global.Apply_Func_C:
					global.Logger.Debug("实现功能C")
				case global.Apply_Func_D:
					global.Logger.Debug("实现功能D")
				case global.Apply_Func_E:
					global.Logger.Debug("实现功能E")
				default:
					global.Logger.Debug("未实现该功能")
				}
			}
		}
	case global.Apply_Status_Audited:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已审核")
		if applyFuncCfg.ApplyAudited.String != "" {
			global.Logger.Debug("开始处理申请单已取消状态处理的功能")
			// 分隔符|*|
			applyAudit := strings.Split(applyFuncCfg.ApplyAudited.String, "|*|")
			global.Logger.Debug("需要实现的功能：", applyAudit)
			for _, value := range applyAudit {
				switch value {
				case global.Apply_Func_A:
					global.Logger.Debug("实现功能A")
				case global.Apply_Func_B:
					global.Logger.Debug("实现功能B")
				case global.Apply_Func_C:
					global.Logger.Debug("实现功能C")
				case global.Apply_Func_D:
					global.Logger.Debug("实现功能D")
				case global.Apply_Func_E:
					global.Logger.Debug("实现功能E")
				default:
					global.Logger.Debug("未实现该功能")
				}
			}
		}
	case global.Apply_Status_Charging:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单主动计费")
		if applyFuncCfg.ApplyCharging.String != "" {
			global.Logger.Debug("开始处理申请单已取消状态处理的功能")
			// 分隔符|*|
			applyCharg := strings.Split(applyFuncCfg.ApplyCharging.String, "|*|")
			global.Logger.Debug("需要实现的功能：", applyCharg)
			for _, value := range applyCharg {
				switch value {
				case global.Apply_Func_A:
					global.Logger.Debug("实现功能A")
				case global.Apply_Func_B:
					global.Logger.Debug("实现功能B")
				case global.Apply_Func_C:
					global.Logger.Debug("实现功能C")
				case global.Apply_Func_D:
					global.Logger.Debug("实现功能D")
				case global.Apply_Func_E:
					global.Logger.Debug("实现功能E")
				default:
					global.Logger.Debug("未实现该功能")
				}
			}
		}
	case global.Apply_Status_Diagnose:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单开始申请远程诊断")
		rcqfby.SendRemoteDiagnoseApplyData(data.HospitalID, data.PARAM.ParamValue)
	case global.Apply_Status_ViewRemote:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单开始申请远程查看")
		rcqfby.SendRemoteViewApplyData(data.HospitalID, data.PARAM.ParamValue)
	default:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单未知状态")
	}
}
