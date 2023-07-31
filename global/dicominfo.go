package global

import "database/sql"

// DICOM 信息
type DicomInfo struct {
	HospitalID        string // 医院ID
	AccessionNumber   string `json:"accession_number"`    // 检查号
	StudyInstanceUid  string `json:"study_instance_uid"`  // DICOM中检查的唯一编号
	SeriesInstanceUid string `json:"series_instance_uid"` // DICOM中序列的唯一编号
	SopInstanceUid    string `json:"sop_instance_uid"`    // DICOM中图像的唯一编号
	DicomFileName     string `json:"dicom_file_name"`     // DICOM文件的存储路径
	User              string `json:"user"`                // 若为空，则必须支持匿名访问
	Password          string `json:"password"`            // 若为空，则必须支持匿名访问
	UpdateTime        string `json:"update_time"`         // 增量数据时间戳，yyyy-MM-dd hh:mm:ss
}

type DicomDB struct {
	AccessionNumber   sql.NullString // 检查号
	StudyInstanceUid  sql.NullString // DICOM中检查的唯一编号
	SeriesInstanceUid sql.NullString // DICOM中序列的唯一编号
	SopInstanceUid    sql.NullString // DICOM中图像的唯一编号
	DicomFileName     sql.NullString // DICOM文件的存储路径
	User              sql.NullString // 若为空，则必须支持匿名访问
	Password          sql.NullString // 若为空，则必须支持匿名访问
	UpdateTime        sql.NullString // 增量数据时间戳，yyyy-MM-dd hh:mm:ss
}

var (
	DicomDataChan chan DicomInfo
)
