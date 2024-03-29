package global

// 系统数据库相关变量

import "database/sql"

var (
	// 字典数据
	DictDatas  []DictDataObject
	SystemData SystemDataObject
	// 字典表中未配置CODE
	OtherCode int = 9999
)

// HIS厂商类型
const (
	HIS_Type_ZLHIS int = 10001 // 中联HIS
	HIS_Type_WDHIS int = 10002 // 万达区域HIS
)

// HIS数据库类型
const (
	HisMysql  string = "mysql"  // HIS数据库mysql
	HisOracle string = "godror" // HIS数据库Oracle
)

// HIS厂商相关信息
type HisConfig struct {
	HISType         sql.NullInt16  // HIS类型：对应HIS厂商类型
	HISTypeName     sql.NullString // HIS类型名：HIS厂商名
	ApplyMZDBType   sql.NullString // 申请单门诊数据库连接类型
	ApplyMZDBConn   sql.NullString // 申请单门诊数据库连接字符串
	ApplyMZViewName sql.NullString // 申请单门诊视图名
	ApplyZYDBType   sql.NullString // 申请单住院数据库连接类型
	ApplyZYDBConn   sql.NullString // 申请单住院数据库连接字符串
	ApplyZYViewName sql.NullString // 申请单住院视图名
	ApplyTJDBType   sql.NullString // 申请单体检数据库连接类型
	ApplyTJDBConn   sql.NullString // 申请单体检数据库连接字符串
	ApplyTJViewName sql.NullString // 申请单体检视图名
	HISInterfaceURL sql.NullString // HIS回写接口URL
}

// 医院相关信息
type HospitalConfig struct {
	HospitalId       sql.NullString // 医院ID
	HospitalName     sql.NullString // 医院名
	HISType          sql.NullInt16  // HIS类型：对应HIS厂商类型
	PacsDBType       sql.NullString // PACS数据库连接类型
	PacsDBConn       sql.NullString // PACS数据库连接字符串
	DicomView        sql.NullString // DICOM视图名
	ApplyView        sql.NullString // PACS提供的申请单视图名
	PacsInterfaceURL sql.NullString // PACS回写接口URL
	UploadTime       sql.NullString // PACS数据上传时记录的时间
}

// 医院申请单状态功能模块配置表
type ApplyFuncConfig struct {
	HospitalId        sql.NullString // 医院ID
	HISType           sql.NullInt16  // HIS类型：对应HIS厂商类型
	ApplyCanceled     sql.NullString // 申请单状态--已取消（多功能通过|*|分割）
	ApplyRegistered   sql.NullString // 申请单状态--已报到（多功能通过|*|分割）
	ApplyChecked      sql.NullString // 申请单状态--已检查（多功能通过|*|分割）
	ApplyDrafted      sql.NullString // 申请单状态--已起草（多功能通过|*|分割）
	ApplyWaitaudit    sql.NullString // 申请单状态--待审核（多功能通过|*|分割）
	ApplyAudited      sql.NullString // 申请单状态--已审核（多功能通过|*|分割）
	ApplyCharging     sql.NullString // 申请单状态--主动计费（多功能通过|*|分割）
	ApplyDiagnose     sql.NullString // 申请单状态--远程诊断（多功能通过|*|分割）
	ApplyViewRemote   sql.NullString // 申请单状态-- 远程查看（多功能通过|*|分割）
	ApplyAuditeRemote sql.NullString // 申请单状态--远程审核（多功能通过|*|分割）
}

type DictConfig struct {
	DictType    sql.NullInt16  // 字段类型
	DictCode    sql.NullInt16  // 字典Code
	DictName    sql.NullString // 字典Name
	DictHisCode sql.NullString // HisCode
	DictHisName sql.NullString // HisName
}

type DictDataObject struct {
	Type    int
	Code    int
	Name    string
	HisCode string
	HisName string
}

// 系统配置

type SystemCfg struct {
	QYPacsType         sql.NullString // 区域PACS数据库连接类型
	QYPacsConn         sql.NullString // 区域PACS数据库连接字符串
	QYPacsInterfaceUrl sql.NullString // 区域PACS接口URL
}

type SystemDataObject struct {
	QYPacsType         string // 区域PACS数据库连接类型
	QYPacsConn         string // 区域PACS数据库连接字符串
	QYPacsInterfaceUrl string // 区域PACS接口URL
}
