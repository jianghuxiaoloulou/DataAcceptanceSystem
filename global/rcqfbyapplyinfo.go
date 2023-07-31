package global

import "database/sql"

// 任城区妇保院

type RcqfbtApplyInfo struct {
	HisApplyID            sql.NullString  // 申请单id(申请单号)
	HisApplyJLID          sql.NullString  // 申请记录ID
	PatientName           sql.NullString  // 患者姓名
	PatientSpellName      sql.NullString  // 患者姓名拼音
	PatientTypeCode       sql.NullString  // 就诊类型，比如OP/IH/PE，其中OP为门诊，IH为住院,PE体检
	PatientTypeName       sql.NullString  // 就诊类型名称，比如门诊/住院/体检
	MedicalRecordNumber   sql.NullString  // 病历号
	SexCode               sql.NullString  // 性别代码，比如MAN/FEMALE，MAN-男，FEMALE-女
	SexName               sql.NullString  // 患者性别，比如男/女
	Age                   sql.NullInt16   // 患者年龄
	AgeUnit               sql.NullString  // 年龄单位，比如岁/月/周/天
	Birthday              sql.NullString  // 出生日期，yyyy-MM-dd
	ModalityCode          sql.NullString  // 检查类型，比如CT/MR/DX/US/ES等
	ProjectCode           sql.NullString  // 检查项目代码，多个使用|*|隔开
	ProjectName           sql.NullString  // 检查方法名称 多个使用|*|隔开
	ProjectFee            sql.NullString  // 检查项目费用 多个使用|*|隔开
	ProjectNote           sql.NullString  // 检查项目注意事项,多个使用|*|隔开
	ProjectDetailID       sql.NullString  // 检查项目明细ID，多个使用|*|隔开
	BodypartCode          sql.NullString  // 部位代码,多个使用|*|号隔开（检查部位需要和检查项目一一对应)
	BodyPart              sql.NullString  // 部位部位名称,多个使用|*|隔开（每个检查项目需要对应一个检查部位)
	ProjectCount          sql.NullInt16   // 检查项目数量
	ClinicNumber          sql.NullString  // 门诊号/住院号/体检号
	VisitCardNumber       sql.NullString  // 就诊卡号
	PhoneNumber           sql.NullString  // 患者电话
	PatientSectionCode    sql.NullString  // 住院病区id
	PatientSectionName    sql.NullString  // 住院病区
	SickbedNumber         sql.NullString  // 住院床位号
	RequestTime           sql.NullString  // 申请时间，yyyy-MM-dd hh:mi:ss
	IdCardNumber          sql.NullString  // 身份证号码
	Address               sql.NullString  // 家庭住址
	ClinicalDiagnosis     sql.NullString  // 临床诊断
	MedicalHistory        sql.NullString  // 病史信息
	RequestDepartmentCode sql.NullString  // 申请科室id
	RequestDepartmentName sql.NullString  // 申请科室
	RequestDoctorCode     sql.NullString  // 申请医生id
	RequestDoctorName     sql.NullString  // 申请医生名
	CheckNote             sql.NullString  // 检查备注
	FilmCount             sql.NullInt16   // 胶片数量
	FilmType              sql.NullInt16   // 胶片类型，0-无，1-传统，2-数字，3-传统+数字
	GraphicReport         sql.NullInt16   // 是否图文报告，0-否，1-是
	Emergency             sql.NullInt16   // 急诊标志，0-否，1-是
	IsolationFlag         sql.NullInt16   // 隔离标志:（0-否，1-是）
	GreenchanFlag         sql.NullInt16   // 绿色通道标志:1-是，0-否
	Fee                   sql.NullFloat64 // 费用
	RmethodName           sql.NullString  // 检查方法
	AccessionNumber       sql.NullString  // 检查号
	PatientCode           sql.NullString  // 患者编号
	HisPatientId          sql.NullString  // his患者唯一id
	RegisterStatus        sql.NullInt16   // 申请单状态
	RegisterDoctorId      sql.NullString  // 登记医生id
	RegisterDoctorCode    sql.NullString  // 登记医生编码
	RegisterDoctorName    sql.NullString  // 登记医生名称
	RegisterTime          sql.NullString  // 登记时间
	QueueNumber           sql.NullString  // 排队号
	DeviceId              sql.NullString  // 机房id
	DeviceCode            sql.NullString  // 机房编码
	DeviceName            sql.NullString  // 机房名称
	StudyDoctorId         sql.NullString  // 检查医生id
	StudyDoctorCode       sql.NullString  // 检查医生编码
	StudyDoctorName       sql.NullString  // 检查医生名称
	StudyTime             sql.NullString  // 检查时间
	AssistDoctorId        sql.NullString  // 辅助医生id
	AssistDoctorCode      sql.NullString  // 辅助医生编码
	AssistDoctorName      sql.NullString  // 辅助医生名称
	OperationDoctorId     sql.NullString  // 手术医生id
	OperationDoctorCode   sql.NullString  // 手术医生编码
	OperationDoctorName   sql.NullString  // 手术医生名称
}

type RcqfbtApplyData struct {
	HospitalId            string        `json:"hospitalId"`            // 医院id
	RegisterId            string        `json:"registerId"`            // 登记单id
	PatientCode           string        `json:"patientCode"`           // 患者编号
	HisPatientId          string        `json:"hisPatientId"`          // his患者唯一id
	PatientName           string        `json:"patientName"`           // 患者名称
	PatientSpellName      string        `json:"patientSpellName"`      // 患者拼音名
	PatientSexCode        int           `json:"patientSexCode"`        // 患者性别编码
	PatientSexName        string        `json:"patientSexName"`        // 患者性别名称
	Age                   int           `json:"age"`                   // 年龄
	AgeUnitCode           int           `json:"ageUnitCode"`           // 年龄单位编码
	AgeUnitName           string        `json:"ageUnitName"`           // 年龄单位名称
	Birthday              string        `json:"birthday"`              // 出生日期
	IdCardNumber          string        `json:"idCardNumber"`          // 身份证号
	MedicareCardNumber    string        `json:"medicareCardNumber"`    // 医保卡号
	PhoneNumber           string        `json:"phoneNumber"`           // 手机号
	Address               string        `json:"address"`               // 地址
	ChiefComplaint        string        `json:"chiefComplaint"`        // 患者主诉
	ClinicalManifestation string        `json:"clinicalManifestation"` // 临床表现
	ClinicalDiagnosis     string        `json:"clinicalDiagnosis"`     // 临床诊断
	MedicalHistory        string        `json:"medicalHistory"`        // 病史记录
	CheckMemo             string        `json:"checkMemo"`             // 检查备注
	RequestId             string        `json:"requestId"`             // his申请单记录ID
	RequestNumber         string        `json:"requestNumber"`         // his申请单号id
	RequestTime           string        `json:"requestTime"`           // 申请时间
	MergencyStatus        int           `json:"mergencyStatus"`        // 急诊状态
	ClinicNumber          string        `json:"clinicNumber"`          // 就诊号
	MedicalRecordNumber   string        `json:"medicalRecordNumber"`   // 病历号
	PatientTypeCode       int           `json:"patientTypeCode"`       // 患者类型编码
	PatientTypeName       string        `json:"patientTypeName"`       // 患者类型名称
	ModalityCode          int           `json:"modalityCode"`          // 检查类型编码
	ModalityName          string        `json:"modalityName"`          // 检查类型名称
	RequestDoctorCode     string        `json:"requestDoctorCode"`     // 申请医生编码
	RequestDoctorName     string        `json:"requestDoctorName"`     // 申请医生名称
	RequestDepartmentCode string        `json:"requestDepartmentCode"` // 申请科室编码
	RequestDepartmentName string        `json:"requestDepartmentName"` // 申请科室名称
	PatientSectionCode    string        `json:"patientSectionCode"`    // 病区编码
	PatientSectionName    string        `json:"patientSectionName"`    // 病区名称
	SickbedNumberCode     string        `json:"sickbedNumberCode"`     // 床位号编码
	SickbedNumberName     string        `json:"sickbedNumberName"`     // 床位号名称
	GraphicReport         int           `json:"graphicReport"`         // 图文报告
	FilmType              int           `json:"filmType"`              // 胶片类型
	FilmNum               int           `json:"filmNum"`               // 胶片数量
	IsolationSign         string        `json:"isolationSign"`         // 隔离
	IsGreenChannel        int           `json:"isGreenChannel"`        // 绿色通道
	AccessionNumber       string        `json:"accessionNumber"`       // 检查号
	RegisterStatus        int           `json:"registerStatus"`        // 申请单状态
	RegisterDoctorId      string        `json:"registerDoctorId"`      // 登记医生id
	RegisterDoctorCode    string        `json:"registerDoctorCode"`    // 登记医生编码
	RegisterDoctorName    string        `json:"registerDoctorName"`    // 登记医生名称
	RegisterTime          string        `json:"registerTime"`          // 登记时间
	QueueNumber           string        `json:"queueNumber"`           // 排队号
	DeviceId              string        `json:"deviceId"`              // 机房id
	DeviceCode            string        `json:"deviceCode"`            // 机房编码
	DeviceName            string        `json:"deviceName"`            // 机房名称
	TotalFee              float64       `json:"totalFee"`              // 总费用
	StudyDoctorId         string        `json:"studyDoctorId"`         // 检查医生id
	StudyDoctorCode       string        `json:"studyDoctorCode"`       // 检查医生编码
	StudyDoctorName       string        `json:"studyDoctorName"`       // 检查医生名称
	StudyTime             string        `json:"studyTime"`             // 检查时间
	AssistDoctorId        string        `json:"assistDoctorId"`        // 辅助医生id
	AssistDoctorCode      string        `json:"assistDoctorCode"`      // 辅助医生编码
	AssistDoctorName      string        `json:"assistDoctorName"`      // 辅助医生名称
	OperationDoctorId     string        `json:"operationDoctorId"`     // 手术医生id
	OperationDoctorCode   string        `json:"operationDoctorCode"`   // 手术医生编码
	OperationDoctorName   string        `json:"operationDoctorName"`   // 手术医生名称
	BodyPartList          []BodyPartStr `json:"bodyPartList"`
}

type BodyPartStr struct {
	BodyPartId      string           `json:"bodyPartId"`
	BodyPartCode    string           `json:"bodyPartCode"`
	BodyPartName    string           `json:"bodyPartName"`
	HisBodyPartCode string           `json:"hisBodyPartCode"`
	ProjectList     []ProjectItemStr `json:"projectList"`
}

type ProjectItemStr struct {
	RequestDetailId string  `json:"requestDetailId"`
	ProjectId       string  `json:"projectId"`
	ProjectCode     string  `json:"projectCode"`
	ProjectName     string  `json:"projectName"`
	HisProjectCode  string  `json:"hisProjectCode"`
	Fee             float64 `json:"fee"`
	ProjectNote     string  `json:"projectNote"`
}

// 飞利浦PACS远程申请时请求的数据
type FLPPACSApplyData struct {
	SiteName         string          `json:"SiteName"`         // 区域站点名称（具体请求医院）Y
	PatientID        string          `json:"PatientID"`        // PACS 系统病人ID 号，PACS 中病人唯一标志 Y
	LocalName        string          `json:"LocalName"`        // 病人姓名 Y
	EnglishName      string          `json:"EnglishName"`      // 病人拼音 Y
	ReferenceNo      string          `json:"ReferenceNo"`      // 身份证号
	Birthday         string          `json:"Birthday"`         // 生日(YYYY-MM-DD) Y
	Gender           string          `json:"Gender"`           // 性别(男、女、未知) Y
	Address          string          `json:"Address"`          // 住址
	Telephone        string          `json:"Telephone"`        // 手机号
	RemotePID        string          `json:"RemotePID"`        // HIS 病人ID 号
	Marriage         string          `json:"Marriage"`         // 婚姻状态
	AccNo            string          `json:"AccNo"`            // 放射编号 Y accessionnumber
	ApplyDept        string          `json:"ApplyDept"`        // 申请部门
	ApplyDoctor      string          `json:"ApplyDoctor"`      // 申请医生
	StudyInstanceUID string          `json:"StudyInstanceUID"` // 图像的SUID Y
	CardNo           string          `json:"CardNo"`           // 就诊卡号
	InhospitalNo     string          `json:"InhospitalNo"`     // 住院号
	ClinicNo         string          `json:"ClinicNo"`         // 门诊号
	PatientType      string          `json:"PatientType"`      // 病人来源（住院、门诊、急诊、体检等）Y
	Observation      string          `json:"Observation"`      // 诊断
	HealthHistory    string          `json:"HealthHistory"`    // 病史
	IsEmergency      int             `json:"IsEmergency"`      // 是否急诊(1 急诊；0 非急诊) Y
	BedNo            string          `json:"BedNo"`            // 床号
	CurrentAge       string          `json:"CurrentAge"`       // 年龄(单位岁、月、天) Y
	Registrar        string          `json:"Registrar"`        // 登记员工号 Y
	RegisterDt       string          `json:"RegisterDt"`       // 登记时间 Y
	Technician       string          `json:"Technician"`       // 检查技师工号 Y
	ExamineDt        string          `json:"ExamineDt"`        // 检查时间 Y
	CheckItems       []FLP_CheckItem `json:"Items"`            // 检查项目数据
}

type FLP_CheckItem struct {
	ProcedureCode string `json:"ProcedureCode"` // 检查项目代码 Y
	CheckingItem  string `json:"CheckingItem"`  // 检查项目名称 Y
	ModalityType  string `json:"ModalityType"`  // 检查类型 Y
	Modality      string `json:"Modality"`      // 检查设备 Y
	RemoteRPID    string `json:"RemoteRPID"`    // 申请单号
}

// 获取的区域PACS申请单数据
type QYPacsApplyData struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    RcqfbtApplyData `json:"data"`
}
