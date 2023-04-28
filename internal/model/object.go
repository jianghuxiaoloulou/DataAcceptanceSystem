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

// 获取申请单数据
func GetZLHisViewApply(sql string) (data []global.ApplyFormResultData) {
	global.Logger.Debug("开始查询视图数据.....")
	var err error
	err = global.OracleDBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.OracleDBEngine, _ = NewOracleDBEngine(global.DatabaseSetting)
	}
	rows, err := global.OracleDBEngine.Query(sql)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		key := ZLHISApply{}
		rows.Scan(&key.Apply_Info.Apply_id, &key.Pat_Info.Pat_name, &key.Apply_Info.Apply_pat_type_code, &key.Apply_Info.Apply_pat_type,
			&key.Apply_Info.Apply_medical_record, &key.Pat_Info.Pat_sex_code, &key.Pat_Info.Pat_sex, &key.Pat_Info.Pat_age, &key.Pat_Info.Pat_age_unit,
			&key.Pat_Info.Pat_brsdate, &key.Apply_Info.Apply_study_type, &key.Apply_Info.Apply_checkitems_id, &key.Apply_Info.Apply_checkitems_name,
			&key.Apply_Info.Apply_bodyparts_id, &key.Apply_Info.Apply_bodyparts_name, &key.Apply_Info.Apply_clinic_id, &key.Apply_Info.Apply_clinic_id,
			&key.Apply_Info.Apply_visit_card_no, &key.Pat_Info.Pat_tel, &key.Apply_Info.Apply_section_id, &key.Apply_Info.Apply_section, &key.Apply_Info.Apply_sicked_index,
			&key.Apply_Info.Apply_time, &key.Apply_Info.Apply_details_id, &key.Pat_Info.Pat_idno, &key.Pat_Info.Pat_addr, &key.Apply_Info.Apply_clinical_diagnosis,
			&key.Apply_Info.Apply_illness_history, &key.Apply_Info.Apply_department_id, &key.Apply_Info.Apply_department, &key.Apply_Info.Apply_doctor_id,
			&key.Apply_Info.Apply_doctor, &key.Apply_Info.Apply_check_note, &key.Apply_Info.Apply_film_count, &key.Apply_Info.Apply_film_flag, &key.Apply_Info.Apply_report_flag,
			&key.Apply_Info.Apply_mergency_status, &key.Apply_Info.Apply_fee)

		patinfo := global.PatientInfo{
			Pat_name:     key.Pat_Info.Pat_name.String,
			Pat_sex_code: key.Pat_Info.Pat_sex_code.String,
			Pat_sex:      key.Pat_Info.Pat_sex.String,
			Pat_age:      int(key.Pat_Info.Pat_age.Int16),
			Pat_age_unit: key.Pat_Info.Pat_age_unit.String,
			Pat_brsdate:  key.Pat_Info.Pat_brsdate.String,
			Pat_tel:      key.Pat_Info.Pat_tel.String,
			Pat_idno:     key.Pat_Info.Pat_idno.String,
			Pat_addr:     key.Pat_Info.Pat_addr.String,
		}
		applyinfo := global.ApplyInfo{
			Apply_id:                 key.Apply_Info.Apply_id.String,
			Apply_pat_type_code:      key.Apply_Info.Apply_pat_type_code.String,
			Apply_pat_type:           key.Apply_Info.Apply_pat_type.String,
			Apply_medical_record:     key.Apply_Info.Apply_medical_record.String,
			Apply_study_type:         key.Apply_Info.Apply_study_type.String,
			Apply_checkitems_id:      key.Apply_Info.Apply_checkitems_id.String,
			Apply_checkitems_name:    key.Apply_Info.Apply_checkitems_name.String,
			Apply_bodyparts_id:       key.Apply_Info.Apply_bodyparts_id.String,
			Apply_bodyparts_name:     key.Apply_Info.Apply_bodyparts_name.String,
			Apply_clinic_id:          key.Apply_Info.Apply_clinic_id.String,
			Apply_visit_card_no:      key.Apply_Info.Apply_visit_card_no.String,
			Apply_section_id:         key.Apply_Info.Apply_section_id.String,
			Apply_section:            key.Apply_Info.Apply_section.String,
			Apply_sicked_index:       key.Apply_Info.Apply_sicked_index.String,
			Apply_time:               key.Apply_Info.Apply_time.String,
			Apply_details_id:         key.Apply_Info.Apply_details_id.String,
			Apply_clinical_diagnosis: key.Apply_Info.Apply_clinical_diagnosis.String,
			Apply_illness_history:    key.Apply_Info.Apply_illness_history.String,
			Apply_department_id:      key.Apply_Info.Apply_department_id.String,
			Apply_department:         key.Apply_Info.Apply_department.String,
			Apply_doctor_id:          key.Apply_Info.Apply_doctor_id.String,
			Apply_doctor:             key.Apply_Info.Apply_doctor.String,
			Apply_check_note:         key.Apply_Info.Apply_check_note.String,
			Apply_film_count:         int(key.Apply_Info.Apply_film_count.Int16),
			Apply_film_flag:          key.Apply_Info.Apply_film_flag.String,
			Apply_report_flag:        key.Apply_Info.Apply_report_flag.String,
			Apply_mergency_status:    key.Apply_Info.Apply_mergency_status.String,
			Apply_fee:                key.Apply_Info.Apply_fee.String,
		}
		obj := global.ApplyFormResultData{
			Apply_Info: applyinfo,
			Pat_Info:   patinfo,
		}
		data = append(data, obj)
	}
	return
}
