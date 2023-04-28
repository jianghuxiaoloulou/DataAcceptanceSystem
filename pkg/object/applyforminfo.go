package object

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
)

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

// 获取申请单数据
func GetApplyFormData(object global.ApplyFormInfoData) (data []global.ApplyFormResultData) {
	// 1. 通过存储过程获取数据
	// 2. 通过接口获取数据
	// 3. 通过接口推送数据（创建数据库接受数据）
	global.Logger.Debug("开始获取申请单数据，请求参数object: ", object)
	data = ByZLHisViewGetApply(object)
	return
}

// 通过数据库视图获取申请单数据
func ByZLHisViewGetApply(object global.ApplyFormInfoData) (data []global.ApplyFormResultData) {
	global.Logger.Debug("开始通过中联视图获取数据：")
	// 查询视图数据
	param1len := len(object.PARAM)
	param2len := len(object.PARAM2)
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
			param1str += "\"visit_card_number\" = " + object.PARAM[i].ParamValue
		case global.Apply_Param_MZH:
			param1str += "\"outpatient_number\" = " + object.PARAM[i].ParamValue
		case global.Apply_Param_ZYH:
			param1str += "\"inhospital_number\" = " + object.PARAM[i].ParamValue
		case global.Apply_Param_KSID:
			param1str += "\"request_department_code\" = " + object.PARAM[i].ParamValue
		case global.Apply_Param_TJH:
			param1str += "\"request_department_code\" = " + object.PARAM[i].ParamValue
		case global.Apply_Param_MZSQDH:
			param1str += "\"his_request_id\" = " + object.PARAM[i].ParamValue
		case global.Apply_Param_ZYSQDH:
			param1str += "\"his_request_id\" = " + object.PARAM[i].ParamValue
		case global.Apply_Param_TJSQDH:
			param1str += "\"his_request_id\" = " + object.PARAM[i].ParamValue
		case global.Apply_Param_SFZH:
			param1str += "\"id_card_number\" = " + object.PARAM[i].ParamValue
		default:
			param1str += "1 = 1"
		}
	}

	sql += " and("
	sql += param1str
	sql += ")"

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
	sql += " and ("
	sql += param2str
	sql += ")"
	global.Logger.Debug("执行的sql语句是: ", sql)
	data = model.GetZLHisViewApply(sql)
	return
}
