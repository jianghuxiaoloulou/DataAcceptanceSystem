package model

import "WowjoyProject/DataAcceptanceSystem/global"

// 获取报告信息
func GetReportInfo(reguidenc string) (repdata global.ReportKeyData, err error) {
	global.Logger.Info("开始查询报告信息: ", reguidenc)
	sql := `SELECT r.report_id,r.uid_enc,r.sms_status,r.patient_id,r.report_status,r.finding,r.conclusion,r.check_doctor_id,
	r.check_doctor,r.report_doctor,r.report_doctor_id,r.audit_doctor,r.audit_doctor_id,r.study_time,r.init_time,
	r.audit_time,r.report_time,r.register_uid_enc FROM report r 
	LEFT JOIN register_info fi on fi.uid_enc = r.uid_enc
	WHERE fi.register_uid_enc = ?`
	// 获取临时QYPACS数据库引擎
	QyPacsDB, err := NewTempDBEngine(global.SystemData.QYPacsType, global.SystemData.QYPacsConn)
	if err != nil {
		global.Logger.Error("获取临时数据库引擎db err: ", err.Error())
		return repdata, err
	}
	defer QyPacsDB.Close()
	row := QyPacsDB.QueryRow(sql, reguidenc)
	err = row.Scan(&repdata.ReportId, &repdata.Uidenc, &repdata.SmsStatus, &repdata.PatientID, &repdata.ReportStatus, &repdata.Finding, &repdata.Conclusion, &repdata.CheckDoctorId,
		&repdata.CheckDoctor, &repdata.ReportDoctor, &repdata.ReportDoctorID, &repdata.AuditDoctor, &repdata.AuditDoctorID, &repdata.StudyTime, &repdata.InitTime,
		&repdata.AuditTime, &repdata.ReportTime, &repdata.RegisterUidEnc)
	if err != nil {
		global.Logger.Error(err)
		return repdata, err
	}
	return repdata, nil
}

// 获取HIS人员ID信息
func GetHisPersonID(id, hospitalid string) (hisid string) {
	global.Logger.Info("获取HIS人员ID信息: ", id, " 医院ID: ", hospitalid)
	var err error
	sql := `SELECT his_user_id FROM user_hospital WHERE user_id = ? and hospital_id = ?;`
	// 获取临时QYPACS数据库引擎
	QyPacsDB, err := NewTempDBEngine(global.SystemData.QYPacsType, global.SystemData.QYPacsConn)
	if err != nil {
		global.Logger.Error("获取临时数据库引擎db err: ", err.Error())
		return
	}
	defer QyPacsDB.Close()
	row := QyPacsDB.QueryRow(sql, id, hospitalid)
	err = row.Scan(&hisid)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("查询到的hisCode: ", hisid)
	return
}

// 获取登记检查相关信息
func GetQYPACSRegisterInfo(registerid string) (data []global.QYPACSRegisterInfo) {
	global.Logger.Info("开始查询登记相关信息: ", registerid)
	var err error
	sql := `SELECT rb.hospital_id,rb.register_id,rb.accession_number,rb.register_status,
	rb.register_doctor_id,rb.register_doctor_code,rb.register_doctor_name,rb.register_time,
	rb.device_id,rb.device_code,rb.device_name,ri.request_number,rp.request_detail_id,
	rp.project_code,rp.project_name,rsi.study_time,rsi.study_doctor_id,rsi.study_doctor_code,
	rsi.study_doctor_name,rb.update_time  
	FROM register_base rb 
	LEFT JOIN register_request_info ri ON rb.register_id = ri.register_id 
	LEFT JOIN register_project rp ON rp.register_id = rb.register_id 
	LEFT JOIN register_study_info rsi ON rsi.register_id = rb.register_id 
	WHERE rb.register_id = ?;`
	// 获取临时QYPACS数据库引擎
	QyPacsDB, err := NewTempDBEngine(global.SystemData.QYPacsType, global.SystemData.QYPacsConn)
	if err != nil {
		global.Logger.Error("获取临时数据库引擎db err: ", err.Error())
		return
	}
	defer QyPacsDB.Close()
	rows, err := QyPacsDB.Query(sql, registerid)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		key := global.QYPACSRegisterInfo{}
		err = rows.Scan(&key.HospitalID, &key.RegisterId, &key.AccessionNumber, &key.RegisterStatus,
			&key.RegisterDoctorId, &key.RegisterDoctorCode, &key.RegisterDoctorName, &key.RegisterTime,
			&key.DeviceId, &key.DeviceCode, &key.DeviceName, &key.ApplyId, &key.ApplydetailId,
			&key.ProjectCode, &key.ProjectName, &key.StudyTime, &key.StudyDoctorId, &key.StudyDoctorCode,
			&key.StudyDoctorName, &key.UpdateTime)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		data = append(data, key)
	}
	global.Logger.Debug("获取到查询登记相关数据：", data)
	return
}

// 获取报告相关信息
func GetQYPACSReportInfo(registerid string) (data global.QYPACSReportInfo) {
	global.Logger.Info("获取报告相关信息: ", registerid)
	var err error
	sql := `SELECT ri.hospital_id,ri.report_time,ri.report_doctor_id,ri.report_doctor_code,ri.report_doctor_name,
	ri.audit_time,ri.audit_doctor_id,ri.audit_doctor_code,ri.audit_doctor_name,ri.positive_negative_status,
	ri.crisis_status,ri.finding,ri.conclusion 
	FROM report_info ri 
	WHERE ri.register_id = ?;`
	// 获取临时QYPACS数据库引擎
	QyPacsDB, err := NewTempDBEngine(global.SystemData.QYPacsType, global.SystemData.QYPacsConn)
	if err != nil {
		global.Logger.Error("获取临时数据库引擎db err: ", err.Error())
		return
	}
	defer QyPacsDB.Close()
	row := QyPacsDB.QueryRow(sql, registerid)
	err = row.Scan(&data.HospitalID, &data.ReportTime, &data.ReportDoctorId, &data.ReportDoctorCode, &data.ReportDoctorName,
		&data.AuditTime, &data.AuditDoctorId, &data.AuditDoctorCode, &data.AuditDoctorName, &data.PositiveNegativeStatus,
		&data.CrisisStatus, &data.Finding, &data.Conclusion)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("获取到报告相关数据：", data)
	return
}

// 获取危急值相关信息
func GetQYPACSCrisisInfo(registerid string) (data global.QYPACSCrisisInfo) {
	global.Logger.Info("获取危急值相关信息: ", registerid)
	var err error
	sql := `SELECT ri.hospital_id,ri.crisis_content,ri.request_doctor_code,ri.request_doctor_name,ri.request_doctor_phone_number,
	ri.patient_phone_number,ri.process_content,ri.receiver_code,ri.receiver_name,ri.process_time,
	ri.warn_his_status,ri.warn_pacs_status 
	FROM report_crisis_info ri 
	WHERE ri.register_id = ?;`
	// 获取临时QYPACS数据库引擎
	QyPacsDB, err := NewTempDBEngine(global.SystemData.QYPacsType, global.SystemData.QYPacsConn)
	if err != nil {
		global.Logger.Error("获取临时数据库引擎db err: ", err.Error())
		return
	}
	defer QyPacsDB.Close()
	row := QyPacsDB.QueryRow(sql, registerid)
	err = row.Scan(&data.HospitalID, &data.CrisisContent, &data.RequestDoctorCode, &data.RequestDoctorName, &data.RequestDoctorPhoneNumber,
		&data.PatientPhoneNumber, &data.ProcessContent, &data.ReceiverCode, &data.ReceiverName, &data.ProcessTime,
		&data.WarnHisStatus, &data.WarnPacsStatus)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("获取到危急值相关数据：", data)
	return
}

// 根据区域PACS项目ID 获取医院项目映射信息
func GetHospitialItemInfo(hospitalid, systemItemId string) (data global.HospitalItemInfo) {
	global.Logger.Info("获取医院检查项目相关信息")
	var err error
	sql := `SELECT spm.telemedicine_project_id,spm.telemedicine_project_code,spm.telemedicine_project_name,spm.modality_code 
	from sys_telemedicine_project_mapping spm 
	WHERE spm.hospital_id = ? AND spm.sys_project_id = ?;`
	// 获取临时QYPACS数据库引擎
	QyPacsDB, err := NewTempDBEngine(global.SystemData.QYPacsType, global.SystemData.QYPacsConn)
	if err != nil {
		global.Logger.Error("获取临时数据库引擎db err: ", err.Error())
		return
	}
	defer QyPacsDB.Close()
	row := QyPacsDB.QueryRow(sql, hospitalid, systemItemId)
	err = row.Scan(&data.HospitalProjectId, &data.HospitalProjectCode, &data.HospitalProjectName, &data.HospitalModalityCode)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("获取医院检查项目相关信息: ", data)
	return
}

// 获取区域PACS的检查项目信息
func GetQYPacsItemInfo(hospitalid, itemcode, itemname string) (data global.QYPACSProjectInfo) {
	sql := `SELECT sp.project_id,sp.project_code,sp.project_name FROM sys_telemedicine_project_mapping spm 
	LEFT JOIN sys_project sp ON sp.project_id = spm.sys_project_id
	WHERE spm.hospital_id = ? AND spm.telemedicine_project_code = ? AND spm.telemedicine_project_name = ?;`
	QyPacsDB, err := NewTempDBEngine(global.SystemData.QYPacsType, global.SystemData.QYPacsConn)
	if err != nil {
		global.Logger.Error("获取临时数据库引擎db err: ", err.Error())
		return
	}
	defer QyPacsDB.Close()
	row := QyPacsDB.QueryRow(sql, hospitalid, itemcode, itemname)
	err = row.Scan(&data.ProjectID, &data.ProjectCode, &data.ProjectName)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("检查项目信息：", data)
	return
}
