package global

// 报告数据全局变量

// 申请单状态请求
type ReportData struct {
	Bizno string     `json:"bizno" binding:"required"`
	Time  string     `json:"time" binding:"required"`
	PARAM ReportInfo `json:"report_info" binding:"required"`
}

// 报告数据
type ReportKeyData struct {
	ReportId       string `db:"report_id"`        // 报告ID
	Uidenc         string `db:"uid_enc"`          // 检查单ID
	SmsStatus      int    `db:"sms_status"`       // 发送短信状态
	PatientID      string `db:"patient_id"`       // 患者ID
	ReportStatus   string `db:"report_status"`    // 报告状态
	Finding        string `db:"finding"`          // 影像所见
	Conclusion     string `db:"conclusion"`       // 诊断结论
	CheckDoctorId  string `db:"check_doctor_id"`  // 检查医生ID
	CheckDoctor    string `db:"check_doctor"`     // 检查医生
	ReportDoctorID string `db:"report_doctor_id"` // 报告医生ID
	ReportDoctor   string `db:"report_doctor"`    // 报告医生
	AuditDoctorID  string `db:"audit_doctor_id"`  // 审核医生ID
	AuditDoctor    string `db:"audit_doctor"`     // 审核医生
	StudyTime      string `db:"study_time"`       // 检查时间
	InitTime       string `db:"init_time"`        // 起草时间
	AuditTime      string `db:"audit_time"`       // 审核时间
	ReportTime     string `db:"report_time"`      // 报告时间
	RegisterUidEnc string `db:"register_uid_enc"` // HIS申请单
}

// 报告信息
type ReportInfo struct {
	HospitalId             string `json:"hospital_id" binding:"required"`        // 医院ID
	RegisterId             string `json:"register_id" binding:"required"`        // 登记单唯一ID
	ReportTime             string `json:"report_time" binding:"required"`        // 报告时间
	ReportDoctorId         string `json:"report_doctor_id" binding:"required"`   // 报告医生ID
	ReportDoctorCode       string `json:"report_doctor_code" binding:"required"` // 报告医生编号
	ReportDoctorName       string `json:"report_doctor_name" binding:"required"` // 报告医生名称
	HisReportDoctorCode    string `json:"his_report_doctor_code"`                // his报告医生编号
	AuditTime              string `json:"audit_time" binding:"required"`         // 审核时间
	AuditDoctorId          string `json:"audit_doctor_id" binding:"required"`    // 审核医生ID
	AuditDoctorCode        string `json:"audit_doctor_code" binding:"required"`  // 审核医生编号
	AuditDoctorName        string `json:"audit_doctor_name" binding:"required"`  // 审核医生名称
	HisAuditDoctorCode     string `json:"his_audit_doctor_code"`                 // his审核医生编号
	PositiveNegativeStatus int    `json:"positive_negative_status"`              // 阴阳性状态，0-阴性，1-阳性
	CrisisStatus           int    `json:"crisis_status"`                         // 危急值状态，0-阴性，1-阳性
	PathologyStatus        int    `json:"pathology_status"`                      // 病理状态，0-阴性，1-阳性
	Finding                string `json:"finding" binding:"required"`            // 影像所见
	Conclusion             string `json:"conclusion" binding:"required"`         // 诊断结论
	ImageQualityGrade      string `json:"image_quality_grade"`                   // 图像质控等级
	ReportQualityGrade     string `json:"report_quality_grade"`                  // 报告质控等级
	PrintStatus            int    `json:"print_status"`                          // 报告打印状态，0-未打印，1-已打印
}
