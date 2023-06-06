package model

// 中联his 数据包

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"database/sql"
	"strconv"
	"strings"
)

type ZLHISApply struct {
	Pat_Info   PatientInfo
	Apply_Info ApplyInfo
}

// 患者信息
type PatientInfo struct {
	Pat_id         sql.NullString // 患者ID
	Pat_idno       sql.NullString // 身份证号
	Pat_si_no      sql.NullString // 医保号（患者信息）
	Pat_name       sql.NullString // 患者姓名
	Pat_spell_name sql.NullString // 患者姓名拼音
	Pat_sex_code   sql.NullString // 患者性别代码
	Pat_sex        sql.NullString // 患者性别：字典：生理性别字典表
	Pat_brsdate    sql.NullString // 患者出生日期：格式：yyyy-MM-dd
	Pat_age        sql.NullInt16  // 病人接诊年龄
	Pat_age_unit   sql.NullString // 病人年龄单位：字典
	Pat_weight     sql.NullInt16  // 患者体重
	Pat_mari_code  sql.NullInt16  // 婚姻状况
	Pat_tel        sql.NullString // 病人联系电话
	Pat_addr       sql.NullString // 病人联系地址
}

// 申请信息
type ApplyInfo struct {
	Apply_hospital_id            sql.NullInt16  // 医院编号
	Apply_id                     sql.NullString // 申请单号
	Apply_status                 sql.NullInt16  // 申请单状态
	Apply_time                   sql.NullString // 申请时间:格式：yyyy-MM-dd HH:mm:ss
	Apply_department_id          sql.NullString // 申请科室ID
	Apply_department             sql.NullString // 申请科室
	Apply_doctor_id              sql.NullString // 申请医生ID
	Apply_doctor                 sql.NullString // 申请医生
	Apply_pat_type_code          sql.NullString // 患者类别CODE(IH/OP)
	Apply_pat_type               sql.NullString // 患者类别
	Apply_clinic_id              sql.NullString // 门诊号:（门诊特有）住院号:（住院特有）体检健康号:（体检特有）
	Apply_visit_card_no          sql.NullString // 就诊卡号
	Apply_medical_record         sql.NullString // 病历号
	Apply_pat_body_sign          sql.NullString // 体征
	Apply_pat_symptoms           sql.NullString // 症状
	Apply_indications            sql.NullString // 适应症
	Apply_clinical_manifestation sql.NullString // 临床表现
	Apply_clinical_diagnosis     sql.NullString // 临床诊断
	Apply_chief_complaint        sql.NullString // 患者主诉
	Apply_illness_history        sql.NullString // 病人现病史
	Apply_report_flag            sql.NullString // 图文报告标志
	Apply_film_count             sql.NullInt16  // 胶片数量
	Apply_film_flag              sql.NullString // 胶片标志
	Apply_fee                    sql.NullString // 费用
	Apply_vip_flag               sql.NullInt16  // 病人VIP标志
	Apply_isolation_flag         sql.NullInt16  // 隔离标志
	Apply_mergency_status        sql.NullString // 急诊状态
	Apply_study_type             sql.NullString // 检查类型 （CT/MR 检查类别字典表）
	Apply_bodyparts_id           sql.NullString // 检查部位ID:多部位用|分割
	Apply_bodyparts_name         sql.NullString // 检查部位名称：多部位用|分割
	Apply_checkitems_id          sql.NullString // 检查项目ID：多个用|分割
	Apply_checkitems_name        sql.NullString // 检查项目名称：多个用|分割
	Apply_details_id             sql.NullString // 项目明细ID：多个用|分割
	Apply_check_note             sql.NullString // 检查备注
	Apply_check_room             sql.NullString // 检查机房
	Apply_section_id             sql.NullString // 病区名称ID：（住院特有）
	Apply_section                sql.NullString // 病区名称：（住院特有）
	Apply_sicked_index           sql.NullString // 床号：（住院特有）
}

// 获取申请单数据
func GetZLHisViewApply(sql string) (data []global.ApplyFormResultData) {
	global.Logger.Debug("开始查询视图数据.....")
	var err error
	err = global.OracleDBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.OracleDBEngine, _ = NewOracleDBEngine(global.DatabaseSetting)
	}
	rows, err := global.OracleDBEngine.Query(sql)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		key := ZLHISApply{}
		rows.Scan(&key.Apply_Info.Apply_id, &key.Pat_Info.Pat_name, &key.Apply_Info.Apply_pat_type_code, &key.Apply_Info.Apply_pat_type,
			&key.Apply_Info.Apply_medical_record, &key.Pat_Info.Pat_sex_code, &key.Pat_Info.Pat_sex, &key.Pat_Info.Pat_age, &key.Pat_Info.Pat_age_unit,
			&key.Pat_Info.Pat_brsdate, &key.Apply_Info.Apply_study_type, &key.Apply_Info.Apply_checkitems_id, &key.Apply_Info.Apply_checkitems_name,
			&key.Apply_Info.Apply_bodyparts_id, &key.Apply_Info.Apply_bodyparts_name, &key.Apply_Info.Apply_clinic_id, &key.Apply_Info.Apply_clinic_id,
			&key.Apply_Info.Apply_visit_card_no, &key.Pat_Info.Pat_tel, &key.Apply_Info.Apply_section_id, &key.Apply_Info.Apply_section, &key.Apply_Info.Apply_sicked_index,
			&key.Apply_Info.Apply_time, &key.Apply_Info.Apply_details_id, &key.Pat_Info.Pat_idno, &key.Pat_Info.Pat_addr, &key.Apply_Info.Apply_clinical_diagnosis,
			&key.Apply_Info.Apply_illness_history, &key.Apply_Info.Apply_department_id, &key.Apply_Info.Apply_department, &key.Apply_Info.Apply_doctor_id,
			&key.Apply_Info.Apply_doctor, &key.Apply_Info.Apply_check_note, &key.Apply_Info.Apply_film_count, &key.Apply_Info.Apply_film_flag, &key.Apply_Info.Apply_report_flag,
			&key.Apply_Info.Apply_mergency_status, &key.Apply_Info.Apply_fee)

		var sex int
		if key.Pat_Info.Pat_sex.String == "男" {
			sex = 1
		} else if key.Pat_Info.Pat_sex.String == "女" {
			sex = 2
		} else {
			sex = 9
		}

		var ageUnit = 1
		if key.Pat_Info.Pat_age_unit.String == "岁" {
			ageUnit = 1
		} else if key.Pat_Info.Pat_age_unit.String == "月" {
			ageUnit = 2
		} else if key.Pat_Info.Pat_age_unit.String == "周" {
			ageUnit = 3
		} else if key.Pat_Info.Pat_age_unit.String == "天" {
			ageUnit = 4
		} else if key.Pat_Info.Pat_age_unit.String == "时" {
			ageUnit = 5
		} else if key.Pat_Info.Pat_age_unit.String == "分" {
			ageUnit = 6
		} else if key.Pat_Info.Pat_age_unit.String == "秒" {
			ageUnit = 7
		}

		patinfo := global.PatientInfo{
			Pat_name:     key.Pat_Info.Pat_name.String,
			Pat_sex_code: key.Pat_Info.Pat_sex_code.String,
			Pat_sex:      sex,
			Pat_age:      int(key.Pat_Info.Pat_age.Int16),
			Pat_age_unit: ageUnit,
			Pat_brsdate:  key.Pat_Info.Pat_brsdate.String[0:10],
			Pat_tel:      key.Pat_Info.Pat_tel.String,
			Pat_idno:     key.Pat_Info.Pat_idno.String,
			Pat_addr:     key.Pat_Info.Pat_addr.String,
		}

		var pattype int
		if key.Apply_Info.Apply_pat_type.String == "门诊" {
			pattype = 1
		} else if key.Apply_Info.Apply_pat_type.String == "住院" {
			pattype = 2
		} else if key.Apply_Info.Apply_pat_type.String == "体检" {
			pattype = 3
		} else {
			pattype = 9
		}

		var studytype int
		if key.Apply_Info.Apply_study_type.String == "CT" {
			studytype = 3
		} else if key.Apply_Info.Apply_study_type.String == "DR" {
			studytype = 2
		} else if key.Apply_Info.Apply_study_type.String == "MR" {
			studytype = 4
		} else if key.Apply_Info.Apply_study_type.String == "DSA" {
			studytype = 5
		} else if key.Apply_Info.Apply_study_type.String == "US" {
			studytype = 6
		} else if key.Apply_Info.Apply_study_type.String == "ES" {
			studytype = 7
		} else if key.Apply_Info.Apply_study_type.String == "PA" {
			studytype = 8
		} else if key.Apply_Info.Apply_study_type.String == "NM" {
			studytype = 9
		} else if key.Apply_Info.Apply_study_type.String == "PET" {
			studytype = 10
		} else {
			studytype = 99
		}

		film_flag, _ := strconv.Atoi(key.Apply_Info.Apply_film_flag.String)
		report_flag, _ := strconv.Atoi(key.Apply_Info.Apply_report_flag.String)
		mergency_status, _ := strconv.Atoi(key.Apply_Info.Apply_mergency_status.String)

		// 检查项目检查部位处理
		checkBodysCode := strings.Split(key.Apply_Info.Apply_bodyparts_id.String, "|")
		checkBodysName := strings.Split(key.Apply_Info.Apply_bodyparts_name.String, "|")
		checkItemsCode := strings.Split(key.Apply_Info.Apply_checkitems_id.String, "|")
		checkItemsName := strings.Split(key.Apply_Info.Apply_checkitems_name.String, "|")

		var bodymap = make(map[string]bool)
		var bodysarr []global.CheckBody

		for i := 0; i < len(checkBodysCode) && i < len(checkBodysName) && i < len(checkItemsCode) && i < len(checkItemsCode) && i < len(checkItemsName); i++ {
			var item global.CheckItem
			var body global.CheckBody
			var itemsarr []global.CheckItem
			bodycode := checkBodysCode[i]
			bodyname := checkBodysName[i]
			itemcode := checkItemsCode[i]
			itemname := checkItemsName[i]
			if bodymap[bodycode] {
				// 存在
				item.Apply_check_item_code = itemcode
				item.Apply_check_item_name = itemname
				for i := 0; i < len(bodysarr); i++ {
					if (bodysarr[i].Apply_bodypart_code) == bodycode {
						bodysarr[i].Apply_checkItems = append(bodysarr[i].Apply_checkItems, item)
					}
				}

			} else {
				// 不存在
				bodymap[bodycode] = true
				item.Apply_check_item_code = itemcode
				item.Apply_check_item_name = itemname
				itemsarr = append(itemsarr, item)
				body.Apply_bodypart_code = bodycode
				body.Apply_bodypart_name = bodyname
				body.Apply_checkItems = itemsarr
				bodysarr = append(bodysarr, body)
			}
		}

		applyinfo := global.ApplyInfo{
			Apply_id:                 key.Apply_Info.Apply_id.String,
			Apply_pat_type_code:      key.Apply_Info.Apply_pat_type_code.String,
			Apply_pat_type:           pattype,
			Apply_medical_record:     key.Apply_Info.Apply_medical_record.String,
			Apply_study_type:         studytype,
			Apply_checkitems_id:      key.Apply_Info.Apply_checkitems_id.String,
			Apply_checkitems_name:    key.Apply_Info.Apply_checkitems_name.String,
			Apply_bodyparts_id:       key.Apply_Info.Apply_bodyparts_id.String,
			Apply_bodyparts_name:     key.Apply_Info.Apply_bodyparts_name.String,
			Apply_bodys:              bodysarr,
			Apply_clinic_id:          key.Apply_Info.Apply_clinic_id.String,
			Apply_visit_card_no:      key.Apply_Info.Apply_visit_card_no.String,
			Apply_section_id:         key.Apply_Info.Apply_section_id.String,
			Apply_section:            key.Apply_Info.Apply_section.String,
			Apply_sicked_index:       key.Apply_Info.Apply_sicked_index.String,
			Apply_time:               key.Apply_Info.Apply_time.String,
			Apply_details_id:         key.Apply_Info.Apply_details_id.String,
			Apply_clinical_diagnosis: key.Apply_Info.Apply_clinical_diagnosis.String,
			Apply_illness_history:    key.Apply_Info.Apply_illness_history.String,
			Apply_department_id:      key.Apply_Info.Apply_department_id.String,
			Apply_department:         key.Apply_Info.Apply_department.String,
			Apply_doctor_id:          key.Apply_Info.Apply_doctor_id.String,
			Apply_doctor:             key.Apply_Info.Apply_doctor.String,
			Apply_check_note:         key.Apply_Info.Apply_check_note.String,
			Apply_film_count:         int(key.Apply_Info.Apply_film_count.Int16),
			Apply_film_flag:          film_flag,
			Apply_report_flag:        report_flag,
			Apply_mergency_status:    mergency_status,
			Apply_fee:                key.Apply_Info.Apply_fee.String,
		}
		obj := global.ApplyFormResultData{
			Apply_Info: applyinfo,
			Pat_Info:   patinfo,
		}
		data = append(data, obj)
	}
	return
}

// 获取数据总数
func GetDataCount(sql string) (count int) {
	global.Logger.Debug("开始查询条件数据总数.....")
	var err error
	err = global.OracleDBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.OracleDBEngine, _ = NewOracleDBEngine(global.DatabaseSetting)
	}
	rows, err := global.OracleDBEngine.Query(sql)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		count++
	}
	return
}
