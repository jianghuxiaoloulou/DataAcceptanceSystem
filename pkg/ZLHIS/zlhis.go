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

	url := "global.ObjectSetting.HISURL"
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

	url := "global.ObjectSetting.HISURL"
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

	url := "global.ObjectSetting.HISURL"
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

// 获取申请单数据
func GetApplyData(hospital global.HospitalConfig, object global.ApplyFormInfoData) (count int, data []global.ApplyFormResultData) {
	// 1.获取HIS配置信息
	hisconfig, err := model.GetHisConfig(int(hospital.HISType.Int16))
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("获取到的HIS厂商信息：", hisconfig)
	switch hisconfig.ApplyMZDBType.String {
	case global.HisMysql:
		global.Logger.Debug("通过Mysql数据库获取数据")
		count, data = ByMysqlGetApplyData(hisconfig, object)
	case global.HisOracle:
		global.Logger.Debug("通过Oracle数据库获取数据")
		count, data = ByOracleGetApplyData(hisconfig, object)
	default:
		global.Logger.Error("未实现该方法")
	}
	return
}

// 通过MySql数据库获取申请单数据
func ByMysqlGetApplyData(his global.HisConfig, object global.ApplyFormInfoData) (count int, data []global.ApplyFormResultData) {
	global.Logger.Debug("开始通过中联MYSQL视图获取数据：")
	var sql string
	sql = `select his_apply_id,his_apply_jlid,patient_name,patient_spell_name,patient_type_code,patient_type_name,medical_record_number,
		sex_code,sex_name,age,age_unit,birthday,modality_code,project_code,project_name,project_fee,project_note,project_detail_id,
		bodypart_code,bodypart,project_count,clinic_number,visit_card_number,phone_number,patient_section_code,patient_section_name,
		sickbed_number,request_time,id_card_number,address,clinical_diagnosis,medical_history,request_department_code,
		request_department_name,request_doctor_code,request_doctor_name,check_note,film_count,film_type,graphic_report,
		emergency,isolation_flag,greenchan_flag,fee,rmethod_name`
	sql += " from " + his.ApplyMZViewName.String + " where 1 = 1 "

	// 就诊类型参数
	var parampatType string
	for index, value := range object.PatientType {
		// 通过字典获取数据
		name := model.GetDictName(value)
		global.Logger.Debug("获取的患者就诊类型名为：", name)
		if index > 0 {
			parampatType += ","
		}
		parampatType += "'" + name + "'"
	}

	if len(parampatType) > 0 {
		sql += " and("
		sql += "patient_type_name in ("
		sql += parampatType
		sql += "))"
	}

	// 检查类型
	var paramstuType string
	for index, value := range object.StudyType {
		// 通过字典获取数据
		name := model.GetDictName(value)
		global.Logger.Debug("获取的检查诊类型名为：", name)
		if index > 0 {
			paramstuType += ","
		}
		paramstuType += "'" + name + "'"
	}
	if len(paramstuType) > 0 {
		sql += " and("
		sql += "modality_code in ("
		sql += paramstuType
		sql += "))"
	}

	// 急诊状态 2-全部数据，1-急诊数据,0-非急诊数据
	switch object.MergencySta {
	case global.Mergency_Type_True:
		sql += " and emergency = " + strconv.Itoa(object.MergencySta) + ""
	case global.Mergency_Type_False:
		sql += " and emergency = " + strconv.Itoa(object.MergencySta) + ""
	default:
	}

	// 参数关键值
	param1len := len(object.PARAM)
	var param1str string
	for i := 0; i < param1len; i++ {
		if i > 0 {
			param1str += " or "
		}
		switch object.PARAM[i].ParamType {
		case global.Apply_Param_JZKH:
			param1str += "visit_card_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_MZH:
			param1str += "clinic_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_ZYH:
			param1str += "clinic_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_BLH:
			param1str += "medical_record_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_TJH:
			param1str += "clinic_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_MZSQDH:
			param1str += "his_apply_id = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_ZYSQDH:
			param1str += "his_apply_id = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_TJSQDH:
			param1str += "his_apply_id = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_SFZH:
			param1str += "id_card_number = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_XM:
			param1str += "patient_name = '" + object.PARAM[i].ParamValue + "'"
		default:
			param1str += "1 = 1"
		}
	}
	if param1len > 0 {
		sql += " and("
		sql += param1str
		sql += ")"
	}

	// 参数时间
	param2len := len(object.PARAM2)
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

	// 参数排序 0-降序 1-升序；默认降序
	if 1 == object.SortType {
		sql += " order by request_time asc"
	} else {
		sql += " order by request_time desc"
	}

	// 获取临时数据库引擎
	hisDB, err := model.NewTempDBEngine(his.ApplyMZDBType.String, his.ApplyMZDBConn.String)
	if err != nil {
		global.Logger.Error(err)
		return
	}

	// 单独返回条件数据总数(不仅仅是分页后的数据)
	count = model.GetDataCount(hisDB, sql)

	// 分页
	pagenum := (object.PageNum - 1) * object.PageSize
	// 参数分页
	if object.PageNum >= 1 && object.PageSize > 0 {
		sql += " limit "
		sql += strconv.Itoa(pagenum) + "," + strconv.Itoa(object.PageSize)
	}
	global.Logger.Debug("执行的sql语句是: ", sql)
	data = model.GetZLHisViewApply(hisDB, sql)

	return
}

// 通过Oracle数据库获取申请单数据
func ByOracleGetApplyData(his global.HisConfig, object global.ApplyFormInfoData) (count int, data []global.ApplyFormResultData) {
	global.Logger.Debug("开始通过中联ORACLE视图获取数据：")
	var sql string
	sql = `select "his_apply_id","his_apply_jlid","patient_name","patient_spell_name","patient_type_code","patient_type_name","medical_record_number",
	"sex_code","sex_name","age","age_unit","birthday","modality_code","project_code","project_name","project_fee","project_note","project_detail_id",
	"bodypart_code","bodypart","project_count","clinic_number","visit_card_number","phone_number","patient_section_code","patient_section_name",
	"sickbed_number","request_time","id_card_number","address","clinical_diagnosis","medical_history","request_department_code",
	"request_department_name","request_doctor_code","request_doctor_name","check_note","film_count","film_type","graphic_report",
	"emergency","isolation_flag","greenchan_flag","fee","rmethod_name"`
	sql += " from " + his.ApplyMZViewName.String + " where 1 = 1 "

	// 就诊类型参数
	var parampatType string
	for index, value := range object.PatientType {
		// 通过字典获取数据
		name := model.GetDictName(value)
		global.Logger.Debug("获取的患者就诊类型名为：", name)
		if index > 0 {
			parampatType += ","
		}
		parampatType += "'" + name + "'"
	}

	if len(parampatType) > 0 {
		sql += " and("
		sql += "\"patient_type_name\" in ("
		sql += parampatType
		sql += "))"
	}

	// 检查类型
	var paramstuType string
	for index, value := range object.StudyType {
		// 通过字典获取数据
		name := model.GetDictName(value)
		global.Logger.Debug("获取的检查诊类型名为：", name)
		if index > 0 {
			paramstuType += ","
		}
		paramstuType += "'" + name + "'"
	}
	if len(paramstuType) > 0 {
		sql += " and("
		sql += "\"modality_code\" in ("
		sql += paramstuType
		sql += "))"
	}

	// 急诊状态 2-全部数据，1-急诊数据,0-非急诊数据
	switch object.MergencySta {
	case global.Mergency_Type_True:
		sql += " and \"emergency\" = " + strconv.Itoa(object.MergencySta) + ""
	case global.Mergency_Type_False:
		sql += " and \"emergency\" = " + strconv.Itoa(object.MergencySta) + ""
	default:
	}

	// 参数关键值
	param1len := len(object.PARAM)
	var param1str string
	for i := 0; i < param1len; i++ {
		if i > 0 {
			param1str += " or "
		}
		switch object.PARAM[i].ParamType {
		case global.Apply_Param_JZKH:
			param1str += "\"visit_card_number\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_MZH:
			param1str += "\"clinic_number\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_ZYH:
			param1str += "\"clinic_number\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_BLH:
			param1str += "\"medical_record_number\" = '" + object.PARAM[i].ParamValue + "'"
		case global.Apply_Param_TJH:
			param1str += "\"clinic_number\" = '" + object.PARAM[i].ParamValue + "'"
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
		default:
			param1str += "1 = 1"
		}
	}
	if param1len > 0 {
		sql += " and("
		sql += param1str
		sql += ")"
	}

	// 参数时间
	param2len := len(object.PARAM2)
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

	// 参数排序
	if 1 == object.SortType {
		sql += " order by \"request_time\" asc"
	} else {
		sql += " order by \"request_time\" desc"
	}

	// 获取临时数据库引擎
	hisDB, err := model.NewTempDBEngine(his.ApplyMZDBType.String, his.ApplyMZDBConn.String)
	if err != nil {
		global.Logger.Error(err)
		return
	}

	// 单独返回条件数据总数(不仅仅是分页后的数据)
	count = model.GetDataCount(hisDB, sql)

	// 分页
	// Oracle 12c及以上版本支持的新特性
	pagenum := (object.PageNum - 1) * object.PageSize

	if object.PageNum >= 0 && object.PageSize > 0 {
		sql += " OFFSET "
		sql += strconv.Itoa(pagenum)
		sql += " ROWS "
		sql += "FETCH NEXT " + strconv.Itoa(object.PageSize)
		sql += " ROWS ONLY"
	}

	global.Logger.Debug("执行的sql语句是: ", sql)
	data = model.GetZLHisViewApply(hisDB, sql)
	return
}
