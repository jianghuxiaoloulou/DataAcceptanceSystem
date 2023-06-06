package zlhis

//中联his包
import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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
	ApplyInfo []ZLHisApplyInfo `json:"apply_info"`
	ApplyItem []ZLHisApplyItem `json:"apply_item"`
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
	ApplyInfo []ZLHisReport    `json:"apply_info"`
	RptInfo   []ZLHisRptInfo   `json:"rpt_info"`
}

type RepRequest struct {
	INPUT RptRequestInput `json:"input"`
}

type ZLHisReqInfo struct {
	QueryKey     string `json:"query_key"`
	QueryContent string `json:"query_content"`
	Pvid         string `json:"pvid"`
}

type ZLHisReqInfo2 struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ZLHisRequest struct {
	INPUT ZLHisRequestInput `json:"input"`
}

type ZLHisRequestInput struct {
	HEAD     ZLHisRequestHead `json:"head"`
	ReqInfo  []ZLHisReqInfo   `json:"req_info"`
	ReqInfo2 []ZLHisReqInfo2  `json:"req_info2"`
}

type ZLHisRequestHead struct {
	Bizno    string `json:"bizno"`
	Sysno    string `json:"sysno"`
	Tarno    string `json:"tarno"`
	Time     string `json:"time"`
	ActionNo string `json:"action_no"`
}

// 接收检查申请，PACS检查报到后回写ZLHIS(5004)[已检查]
func RegisteredWriteBack(data global.ApplyFormStatusData) {
	reqinfo1 := ZLHisApplyInfo{
		ApplyId:     data.PARAM.ParamValue,
		ApplyStatus: "3",
	}
	var reqinfo1Arr []ZLHisApplyInfo
	reqinfo1Arr = append(reqinfo1Arr, reqinfo1)

	reqinfo2 := ZLHisApplyItem{
		ChkinPs: "",
	}
	var reqinfo2Arr []ZLHisApplyItem
	reqinfo2Arr = append(reqinfo2Arr, reqinfo2)

	requestHead := ZLHisRequestHead{
		Bizno:    "5004",
		Sysno:    "PACS",
		Tarno:    "HIS",
		Time:     time.Now().Format("20060102150405"),
		ActionNo: "",
	}
	requestInout := RegisteredRequestInput{
		HEAD:      requestHead,
		ApplyInfo: reqinfo1Arr,
		ApplyItem: reqinfo2Arr,
	}
	request := RegisteredRequest{
		INPUT: requestInout,
	}
	requestData, err := json.Marshal(request)
	global.Logger.Debug("回写中联请求的数据：", requestData)

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
	var reqinfo1Arr []ZLHisApplyInfo
	reqinfo1Arr = append(reqinfo1Arr, reqinfo1)

	reqinfo2 := ZLHisApplyItem{
		ChkinPs: "",
	}

	var reqinfo2Arr []ZLHisApplyItem
	reqinfo2Arr = append(reqinfo2Arr, reqinfo2)

	requestHead := ZLHisRequestHead{
		Bizno:    "5005",
		Sysno:    "PACS",
		Tarno:    "HIS",
		Time:     time.Now().Format("20060102150405"),
		ActionNo: "",
	}
	requestInout := RegisteredRequestInput{
		HEAD:      requestHead,
		ApplyInfo: reqinfo1Arr,
		ApplyItem: reqinfo2Arr,
	}
	request := RegisteredRequest{
		INPUT: requestInout,
	}
	requestData, err := json.Marshal(request)
	global.Logger.Debug("回写中联请求的数据：", requestData)

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
	var applyinfoArr []ZLHisReport
	applyinfoArr = append(applyinfoArr, applyinfo)

	requestHead := ZLHisRequestHead{
		Bizno:    "5008",
		Sysno:    "PACS",
		Tarno:    "HIS",
		Time:     time.Now().Format("20060102150405"),
		ActionNo: "",
	}
	requestInout := RptRequestInput{
		HEAD:      requestHead,
		ApplyInfo: applyinfoArr,
		RptInfo:   reqinfoArr,
	}
	request := RepRequest{
		INPUT: requestInout,
	}
	requestData, err := json.Marshal(request)
	global.Logger.Debug("回写中联请求的数据：", requestData)

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

// 通过数据库视图获取申请单数据（Oracle数据库）
func ByZLHisViewGetApply(object global.ApplyFormInfoData) (count int, data []global.ApplyFormResultData) {
	global.Logger.Debug("开始通过中联视图获取数据：")
	// 查询视图数据
	param1len := len(object.PARAM)
	param2len := len(object.PARAM2)

	var patType []string
	var stuType []string

	for _, value := range object.PatientType {
		switch value {
		case global.Pat_Type_MZ:
			patType = append(patType, "门诊")
		case global.Pat_Type_ZY:
			patType = append(patType, "住院")
		case global.Pat_Type_TJ:
			patType = append(patType, "体检")
		default:
			patType = append(patType, "其他")
		}
	}

	for _, stu := range object.StudyType {
		switch stu {
		case global.Study_Type_XRay:
			stuType = append(stuType, "X-Ray")
		case global.Study_Type_DR:
			stuType = append(stuType, "DR")
		case global.Study_Type_CT:
			stuType = append(stuType, "CT")
		case global.Study_Type_MR:
			stuType = append(stuType, "MR")
		case global.Study_Type_DSA:
			stuType = append(stuType, "DSA")
		case global.Study_Type_US:
			stuType = append(stuType, "US")
		case global.Study_Type_ES:
			stuType = append(stuType, "ES")
		case global.Study_Type_PA:
			stuType = append(stuType, "PA")
		case global.Study_Type_NM:
			stuType = append(stuType, "NM")
		case global.Study_Type_PET:
			stuType = append(stuType, "PET")
		default:
			stuType = append(stuType, "OT")
		}
	}

	sql := `select "his_request_id","patient_name","patient_type_code","patient_type_name","medical_record_number",
	"sex_code","sex_name",regexp_substr("age",'[0-9]+') age,replace("age", regexp_substr("age",'[0-9]+'), '') age_unit,
	"birthday","modality_code","project_code","project_name","bodypart_code","bodypart",
	"outpatient_number","inhospital_number","visit_card_number","phone_number","inp_ward_id","patient_section_name",
	"sickbed_number","request_time","his_request_detail_id","id_card_number","address","clinical_diagnosis","medical_history",
	"request_department_code","request_department_name","request_doctor_code","request_doctor_name","check_note","film_count",
	"film_type","graphic_report","emergency","fee" 
	from V_PACS_HZ where 1 = 1`
	// 参数1
	var param1str string
	for i := 0; i < param1len; i++ {
		if i > 0 {
			param1str += " or "
		}
		switch object.PARAM[i].ParamType {
		case global.Apply_Param_JZKH:
			param1str += "\"visit_card_number\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_MZH:
			param1str += "\"outpatient_number\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_ZYH:
			param1str += "\"inhospital_number\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_BLH:
			param1str += "\"medical_record_number\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_TJH:
			param1str += "\"outpatient_number\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_MZSQDH:
			param1str += "\"his_request_id\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_ZYSQDH:
			param1str += "\"his_request_id\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_TJSQDH:
			param1str += "\"his_request_id\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_SFZH:
			param1str += "\"id_card_number\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_XM:
			param1str += "\"patient_name\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_JZ:
			param1str += "\"emergency\" = '" + object.PARAM[i].ParamValue + "'"
		default:
			param1str += "1 = 1"
		}
	}
	if param1len > 0 {
		sql += " and("
		sql += param1str
		sql += ")"
	}

	// 参数就诊类型
	var parampatType string
	for i := 0; i < len(patType); i++ {
		if i > 0 {
			parampatType += ","
		}
		parampatType += "'" + patType[i] + "'"
	}
	if len(patType) > 0 {
		sql += " and("
		sql += "\"patient_type_name\" in ("
		sql += parampatType
		sql += "))"
	}

	// 参数检查类型
	var paramstuType string
	for i := 0; i < len(stuType); i++ {
		if i > 0 {
			paramstuType += ","
		}
		paramstuType += "'" + stuType[i] + "'"
	}
	if len(stuType) > 0 {
		sql += " and("
		sql += "\"modality_code\" in ("
		sql += paramstuType
		sql += "))"
	}

	// 参数2
	var param2str string
	for i := 0; i < param2len; i++ {
		if i > 0 {
			param2str += " or "
		}
		if (object.PARAM2[i].StartDate != "") && (object.PARAM2[i].EndDate != "") {
			param2str += " \"request_time\" between to_date('" + object.PARAM2[i].StartDate + "','yyyy-mm-dd hh24:mi:ss') and to_date('" + object.PARAM2[i].EndDate + "','yyyy-mm-dd hh24:mi:ss')"
		} else {
			param2str += "1=1"
		}
	}
	if param2len > 0 {
		sql += " and ("
		sql += param2str
		sql += ")"
	}

	// 排序
	// 单独返回条件数据总数(不仅仅是分页后的数据)
	count = model.GetDataCount(sql)

	// 分页
	if object.StartSize >= 0 && object.EndSize > 0 {
		sql += " and rownum between "
		sql += strconv.Itoa(object.StartSize) + " and " + strconv.Itoa(object.EndSize)
	}

	global.Logger.Debug("执行的sql语句是: ", sql)
	data = model.GetZLHisViewApply(sql)
	return
}

// 通过MySql数据库连接获取数据
func ByZLHisMysqlView(object global.ApplyFormInfoData) (count int, data []global.ApplyFormResultData) {
	global.Logger.Debug("开始通过中联视图获取数据：")
	// 查询视图数据
	param1len := len(object.PARAM)
	param2len := len(object.PARAM2)

	var patType []string
	var stuType []string

	for _, value := range object.PatientType {
		switch value {
		case global.Pat_Type_MZ:
			patType = append(patType, "门诊")
		case global.Pat_Type_ZY:
			patType = append(patType, "住院")
		case global.Pat_Type_TJ:
			patType = append(patType, "体检")
		default:
			patType = append(patType, "其他")
		}
	}

	for _, stu := range object.StudyType {
		switch stu {
		case global.Study_Type_XRay:
			stuType = append(stuType, "X-Ray")
		case global.Study_Type_DR:
			stuType = append(stuType, "DR")
		case global.Study_Type_CT:
			stuType = append(stuType, "CT")
		case global.Study_Type_MR:
			stuType = append(stuType, "MR")
		case global.Study_Type_DSA:
			stuType = append(stuType, "DSA")
		case global.Study_Type_US:
			stuType = append(stuType, "US")
		case global.Study_Type_ES:
			stuType = append(stuType, "ES")
		case global.Study_Type_PA:
			stuType = append(stuType, "PA")
		case global.Study_Type_NM:
			stuType = append(stuType, "NM")
		case global.Study_Type_PET:
			stuType = append(stuType, "PET")
		default:
			stuType = append(stuType, "OT")
		}
	}

	sql := `select his_request_id,patient_name,patient_type_code,patient_type_name,medical_record_number,
	sex_code,sex_name,age,age_unit,birthday,modality_code,project_code,project_name,bodypart_code,bodypart,
	outpatient_number,inhospital_number,visit_card_number,phone_number,inp_ward_id,patient_section_name,
	sickbed_number,request_time,his_request_detail_id,id_card_number,address,clinical_diagnosis,medical_history,
	request_department_code,request_department_name,request_doctor_code,request_doctor_name,check_note,film_count,
	film_type,graphic_report,emergency,fee 
	from V_PACS_HZ where 1 = 1`
	// 参数1
	var param1str string
	for i := 0; i < param1len; i++ {
		if i > 0 {
			param1str += " or "
		}
		switch object.PARAM[i].ParamType {
		case global.Apply_Param_JZKH:
			param1str += "visit_card_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_MZH:
			param1str += "outpatient_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_ZYH:
			param1str += "inhospital_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_BLH:
			param1str += "medical_record_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_TJH:
			param1str += "outpatient_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_MZSQDH:
			param1str += "his_request_id = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_ZYSQDH:
			param1str += "his_request_id = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_TJSQDH:
			param1str += "his_request_id = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_SFZH:
			param1str += "id_card_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_XM:
			param1str += "patient_name = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_JZ:
			param1str += "emergency = '" + object.PARAM[i].ParamValue + "'"
		default:
			param1str += "1 = 1"
		}
	}
	if param1len > 0 {
		sql += " and("
		sql += param1str
		sql += ")"
	}
	// 参数就诊类型
	var parampatType string
	for i := 0; i < len(patType); i++ {
		if i > 0 {
			parampatType += ","
		}
		parampatType += "'" + patType[i] + "'"
	}
	if len(patType) > 0 {
		sql += " and("
		sql += "patient_type_name in ("
		sql += parampatType
		sql += "))"

	}

	// 参数检查类型
	var paramstuType string
	for i := 0; i < len(stuType); i++ {
		if i > 0 {
			paramstuType += ","
		}
		paramstuType += "'" + stuType[i] + "'"
	}
	if len(stuType) > 0 {
		sql += " and("
		sql += "modality_code in ("
		sql += paramstuType
		sql += "))"
	}

	// 参数2
	var param2str string
	for i := 0; i < param2len; i++ {
		if i > 0 {
			param2str += " or "
		}
		if (object.PARAM2[i].StartDate != "") && (object.PARAM2[i].EndDate != "") {
			param2str += " request_time between '" + object.PARAM2[i].StartDate + "' and '" + object.PARAM2[i].EndDate + "'"
		} else {
			param2str += "1=1"
		}
	}
	if param2len > 0 {
		sql += " and ("
		sql += param2str
		sql += ")"
	}

	// 排序
	// 单独返回条件数据总数(不仅仅是分页后的数据)
	count = model.GetDataCount(sql)
	// 参数分页
	if object.StartSize >= 0 && object.EndSize > 0 {
		sql += " limit "
		sql += strconv.Itoa(object.StartSize) + "," + strconv.Itoa(object.EndSize)
	}
	global.Logger.Debug("执行的sql语句是: ", sql)
	data = model.GetZLHisViewApply(sql)
	return
}
