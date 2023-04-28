package global

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
