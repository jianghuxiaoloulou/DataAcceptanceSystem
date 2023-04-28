package model

import "database/sql"

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
