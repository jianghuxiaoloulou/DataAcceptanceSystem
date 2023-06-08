package global

// 就诊类型
const (
	Pat_Type_MZ int = 1 // 门诊
	Pat_Type_ZY int = 2 // 住院
	Pat_Type_TJ int = 3 // 体检
	Pat_Type_OT int = 9 // 其它
)

// 检查类型
const (
	Study_Type_XRay int = 1  // X-Ray
	Study_Type_DR   int = 2  // DR
	Study_Type_CT   int = 3  // CT
	Study_Type_MR   int = 4  // MR
	Study_Type_DSA  int = 5  // DSA
	Study_Type_US   int = 6  // US
	Study_Type_ES   int = 7  // ES
	Study_Type_PA   int = 8  // PA
	Study_Type_NM   int = 9  // NM
	Study_Type_PET  int = 10 // PET
	Study_Type_OT   int = 99 // OT
)

// 申请单请求参数类型
const (
	Apply_Param_JZKH   int = 1  // 就诊卡号
	Apply_Param_MZH    int = 2  // 门诊号
	Apply_Param_ZYH    int = 3  // 住院号
	Apply_Param_BLH    int = 4  // 病历号
	Apply_Param_TJH    int = 5  // 体检号
	Apply_Param_MZSQDH int = 6  // 门诊申请单号
	Apply_Param_ZYSQDH int = 7  // 住院申请单号
	Apply_Param_TJSQDH int = 8  // 体检申请单号
	Apply_Param_SFZH   int = 9  // 身份证号
	Apply_Param_XM     int = 10 // 患者姓名
	Apply_Param_JZ     int = 11 //急诊
)

// 申请单信息请求
type ApplyFormInfoData struct {
	Bizno       string   `json:"bizno" binding:"required"`
	Time        string   `json:"time" binding:"required"`
	PatientType []int    `json:"pat_type"`
	StudyType   []int    `json:"study_type"`
	StartSize   int      `json:"page_num"`
	EndSize     int      `json:"page_size"`
	SortType    int      `json:"sort_type"`
	SortValue   int      `json:"sort_value"`
	PARAM       []Param  `json:"req_info"`
	PARAM2      []Param2 `json:"req_info2"`
}

type Param2 struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// 申请单返回数据
type ApplyFormResultData struct {
	Pat_Info   PatientInfo `json:"pat_info"`
	Apply_Info ApplyInfo   `json:"apply_info"`
}

// 患者信息
type PatientInfo struct {
	Pat_id         string `json:"pat_id" binding:"required"`         // 患者ID
	Pat_idno       string `json:"pat_idno"`                          // 身份证号
	Pat_si_no      string `json:"pat_si_no"`                         // 医保号（患者信息）
	Pat_name       string `json:"pat_name" binding:"required"`       // 患者姓名
	Pat_spell_name string `json:"pat_spell_name" binding:"required"` // 患者姓名拼音
	Pat_sex_code   string `json:"pat_sex_code" binding:"required"`   // 患者性别代码
	Pat_sex        int    `json:"pat_sex" binding:"required"`        // 患者性别：字典：生理性别字典表
	Pat_brsdate    string `json:"pat_brsdate" binding:"required"`    // 患者出生日期：格式：yyyy-MM-dd
	Pat_age        int    `json:"pat_age" binding:"required"`        // 病人接诊年龄
	Pat_age_unit   int    `json:"pat_age_unit" binding:"required"`   // 病人年龄单位：字典
	Pat_weight     int    `json:"pat_weight"`                        // 患者体重
	Pat_mari_code  int    `json:"pat_mari_code"`                     // 婚姻状况
	Pat_tel        string `json:"pat_tel"`                           // 病人联系电话
	Pat_addr       string `json:"pat_addr"`                          // 病人联系地址
}

// 检查部位
type CheckBody struct {
	Apply_bodypart_code string      `json:"apply_bodypart_code"` // 检查部位代码
	Apply_bodypart_name string      `json:"apply_bodypart_name"` // 检查部位名字
	Apply_checkItems    []CheckItem `json:"apply_checkitems"`    // 检查项目
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
	Apply_hospital_id            string      `json:"apply_hospital_id" binding:"required"` // 医院编号
	Apply_id                     string      `json:"apply_id" binding:"required"`          // 申请单号
	Apply_jlid                   string      `json:"apply_jlid"`                           // 申请记录ID
	Apply_status                 int         `json:"apply_status" binding:"required"`      // 申请单状态
	Apply_time                   string      `json:"apply_time" binding:"required"`        // 申请时间:格式：yyyy-MM-dd HH:mm:ss
	Apply_department_id          string      `json:"apply_department_id"`                  // 申请科室ID
	Apply_department             string      `json:"apply_department"`                     // 申请科室
	Apply_doctor_id              string      `json:"apply_doctor_id"`                      // 申请医生ID
	Apply_doctor                 string      `json:"apply_doctor"`                         // 申请医生
	Apply_pat_type_code          string      `json:"apply_pat_type_code"`                  // 患者类别CODE(IH/OP/PE)
	Apply_pat_type               int         `json:"apply_pat_type"`                       // 患者类别(住院/门诊/体检)
	Apply_clinic_id              string      `json:"apply_clinic_id" binding:"required"`   // 门诊号:（门诊特有）住院号:（住院特有）体检健康号:（体检特有）
	Apply_visit_card_no          string      `json:"apply_visit_card_no"`                  // 就诊卡号
	Apply_medical_record         string      `json:"apply_medical_record"`                 // 病历号
	Apply_pat_body_sign          string      `json:"apply_pat_body_sign"`                  // 体征
	Apply_pat_symptoms           string      `json:"apply_pat_symptoms"`                   // 症状
	Apply_indications            string      `json:"apply_indications"`                    // 适应症
	Apply_clinical_manifestation string      `json:"apply_clinical_manifestation"`         // 临床表现
	Apply_clinical_diagnosis     string      `json:"apply_clinical_diagnosis"`             // 临床诊断
	Apply_chief_complaint        string      `json:"apply_chief_complaint"`                // 患者主诉
	Apply_illness_history        string      `json:"apply_illness_history"`                // 病人现病史
	Apply_report_flag            int         `json:"apply_report_flag" binding:"required"` // 图文报告标志
	Apply_film_count             int         `json:"apply_film_count" binding:"required"`  // 胶片数量
	Apply_film_flag              int         `json:"apply_film_flag" binding:"required"`   // 胶片标志
	Apply_fee                    string      `json:"fee"`                                  // 费用
	Apply_vip_flag               int         `json:"apply_vip_flag"`                       // 病人VIP标志
	Apply_isolation_flag         int         `json:"apply_isolation_flag"`                 // 隔离标志
	Apply_mergency_status        int         `json:"apply_mergency_status"`                // 急诊状态
	Apply_greenchan_flag         int         `json:"apply_greenchan_flag"`                 // 绿色通道标志:1-是，0-否
	Apply_study_type             int         `json:"apply_study_type"`                     // 检查类型 （CT/MR 检查类别字典表）
	Apply_bodyparts_id           string      `json:"apply_bodyparts_id"`                   // 检查部位ID:多部位用|分割
	Apply_bodyparts_name         string      `json:"apply_bodyparts_name"`                 // 检查部位名称：多部位用|分割
	Apply_checkitems_id          string      `json:"apply_checkitems_id"`                  // 检查项目ID：多个用|分割
	Apply_checkitems_name        string      `json:"apply_checkitems_name"`                // 检查项目名称：多个用|分割
	Apply_details_id             string      `json:"apply_details_id"`                     // 项目明细ID：多个用|分割
	Apply_check_note             string      `json:"apply_check_note"`                     // 检查备注
	Apply_bodys                  []CheckBody `json:"apply_bodys" binding:"required"`       // 检查部位
	Apply_check_room             string      `json:"apply_check_room"`                     // 检查机房
	Apply_section_id             string      `json:"apply_section_id"`                     // 病区名称ID：（住院特有）
	Apply_section                string      `json:"apply_section"`                        // 病区名称：（住院特有）
	Apply_sicked_index           string      `json:"apply_sicked_index"`                   // 床号：（住院特有）
}

// 返回结果
// 申请单状态返回
type ApplyFormInfoResult struct {
	Bizno     string                `json:"bizno"`
	Time      string                `json:"time"`
	PARAM     AckInfo               `json:"ack_info"`
	DataCount int                   `json:"data_count"`
	DATA      []ApplyFormResultData `json:"data"`
}
