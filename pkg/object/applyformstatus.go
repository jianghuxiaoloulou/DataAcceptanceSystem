package object

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	rcqfby "WowjoyProject/DataAcceptanceSystem/pkg/RCQFBY"
	"WowjoyProject/DataAcceptanceSystem/pkg/wdhis"
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
				WorkListFunc(value, data.PARAM.ParamValue)
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
				WorkListFunc(value, data.PARAM.ParamValue)
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
				WorkListFunc(value, data.PARAM.ParamValue)
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
				WorkListFunc(value, data.PARAM.ParamValue)
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
				WorkListFunc(value, data.PARAM.ParamValue)
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
				WorkListFunc(value, data.PARAM.ParamValue)
			}
		}
	case global.Apply_Status_Diagnose:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单开始申请远程诊断")
		if applyFuncCfg.ApplyDiagnose.String != "" {
			global.Logger.Debug("开始处理申请单远程诊断功能....")
			// 分割符|*|
			applyDiag := strings.Split(applyFuncCfg.ApplyDiagnose.String, "|*|")
			global.Logger.Debug("需要实现的功能：", applyDiag)
			for _, value := range applyDiag {
				WorkListFunc(value, data.PARAM.ParamValue)
			}
		}
	case global.Apply_Status_ViewRemote:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单开始申请远程查看")
		if applyFuncCfg.ApplyViewRemote.String != "" {
			global.Logger.Debug("开始处理申请单远程查看功能....")
			// 分割符|*|
			applyViewRemote := strings.Split(applyFuncCfg.ApplyDiagnose.String, "|*|")
			global.Logger.Debug("需要实现的功能：", applyViewRemote)
			for _, value := range applyViewRemote {
				WorkListFunc(value, data.PARAM.ParamValue)
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
				WorkListFunc(value, data.PARAM.ParamValue)
			}
		}
	case global.Apply_Status_AuditeRemote:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单申请远程审核")
		if applyFuncCfg.ApplyAuditeRemote.String != "" {
			global.Logger.Debug("开始处理申请单申请远程审核的功能")
			// 分隔符|*|
			applyCharg := strings.Split(applyFuncCfg.ApplyCharging.String, "|*|")
			global.Logger.Debug("需要实现的功能：", applyCharg)
			for _, value := range applyCharg {
				WorkListFunc(value, data.PARAM.ParamValue)
			}
		}
	default:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单未知状态")
	}
}

// 实现功能清单列表
func WorkListFunc(key, value string) {
	switch key {
	case global.Apply_Func_A:
		global.Logger.Debug("实现功能A：万达区域HIS检查登记回写EX1001")
		wdhis.FuncEX1001(value)
	case global.Apply_Func_B:
		global.Logger.Debug("实现功能B：万达区域HIS检查回写(EX1002)")
		wdhis.FuncEX1002(value)
	case global.Apply_Func_C:
		global.Logger.Debug("实现功能C：万达区域HIS检查报告回写（EX1003）")
		wdhis.FuncEX1003(value)
	case global.Apply_Func_D:
		global.Logger.Debug("实现功能D: 任城区妇保院报告回写")
		rcqfby.ReportDataWriteBack(value)
	case global.Apply_Func_E:
		global.Logger.Debug("实现功能E: 济宁附属医院远程诊断")
		rcqfby.SendRemoteDiagnoseApplyData(value)
	case global.Apply_Func_F:
		global.Logger.Debug("实现功能F: 济宁附属医院远程查看")
		rcqfby.SendRemoteViewApplyData(value)
	case global.Apply_Func_G:
		global.Logger.Debug("实现功能G: 济宁附属医院远程审核")
		rcqfby.SendRemoteAuditeApplyData(value)
	case global.Apply_Func_H:
		global.Logger.Debug("实现功能H: 万达区域HIS检查报告回写（EX1004）")
		wdhis.FuncEX1004(value)
	case global.Apply_Func_I:
		global.Logger.Debug("实现功能I: 万达区域HIS申请单撤销（EX1006）")
		wdhis.FuncEX1006(value)
	case global.Apply_Func_J:
		global.Logger.Debug("实现功能J: 万达区域HIS危急消息通知（EX2000）")
		// wdhis.FuncEX2000(value)
	default:
		global.Logger.Debug("未实现该功能")
	}
}
