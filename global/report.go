package global

import "database/sql"

// 报告数据全局变量

// 申请单状态请求
type ReportData struct {
	Bizno string     `json:"bizno" binding:"required"`
	Time  string     `json:"time" binding:"required"`
	PARAM ReportInfo `json:"report_info"`
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

// 报告回写信息
type ReportInfo struct {
	HospitalId                string                  `json:"hospitalId"`                // 医院ID
	StudyInstanceUid          string                  `json:"studyInstanceUid"`          // Dicom中study_instance_uid
	RegisterId                string                  `json:"registerId"`                // 登记单唯一ID
	RegisterStatus            string                  `json:"registerStatus"`            // 申请单状态
	PositiveNegativeStatus    int                     `json:"positiveNegativeStatus"`    // 阴阳性（0-阴性，1-阳性）
	PathologyStatus           int                     `json:"pathologyStatus"`           // 病理状态（0-阴性，1-阳性）
	Finding                   string                  `json:"finding"`                   // 影像所见
	Conclusion                string                  `json:"conclusion"`                // 诊断结论
	ReportTime                string                  `json:"reportTime"`                // 报告时间
	ReportDoctorId            string                  `json:"reportDoctorId"`            // 报告医生id
	ReportDoctorName          string                  `json:"reportDoctorName"`          // 报告医生名称
	ReportDoctorCode          string                  `json:"reportDoctorCode"`          // 报告医生编码
	HisReportDoctorCode       string                  `json:"hisReportDoctorCode"`       // his报告医生编号
	AuditTime                 string                  `json:"auditTime"`                 // 审核时间
	ReportUpdateTime          string                  `json:"reportUpdateTime"`          // 报告修订时间
	AuditDoctorId             string                  `json:"auditDoctorId"`             // 审核医生id
	AuditDoctorName           string                  `json:"auditDoctorName"`           // 审核医生名称
	AuditDoctorCode           string                  `json:"auditDoctorCode"`           // 审核医生编码
	HisAuditDoctorCode        string                  `json:"hisAuditDoctorCode"`        // his审核医生编号
	InfectiousDiseaseInfoList []InfectiousDiseaseInfo `json:"infectiousDiseaseInfoList"` // 传染病
	ImageQuality              ImageQualityObj         `json:"imageQuality"`              // 图像质控
	InfoQuality               InfoQualityObj          `json:"infoQuality"`               // 报告信息质控
	CrisisInfo                CrisisInfoObj           `json:"crisisInfo"`                // 危机值
}

// 传染病
type InfectiousDiseaseInfo struct {
	InfectiousDiseaseCode string `json:"infectiousDiseaseCode"` // 传染病编码
	InfectiousDiseaseName string `json:"infectiousDiseaseName"` // 传染病名称
}

// 图像质控
type ImageQualityObj struct {
	DoctorAdvice      string               `json:"doctorAdvice"`      // 医生建议
	ImageQualityGrade string               `json:"imageQualityGrade"` // 图像质控等级
	QualityDetailList []ImageQualityDetail `json:"qualityDetailList"` // 图像质控明细
}

// 图像质控明细
type ImageQualityDetail struct {
	ProjectName string `json:"projectName"` // 检查项目名称
	Content     string `json:"content"`     // 评价内容和方法
	Points      int    `json:"points"`      // 扣分

}

// 报告信息质控
type InfoQualityObj struct {
	DoctorAdvice       string                `json:"doctorAdvice"`       // 医生建议
	ReportQualityGrade string                `json:"reportQualityGrade"` // 报告质控等级
	QualityDetailList  []ReportQualityDetail `json:"qualityDetailList"`  // 报告质控明细
}

// 报告质控明细
type ReportQualityDetail struct {
	ProjectName string `json:"projectName"` // 检查项目名称
	Content     string `json:"content"`     // 评价内容和方法
	Points      int    `json:"points"`      // 扣分
}

// 危机值
type CrisisInfoObj struct {
	CrisisStatus             int    `json:"crisisStatus"`             // 危急值状态（0-阴性，1-阳性）
	CrisisContent            string `json:"crisisContent"`            // 危机值内容
	RequestDoctorCode        string `json:"requestDoctorCode"`        // 申请医生代码
	RequestDoctorName        string `json:"requestDoctorName"`        // 申请医生名称
	RequestDoctorPhoneNumber string `json:"requestDoctorPhoneNumber"` // 申请医生手机号
	PatientPhoneNumber       string `json:"patientPhoneNumber"`       // 患者手机号
	ProcessContent           string `json:"processContent"`           // 处理内容
	ReceiverCode             string `json:"receiverCode"`             // 获知人员代码
	ReceiverName             string `json:"receiverName"`             // 获知人员名称
	ProcessTime              string `json:"processTime"`              // 处理时间
	WarnHisStatus            int    `json:"warnHisStatus"`            // 通知HIS状态 0-未通知 1-已通知
	WarnPacsStatus           int    `json:"warnPacsStatus"`           // 通知pacs状态 0-未通知 1-已通知
}

// QYPACS登记检查相关信息
type QYPACSRegisterInfo struct {
	HospitalID         sql.NullString // 医院ID
	RegisterId         sql.NullString // 申请单唯一ID(PACS中)
	AccessionNumber    sql.NullString // 检查号
	RegisterStatus     sql.NullInt16  // 申请单状态
	RegisterDoctorId   sql.NullString // 登记医生id
	RegisterDoctorCode sql.NullString // 登记医生Code
	RegisterDoctorName sql.NullString // 登记医生名称
	RegisterTime       sql.NullString // 登记时间
	DeviceId           sql.NullString // 机房唯一id
	DeviceCode         sql.NullString // 机房编码
	DeviceName         sql.NullString // 机房名称
	ApplyId            sql.NullString // 申请单编号（HIS中）
	ApplydetailId      sql.NullString // his申请明细id（HIS中回写）
	ProjectCode        sql.NullString // 检查项目编码
	ProjectName        sql.NullString // 检查项目名称
	StudyTime          sql.NullString // 检查时间
	StudyDoctorId      sql.NullString // 检查医生id
	StudyDoctorCode    sql.NullString // 检查医生编号
	StudyDoctorName    sql.NullString // 检查医生名称
	UpdateTime         sql.NullString // 登记记录最后更新时间
}

// QYPACS报告相关信息
type QYPACSReportInfo struct {
	HospitalID             sql.NullString // 医院ID
	ReportTime             sql.NullString // 报告时间
	ReportDoctorId         sql.NullString // 报告医生id
	ReportDoctorCode       sql.NullString // 报告医生编号
	ReportDoctorName       sql.NullString // 报告医生名称
	AuditTime              sql.NullString // 审核时间
	AuditDoctorId          sql.NullString // 审核医生id
	AuditDoctorCode        sql.NullString // 审核医生编号
	AuditDoctorName        sql.NullString // 审核医生名称
	PositiveNegativeStatus sql.NullInt16  // 阴阳性状态，0-阴性，1-阳性
	CrisisStatus           sql.NullInt16  // 危急值状态，0-阴性，1-阳性
	Finding                sql.NullString // 影像所见
	Conclusion             sql.NullString // 诊断结论
}

// QYPACS危急值相关信息
type QYPACSCrisisInfo struct {
	HospitalID               sql.NullString // 医院ID
	CrisisContent            sql.NullString // 危机值内容
	RequestDoctorCode        sql.NullString // 申请医生代码
	RequestDoctorName        sql.NullString // 申请医生名称
	RequestDoctorPhoneNumber sql.NullString // 申请医生手机号
	PatientPhoneNumber       sql.NullString // 患者手机号
	ProcessContent           sql.NullString // 处理内容
	ReceiverCode             sql.NullString // 获知人员代码
	ReceiverName             sql.NullString // 获知人员名称
	ProcessTime              sql.NullString // 处理时间
	WarnHisStatus            sql.NullInt16  // 通知HIS状态 0-未通知 1-已通知
	WarnPacsStatus           sql.NullInt16  // 通知pacs状态 0-未通知 1-已通知
}

// QYPACS检查项目信息
type QYPACSProjectInfo struct {
	ProjectID   sql.NullString // 检查项目唯一id
	ProjectCode sql.NullString // 检查项目代码
	ProjectName sql.NullString // 检查项目名称
}

// 医院映射项目信息
type HospitalItemInfo struct {
	HospitalProjectId    sql.NullString // 对应医院检查项目唯一id
	HospitalId           sql.NullString // 所属医院
	HospitalProjectCode  sql.NullString // 检查项目代码
	HospitalProjectName  sql.NullString // 检查项目名称
	HospitalModalityCode sql.NullInt16  // 检查类型
	SystemProjectId      sql.NullString // 映射的项目id
}
