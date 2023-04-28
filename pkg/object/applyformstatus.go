package object

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ZLHisApplyInfo struct {
	ApplyId     string `json:"apply_id"`
	ApplyStatus string `json:"apply_status"`
}

type ZLHisApplyItem struct {
	ChkinPs string `json:"chkin_ps"`
}

type RegisteredRequest struct {
	INPUT RegisteredRequestInput `json:"input"`
}

type RegisteredRequestInput struct {
	HEAD      ZLHisRequestHead `json:"head"`
	ApplyInfo ZLHisApplyInfo   `json:"apply_info"`
	ApplyItem ZLHisApplyItem   `json:"apply_item"`
}

type ZLHisReport struct {
	OrderCtl string `json:"order_ctl"`
	ApplyId  string `json:"apply_id"`
	OrderId  string `json:"order_id"`
}

type ZLHisRptInfo struct {
	IoitemCname    string `json:"ioitem_cname"`
	OrderRptResult string `json:"order_rpt_result"`
	Chkr           string `json:"chkr"`
	ChkTime        string `json:"chk_time"`
}

type RptRequestInput struct {
	HEAD      ZLHisRequestHead `json:"head"`
	ApplyInfo ZLHisReport      `json:"apply_info"`
	RptInfo   []ZLHisRptInfo   `json:"rpt_info"`
}

type RepRequest struct {
	INPUT RptRequestInput `json:"input"`
}

// PACS信息推送
func ApplyFormStatusNotity(data global.ApplyFormStatusData) {
	// 具体业务
	global.Logger.Info("PACS推送的数据：", data)
	switch data.PARAM.ParamType {
	case global.Apply_Status_Canceled:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已取消")
		// 具体业务员
		CanceledWriteBack(data)
	case global.Apply_Status_Registered:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单 已报到")
		// 具体业务员
		RegisteredWriteBack(data)
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
		AuditedWriteBack(data)
	case global.Apply_Status_Other:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单 其它操作")
		// 具体业务员
	default:
		global.Logger.Debug(data.PARAM.ParamValue, " 申请单未知状态")
		// 具体业务员
	}
}

// 接收检查申请，PACS检查报到后回写ZLHIS(5004)[已检查]
func RegisteredWriteBack(data global.ApplyFormStatusData) {
	reqinfo1 := ZLHisApplyInfo{
		ApplyId:     data.PARAM.ParamValue,
		ApplyStatus: "3",
	}
	reqinfo2 := ZLHisApplyItem{
		ChkinPs: "",
	}
	requestHead := ZLHisRequestHead{
		Bizno:    "5004",
		Sysno:    "PACS",
		Tarno:    "HIS",
		Time:     time.Now().Format("20060102150405"),
		ActionNo: "",
	}
	requestInout := RegisteredRequestInput{
		HEAD:      requestHead,
		ApplyInfo: reqinfo1,
		ApplyItem: reqinfo2,
	}
	request := RegisteredRequest{
		INPUT: requestInout,
	}
	global.Logger.Debug("回写中联请求的数据：", request)
	requestData, err := json.Marshal(request)
	url := global.ObjectSetting.HISURL
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if err != nil {
		global.Logger.Debug("http NewRequest err ", err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		global.Logger.Debug("http Do err ", err.Error())
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		global.Logger.Debug("ioutil.ReadAll err ", err.Error())
		return
	}
	global.Logger.Debug(data.PARAM.ParamValue, " :回写中联单返回结果：", string(body))
}

// 取消检查接收，PACS检查取消报到后回写ZLHIS(5005)[取消检查]
func CanceledWriteBack(data global.ApplyFormStatusData) {
	reqinfo1 := ZLHisApplyInfo{
		ApplyId:     data.PARAM.ParamValue,
		ApplyStatus: "2",
	}
	reqinfo2 := ZLHisApplyItem{
		ChkinPs: "",
	}
	requestHead := ZLHisRequestHead{
		Bizno:    "5005",
		Sysno:    "PACS",
		Tarno:    "HIS",
		Time:     time.Now().Format("20060102150405"),
		ActionNo: "",
	}
	requestInout := RegisteredRequestInput{
		HEAD:      requestHead,
		ApplyInfo: reqinfo1,
		ApplyItem: reqinfo2,
	}
	request := RegisteredRequest{
		INPUT: requestInout,
	}
	global.Logger.Debug("回写中联请求的数据：", request)
	requestData, err := json.Marshal(request)
	url := global.ObjectSetting.HISURL
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if err != nil {
		global.Logger.Debug("http NewRequest err ", err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		global.Logger.Debug("http Do err ", err.Error())
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		global.Logger.Debug("ioutil.ReadAll err ", err.Error())
		return
	}
	global.Logger.Debug(data.PARAM.ParamValue, " :回写中联单返回结果：", string(body))
}

// 发送检查报告，PACS发送报告后回写ZLHIS(5008)[回写报告内容]
func AuditedWriteBack(data global.ApplyFormStatusData) {
	// 通过申请单ID 获取报告信息
	repData, _ := model.GetReportInfo(data.PARAM.ParamValue)
	var reqinfoArr []ZLHisRptInfo
	audittime := strings.ReplaceAll(repData.AuditTime, "-", "")
	audittime = strings.ReplaceAll(repData.AuditTime, " ", "")
	audittime = strings.ReplaceAll(repData.AuditTime, ":", "")
	reqinfo1 := ZLHisRptInfo{
		IoitemCname:    "检查所见",
		OrderRptResult: repData.Finding,
		Chkr:           repData.AuditDoctor,
		ChkTime:        audittime,
	}
	reqinfoArr = append(reqinfoArr, reqinfo1)
	reqinfo2 := ZLHisRptInfo{
		IoitemCname:    "诊断意见",
		OrderRptResult: repData.Conclusion,
		Chkr:           repData.AuditDoctor,
		ChkTime:        audittime,
	}
	reqinfoArr = append(reqinfoArr, reqinfo2)

	applyinfo := ZLHisReport{
		OrderCtl: "CN",
		ApplyId:  data.PARAM.ParamValue,
		OrderId:  "",
	}

	requestHead := ZLHisRequestHead{
		Bizno:    "5008",
		Sysno:    "PACS",
		Tarno:    "HIS",
		Time:     time.Now().Format("20060102150405"),
		ActionNo: "",
	}
	requestInout := RptRequestInput{
		HEAD:      requestHead,
		ApplyInfo: applyinfo,
		RptInfo:   reqinfoArr,
	}
	request := RepRequest{
		INPUT: requestInout,
	}
	global.Logger.Debug("回写中联请求的数据：", request)
	requestData, err := json.Marshal(request)
	url := global.ObjectSetting.HISURL
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if err != nil {
		global.Logger.Debug("http NewRequest err ", err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		global.Logger.Debug("http Do err ", err.Error())
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		global.Logger.Debug("ioutil.ReadAll err ", err.Error())
		return
	}
	global.Logger.Debug(data.PARAM.ParamValue, " :回写中联单返回结果：", string(body))
}
