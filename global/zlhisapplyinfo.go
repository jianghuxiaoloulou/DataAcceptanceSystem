package global

import "database/sql"

// 医院相关信息
type ZLHISApplyInfo struct {
	HisApplyID            sql.NullString // 申请单id(申请单号)
	HisApplyJLID          sql.NullString // 申请记录ID
	PatientName           sql.NullString // 患者姓名
	PatientSpellName      sql.NullString // 患者姓名拼音
	PatientTypeCode       sql.NullString // 就诊类型，比如OP/IH/PE，其中OP为门诊，IH为住院,PE体检
	PatientTypeName       sql.NullString // 就诊类型名称，比如门诊/住院/体检
	MedicalRecordNumber   sql.NullString // 病历号
	SexCode               sql.NullString // 性别代码，比如MAN/FEMALE，MAN-男，FEMALE-女
	SexName               sql.NullString // 患者性别，比如男/女
	Age                   sql.NullInt16  // 患者年龄
	AgeUnit               sql.NullString // 年龄单位，比如岁/月/周/天
	Birthday              sql.NullString // 出生日期，yyyy-MM-dd
	ModalityCode          sql.NullString // 检查类型，比如CT/MR/DX/US/ES等
	ProjectCode           sql.NullString // 检查项目代码，多个使用|*|隔开
	ProjectName           sql.NullString // 检查方法名称 多个使用|*|隔开
	ProjectFee            sql.NullString // 检查项目费用 多个使用|*|隔开
	ProjectNote           sql.NullString // 检查项目注意事项,多个使用|*|隔开
	ProjectDetailID       sql.NullString // 检查项目明细ID，多个使用|*|隔开
	BodypartCode          sql.NullString // 部位代码,多个使用|*|号隔开（检查部位需要和检查项目一一对应)
	BodyPart              sql.NullString // 部位部位名称,多个使用|*|隔开（每个检查项目需要对应一个检查部位)
	ProjectCount          sql.NullInt16  // 检查项目数量
	ClinicNumber          sql.NullString // 门诊号/住院号/体检号
	VisitCardNumber       sql.NullString // 就诊卡号
	PhoneNumber           sql.NullString // 患者电话
	PatientSectionCode    sql.NullString // 住院病区id
	PatientSectionName    sql.NullString // 住院病区
	SickbedNumber         sql.NullString // 住院床位号
	RequestTime           sql.NullString // 申请时间，yyyy-MM-dd hh:mi:ss
	IdCardNumber          sql.NullString // 身份证号码
	Address               sql.NullString // 家庭住址
	ClinicalDiagnosis     sql.NullString // 临床诊断
	MedicalHistory        sql.NullString // 病史信息
	RequestDepartmentCode sql.NullString // 申请科室id
	RequestDepartmentName sql.NullString // 申请科室
	RequestDoctorCode     sql.NullString // 申请医生id
	RequestDoctorName     sql.NullString // 申请医生名
	CheckNote             sql.NullString // 检查备注
	FilmCount             sql.NullInt16  // 胶片数量
	FilmType              sql.NullInt16  // 胶片类型，0-无，1-传统，2-数字，3-传统+数字
	GraphicReport         sql.NullInt16  // 是否图文报告，0-否，1-是
	Emergency             sql.NullInt16  // 急诊标志，0-否，1-是
	IsolationFlag         sql.NullInt16  // 隔离标志:（0-否，1-是）
	GreenchanFlag         sql.NullInt16  // 绿色通道标志:1-是，0-否
	Fee                   sql.NullString // 费用
	RmethodName           sql.NullString // 检查方法
}
