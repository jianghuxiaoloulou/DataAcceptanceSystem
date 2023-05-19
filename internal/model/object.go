package model

import "WowjoyProject/DataAcceptanceSystem/global"

// 获取报告信息
func GetReportInfo(reguidenc string) (global.ReportKeyData, error) {
	global.Logger.Info("开始查询报告信息: ", reguidenc)
	sql := `SELECT r.report_id,r.uid_enc,r.sms_status,r.patient_id,r.report_status,r.finding,r.conclusion,r.check_doctor_id,
	r.check_doctor,r.report_doctor,r.report_doctor_id,r.audit_doctor,r.audit_doctor_id,r.study_time,r.init_time,
	r.audit_time,r.report_time,r.register_uid_enc FROM report r 
	LEFT JOIN register_info fi on fi.uid_enc = r.uid_enc
	WHERE fi.register_uid_enc = ?`
	row := global.ReadDBEngine.QueryRow(sql, reguidenc)
	repdata := global.ReportKeyData{}
	err := row.Scan(&repdata.ReportId, &repdata.Uidenc, &repdata.SmsStatus, &repdata.PatientID, &repdata.ReportStatus, &repdata.Finding, &repdata.Conclusion, &repdata.CheckDoctorId,
		&repdata.CheckDoctor, &repdata.ReportDoctor, &repdata.ReportDoctorID, &repdata.AuditDoctor, &repdata.AuditDoctorID, &repdata.StudyTime, &repdata.InitTime,
		&repdata.AuditTime, &repdata.ReportTime, &repdata.RegisterUidEnc)
	if err != nil {
		global.Logger.Error(err)
		return repdata, err
	}
	return repdata, nil
}
