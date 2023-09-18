package global

import "database/sql"

// DICOM 文件获取类型
const (
	Dicom_Type_Share     int = 1 // 匿名共享
	Dicom_Type_PassShare int = 2 // 密码共享
	Dicom_Type_FTP       int = 3 // FTP
)

// DICOM 信息
type DicomInfo struct {
	HospitalID        string // 医院ID
	AccessionNumber   string `json:"accession_number"`    // 检查号
	StudyInstanceUid  string `json:"study_instance_uid"`  // DICOM中检查的唯一编号
	SeriesInstanceUid string `json:"series_instance_uid"` // DICOM中序列的唯一编号
	SopInstanceUid    string `json:"sop_instance_uid"`    // DICOM中图像的唯一编号
	DicomFileName     string `json:"dicom_file_name"`     // DICOM文件的存储路径
	FileType          int    `json:"file_type"`           // 数据类型：1：匿名共享，2：密码共享，3：ftp
	Host              string `json:"host"`                // 主机 例如：127.0.0.1
	Port              string `json:"port"`                // 端口 例如：21
	User              string `json:"user"`                // 若为空，则必须支持匿名访问
	Password          string `json:"password"`            // 若为空，则必须支持匿名访问
	UpdateTime        string `json:"update_time"`         // 增量数据时间戳，yyyy-MM-dd hh:mm:ss
}

type DicomDB struct {
	AccessionNumber   sql.NullString // 检查号
	StudyInstanceUid  sql.NullString // DICOM中检查的唯一编号
	SeriesInstanceUid sql.NullString // DICOM中序列的唯一编号
	SopInstanceUid    sql.NullString // DICOM中图像的唯一编号
	FileType          sql.NullInt16  // 数据类型：1：匿名共享，2：密码共享，3：ftp
	DicomFileName     sql.NullString // DICOM文件的存储路径
	Host              sql.NullString // 主机 例如：127.0.0.1
	Port              sql.NullString // 端口 例如：21
	User              sql.NullString // 若为空，则必须支持匿名访问
	Password          sql.NullString // 若为空，则必须支持匿名访问
	UpdateTime        sql.NullString // 增量数据时间戳，yyyy-MM-dd hh:mm:ss
}

var (
	DicomDataChan chan DicomInfo
)
