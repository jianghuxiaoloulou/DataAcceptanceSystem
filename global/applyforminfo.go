package global

// 申请单请求参数类型
const (
	Apply_Param_JZKH   int = 1 // 就诊卡号
	Apply_Param_MZH    int = 2 // 门诊号
	Apply_Param_ZYH    int = 3 // 住院号
	Apply_Param_KSID   int = 4 // 科室ID
	Apply_Param_TJH    int = 5 // 体检号
	Apply_Param_MZSQDH int = 6 // 门诊申请单号
	Apply_Param_ZYSQDH int = 7 // 住院申请单号
	Apply_Param_TJSQDH int = 8 // 体检申请单号
	Apply_Param_SFZH   int = 9 // 身份证号

)

// 申请单信息请求
type ApplyFormInfoData struct {
	Bizno  string   `json:"bizno" binding:"required"`
	Time   string   `json:"time" binding:"required"`
	PARAM  []Param  `json:"req_info"`
	PARAM2 []Param2 `json:"req_info2"`
}

type Param2 struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	StartSize int    `json:"start_size"`
	EndSize   int    `json:"end_size"`
}

// 申请单返回数据
type ApplyFormResultData struct {
	Pat_Info   PatientInfo `json:"pat_info"`
	Apply_Info ApplyInfo   `json:"apply_info"`
}

// 患者信息
type PatientInfo struct {
	Pat_id         string `db:"" json:"pat_id" binding:"required"`               // 患者ID
	Pat_idno       string `db:"id_card_number" json:"Pat_idno"`                  // 身份证号
	Pat_si_no      string `db:"" json:"pat_si_no"`                               // 医保号（患者信息）
	Pat_name       string `db:"patient_name" json:"Pat_name" binding:"required"` // 患者姓名
	Pat_spell_name string `db:"" json:"pat_spell_name" binding:"required"`       // 患者姓名拼音
	Pat_sex_code   string `db:"sex_code" json:"pat_sex_code" binding:"required"` // 患者性别代码
	Pat_sex        string `db:"sex_name" json:"pat_sex" binding:"required"`      // 患者性别：字典：生理性别字典表
	Pat_brsdate    string `db:"birthday" json:"pat_brsdate" binding:"required"`  // 患者出生日期：格式：yyyy-MM-dd
	Pat_age        int    `db:"age" json:"pat_age" binding:"required"`           // 病人接诊年龄
	Pat_age_unit   string `db:"age_unit" json:"pat_age_unit" binding:"required"` // 病人年龄单位：字典
	Pat_weight     int    `db:"" json:"pat_weight"`                              // 患者体重
	Pat_mari_code  int    `db:"" json:"pat_mari_code"`                           // 婚姻状况
	Pat_tel        string `db:"phone_number" json:"pat_tel"`                     // 病人联系电话
	Pat_addr       string `db:"address" json:"pat_addr"`                         // 病人联系地址
}

// 检查部位
type CheckBody struct {
	Apply_bodypart_code string      `json:"ap ply_bodypart_code"` // 检查部位代码
	Apply_bodypart_name string      `json:"apply_bodypart_name"`  // 检查部位名字
	Apply_checkItems    []CheckItem `json:"apply_checkitems"`     // 检查项目
}

// 检查项目
type CheckItem struct {
	Apply_check_item_code string `json:"apply_check_item_code" binding:"required"` // 检查项目代码
	Apply_check_item_name string `json:"apply_check_item_name" binding:"required"` // 检查项目名字
	Apply_detail_id       string `json:"apply_detail_id" binding:"required"`       // 申请明细序号
	Apply_checkitem_fee   int    `json:"apply_checkitem_fee" binding:"required"`   // 检查项目费用
	Apply_points_note     string `json:"apply_points_note"`                        // 注意事项
}

// 申请信息
type ApplyInfo struct {
	Apply_hospital_id            int         `db:"" json:"apply_hospital_id" binding:"required"`                // 医院编号
	Apply_id                     string      `db:"his_request_id" json:"apply_id" binding:"required"`           // 申请单号
	Apply_status                 int         `db:"" json:"apply_status" binding:"required"`                     // 申请单状态
	Apply_time                   string      `db:"request_time" json:"apply_time" binding:"required"`           // 申请时间:格式：yyyy-MM-dd HH:mm:ss
	Apply_department_id          string      `db:"request_department_code" json:"apply_department_id"`          // 申请科室ID
	Apply_department             string      `db:"request_department_name" json:"apply_department"`             // 申请科室
	Apply_doctor_id              string      `db:"request_doctor_code" json:"apply_doctor_id"`                  // 申请医生ID
	Apply_doctor                 string      `db:"request_doctor_name" json:"apply_doctor"`                     // 申请医生
	Apply_pat_type_code          string      `db:"patient_type_code" json:"apply_pat_type_code"`                // 患者类别CODE(IH/OP)
	Apply_pat_type               string      `db:"patient_type_name" json:"apply_pat_type"`                     // 患者类别
	Apply_clinic_id              string      `db:"outpatient_number" json:"apply_clinic_id" binding:"required"` // 门诊号:（门诊特有）住院号:（住院特有）体检健康号:（体检特有）
	Apply_visit_card_no          string      `db:"visit_card_number" json:"apply_visit_card_no"`                // 就诊卡号
	Apply_medical_record         string      `db:"medical_record_number" json:"apply_medical_record"`           // 病历号
	Apply_pat_body_sign          string      `db:"" json:"apply_pat_body_sign"`                                 // 体征
	Apply_pat_symptoms           string      `db:"" json:"apply_pat_symptoms"`                                  // 症状
	Apply_indications            string      `db:"" json:"apply_indications"`                                   // 适应症
	Apply_clinical_manifestation string      `db:"" json:"apply_clinical_manifestation"`                        // 临床表现
	Apply_clinical_diagnosis     string      `db:"clinical_diagnosis" json:"apply_clinical_diagnosis"`          // 临床诊断
	Apply_chief_complaint        string      `db:"" json:"apply_chief_complaint"`                               // 患者主诉
	Apply_illness_history        string      `db:"medical_history" json:"apply_illness_history"`                // 病人现病史
	Apply_report_flag            string      `db:"graphic_report" json:"apply_report_flag" binding:"required"`  // 图文报告标志
	Apply_film_count             int         `db:"film_count" json:"apply_film_count" binding:"required"`       // 胶片数量
	Apply_film_flag              string      `db:"film_type" json:"apply_film_flag" binding:"required"`         // 胶片标志
	Apply_fee                    string      `db:"fee" json:"fee"`                                              // 费用
	Apply_vip_flag               int         `db:"" json:"apply_vip_flag"`                                      // 病人VIP标志
	Apply_isolation_flag         int         `db:"" json:"apply_isolation_flag"`                                // 隔离标志
	Apply_mergency_status        string      `db:"emergency" json:"apply_mergency_status" binding:"required"`   // 急诊状态
	Apply_study_type             string      `db:"modality_code" json:"apply_study_type"`                       // 检查类型 （CT/MR 检查类别字典表）
	Apply_bodyparts_id           string      `db:"bodypart_code" json:"apply_bodyparts_id"`                     // 检查部位ID:多部位用|分割
	Apply_bodyparts_name         string      `db:"bodypart" json:"apply_bodyparts_name"`                        // 检查部位名称：多部位用|分割
	Apply_checkitems_id          string      `db:"project_code" json:"apply_checkitems_id"`                     // 检查项目ID：多个用|分割
	Apply_checkitems_name        string      `db:"project_name" json:"apply_checkitems_name"`                   // 检查项目名称：多个用|分割
	Apply_details_id             string      `db:"his_request_detail_id" json:"apply_details_id"`               // 项目明细ID：多个用|分割
	Apply_check_note             string      `db:"check_note" json:"apply_check_note"`                          // 检查备注
	Apply_bodys                  []CheckBody `db:"" json:"apply_bodys" binding:"required"`                      // 检查部位
	Apply_check_room             string      `db:"" json:"apply_check_room"`                                    // 检查机房
	Apply_section_id             string      `db:"inp_ward_id" json:"apply_section_id"`                         // 病区名称ID：（住院特有）
	Apply_section                string      `db:"patient_section_name" json:"apply_section"`                   // 病区名称：（住院特有）
	Apply_sicked_index           string      `db:"sickbed_number" json:"apply_sicked_index"`                    // 床号：（住院特有）
}

// 返回结果
// 申请单状态返回
type ApplyFormInfoResult struct {
	Bizno string                `json:"bizno"`
	Time  string                `json:"time"`
	PARAM AckInfo               `json:"ack_info"`
	DATA  []ApplyFormResultData `json:"data"`
}
