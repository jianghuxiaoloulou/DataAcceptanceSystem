package rcqfby

// 任城区妇保院

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 获取申请单数据
func GetApplyData(hospital global.HospitalConfig, object global.ApplyDicomData) {
	global.Logger.Debug("开始通过SQL SERVER 视图获取申请单数据：")
	var sql string
	sql = `select his_apply_id,his_apply_jlid,patient_name,patient_spell_name,patient_type_code,patient_type_name,medical_record_number,
		sex_code,sex_name,age,age_unit,birthday,modality_code,project_code,project_name,project_fee,project_note,project_detail_id,
		bodypart_code,bodypart,project_count,clinic_number,visit_card_number,phone_number,patient_section_code,patient_section_name,
		sickbed_number,request_time,id_card_number,address,clinical_diagnosis,medical_history,request_department_code,
		request_department_name,request_doctor_code,request_doctor_name,check_note,film_count,film_type,graphic_report,
		emergency,isolation_flag,greenchan_flag,fee,rmethod_name,accession_number,patient_code,his_patient_id,register_status,
		register_doctor_id,register_doctor_code,register_doctor_name,register_time,queue_number,device_id,device_code,device_name,
		study_doctor_id,study_doctor_code,study_doctor_name,assist_doctor_id,assist_doctor_code,assist_doctor_name,
		operation_doctor_id,operation_doctor_code,operation_doctor_name`
	sql += " from dbo." + hospital.ApplyView.String + " where 1 = 1 "
	// 通过申请单时间升序获取数据
	sql += " and (" + " request_time between '" + object.PARAM.StartDate + "' and '" + object.PARAM.EndDate + "'" + ")"
	sql += " order by request_time asc"

	global.Logger.Debug("执行的sql: ", sql)

	// 获取临时数据库引擎
	PacsDB, err := model.NewTempDBEngine(hospital.PacsDBType.String, hospital.PacsDBConn.String)
	if err != nil {
		global.Logger.Error("获取临时数据库引擎db err: ", err.Error())
		return
	}

	data := model.GetRcqfbyApplyData(PacsDB, sql, hospital.HospitalId.String)
	global.Logger.Debug("获取的申请单数据：", data)
	for _, value := range data {
		// 上传申请单数据
		go UploadApplyData(PacsDB, hospital.HospitalId.String, value)
		// 获取DICOM数据
		go GetDicomData(PacsDB, hospital, value.AccessionNumber)
	}
}

// 获取申请单数据
func GetApplyDataTime(hospital global.HospitalConfig, object global.ApplyDicomData) {
	global.Logger.Debug("开始通过SQL SERVER 视图获取申请单数据：")
	var sql string
	sql = `select his_apply_id,his_apply_jlid,patient_name,patient_spell_name,patient_type_code,patient_type_name,medical_record_number,
		sex_code,sex_name,age,age_unit,birthday,modality_code,project_code,project_name,project_fee,project_note,project_detail_id,
		bodypart_code,bodypart,project_count,clinic_number,visit_card_number,phone_number,patient_section_code,patient_section_name,
		sickbed_number,request_time,id_card_number,address,clinical_diagnosis,medical_history,request_department_code,
		request_department_name,request_doctor_code,request_doctor_name,check_note,film_count,film_type,graphic_report,
		emergency,isolation_flag,greenchan_flag,fee,rmethod_name,accession_number,patient_code,his_patient_id,register_status,
		register_doctor_id,register_doctor_code,register_doctor_name,register_time,queue_number,device_id,device_code,device_name,
		study_doctor_id,study_doctor_code,study_doctor_name,assist_doctor_id,assist_doctor_code,assist_doctor_name,
		operation_doctor_id,operation_doctor_code,operation_doctor_name`
	sql += " from dbo." + hospital.ApplyView.String + " where 1 = 1 "
	// 通过申请单时间升序获取数据
	sql += " and (" + " request_time between '" + hospital.UploadTime.String + "' and '" + object.PARAM.EndDate + "'" + ")"
	sql += " order by request_time asc"

	global.Logger.Debug("执行的sql: ", sql)

	// 获取临时数据库引擎
	PacsDB, err := model.NewTempDBEngine(hospital.PacsDBType.String, hospital.PacsDBConn.String)
	if err != nil {
		global.Logger.Error("获取临时数据库引擎db err: ", err.Error())
		return
	}

	data := model.GetRcqfbyApplyData(PacsDB, sql, hospital.HospitalId.String)
	global.Logger.Debug("获取的申请单数据：", data)
	for _, value := range data {
		// 上传申请单数据
		go UploadApplyData(PacsDB, hospital.HospitalId.String, value)
		// 获取DICOM数据
		go GetDicomData(PacsDB, hospital, value.AccessionNumber)
	}
}

// 上传申请单数据
func UploadApplyData(db *sql.DB, hospitalid string, data global.RcqfbtApplyData) {
	global.Logger.Debug("开始执行任城区妇保院申请单数据上传", data)
	reqdata, err := json.Marshal(data)
	global.Logger.Debug("上传申请单发送数据：", string(reqdata))
	if err != nil {
		global.Logger.Error(err)
		return
	}
	url := global.SystemData.QYPacsInterfaceUrl
	url += "//"
	url += "ms-qypacs//v1//outside//register//push"
	global.Logger.Debug("操作的URL: ", url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqdata))
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer resp.Body.Close()
	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Info("resp.Body: ", string(resp_body))
	var result = make(map[string]interface{})
	err = json.Unmarshal(resp_body, &result)
	if err != nil {
		global.Logger.Error("resp.Body err ", err)
		return
	}
	// 解析json
	if vCode, ok := result["code"]; ok {
		resultcode := vCode.(string)
		if resultcode == "0" {
			global.Logger.Info("任城区妇保院申请单数据上传成功：", data.AccessionNumber)
			// 数据上传成功后，更新sys_dict_hospital_config表中upload_time时间
			model.UpdateUploadTiem(data.RequestTime, hospitalid)
		} else {
			global.Logger.Error("任城区妇保院申请单数据上传失败：", data.AccessionNumber)
		}
	}
}

// 获取DICOM数据
func GetDicomData(db *sql.DB, hospital global.HospitalConfig, accessionnumber string) {
	global.Logger.Debug("开始通过SQL SERVER 视图获取DICOM数据：")
	var sql string
	sql = `select accession_number,study_instance_uid,series_instance_uid,sop_instance_uid,file_type,dicom_file_name,host,port,user,password,update_time`
	sql += " from dbo." + hospital.DicomView.String + " where accession_number = "
	sql += "'" + accessionnumber + "'"
	global.Logger.Debug("执行的sql: ", sql)
	model.GetRcqfbyDicomData(db, sql, hospital.HospitalId.String)
}

// 上传DICOM数据
func UploadDicomData(data global.DicomInfo) {
	global.Logger.Debug("开始执行任城区妇保院DICOM文件上传", data)
	url := global.SystemData.QYPacsInterfaceUrl
	url += "//"
	url += "ms-dcm-workstation//v1//dicom//store"
	url += "//"
	url += data.HospitalID
	url += "//"
	url += data.AccessionNumber
	global.Logger.Debug("操作的URL: ", url)
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", data.DicomFileName)
	if err != nil {
		global.Logger.Error("error writing to buffer")
		return
	}
	// 判断获取DICOM文件类型
	if data.FileType == global.Dicom_Type_FTP {
		global.Logger.Debug("开始通过FTP方式获取DICOM影像数据")

		// 连接FTP服务器
		conn, err := net.Dial("tcp", data.Host+":"+data.Port)
		if err != nil {
			global.Logger.Debug("连接FTP服务器失败:", err)
			return
		}
		defer conn.Close()

		// 读取服务器返回的欢迎消息
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			global.Logger.Debug("读取欢迎消息失败:", err)
			return
		}
		global.Logger.Debug(string(buff[:n]))

		// 发送登录命令
		cmd := "USER " + data.User + "\r\n"
		_, err = conn.Write([]byte(cmd))
		if err != nil {
			global.Logger.Debug("发送登录命令失败:", err)
			return
		}

		n, err = conn.Read(buff)
		if err != nil {
			global.Logger.Debug("读取登录响应失败:", err)
			return
		}
		global.Logger.Debug(string(buff[:n]))

		// 发送密码命令
		cmd = "PASS " + data.Password + "\r\n"
		_, err = conn.Write([]byte(cmd))
		if err != nil {
			global.Logger.Debug("发送密码命令失败:", err)
			return
		}

		n, err = conn.Read(buff)
		if err != nil {
			global.Logger.Debug("读取密码响应失败:", err)
			return
		}
		global.Logger.Debug(string(buff[:n]))

		// 发送 PASV 命令，进入被动模式
		cmd = "PASV\r\n"
		_, err = conn.Write([]byte(cmd))
		if err != nil {
			global.Logger.Debug("发送PASV命令失败:", err)
			return
		}

		n, err = conn.Read(buff)
		if err != nil {
			global.Logger.Debug("读取PASV响应失败:", err)
			return
		}
		global.Logger.Debug(string(buff[:n]))

		// 解析被动模式返回的地址和端口
		_, port := parsePasvResponse(string(buff[:n]))

		// 主动连接到服务器的数据端口
		dataConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", data.Host, port))
		if err != nil {
			global.Logger.Debug("连接数据端口失败:", err)
			return
		}
		defer dataConn.Close()
		// 发送下载文件命令
		cmd = "RETR " + data.DicomFileName + "\r\n"
		_, err = conn.Write([]byte(cmd))
		if err != nil {
			global.Logger.Debug("发送下载文件命令失败:", err)
			return
		}

		n, err = conn.Read(buff)
		if err != nil {
			global.Logger.Debug("读取下载文件响应失败:", err)
			return
		}
		global.Logger.Debug(string(buff[:n]))

		/////////
		_, err = io.Copy(fileWriter, dataConn)
		if err != nil {
			global.Logger.Error("拷贝数据流量错误: ", err)
			return
		}
		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()
		resp, err := http.Post(url, contentType, bodyBuf)
		if err != nil {
			global.Logger.Error("http post err : ", err)
			return
		}
		defer resp.Body.Close()
		resp_body, _ := io.ReadAll(resp.Body)
		global.Logger.Info("resp.Body: ", string(resp_body))
		global.Logger.Info("Dicom文件上传成功：", data.DicomFileName)

	} else if data.FileType == global.Dicom_Type_Share {
		global.Logger.Debug("开始通过匿名访问共享的方式获取DICOM影像数据")
		//打开文件句柄操作
		fh, err := os.Open(data.DicomFileName)
		if err != nil {
			global.Logger.Error("error opening file")
			return
		}
		//iocopy
		_, err = io.Copy(fileWriter, fh)
		if err != nil {
			return
		}
		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()
		fh.Close()
		resp, err := http.Post(url, contentType, bodyBuf)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		resp_body, _ := io.ReadAll(resp.Body)
		global.Logger.Info("resp.Body: ", string(resp_body))
	}
}

// 上传报告数据到区域PACS
func UploadReportData(data global.ReportInfo) {
	global.Logger.Debug("开始执行济宁医学院报告数据上传区域PACS", data)
	reqdata, err := json.Marshal(data)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	url := global.SystemData.QYPacsInterfaceUrl
	url += "//"
	url += "ms-qypacs//v1//outside//report//writeBack"
	global.Logger.Debug("操作的URL: ", url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqdata))
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer resp.Body.Close()
	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Info("resp.Body: ", string(resp_body))
	var result = make(map[string]interface{})
	err = json.Unmarshal(resp_body, &result)
	if err != nil {
		global.Logger.Error("resp.Body err ", err)
		return
	}
	// 解析json
	if vCode, ok := result["code"]; ok {
		resultcode := vCode.(string)
		if resultcode == "0" {
			global.Logger.Info("任城区妇保院报告数据上传成功：", data.RegisterId)
		} else {
			global.Logger.Error("任城区妇保院报告数据上传失败：", data.RegisterId)
		}
	}
}

// 通过存储过程回写数据到妇保院
func WriteBackProc(data global.RcqfbtApplyData) {
	global.Logger.Debug("开始执行报告数据通过存储过程上传到妇保院", data)
	// 1.通过HospitalID 获取医院相关数据库连接信息
	hospitalConfig, err := model.GetHospitalConfig(data.HospitalId)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("获取的医院相关连接信息：", hospitalConfig)
	// 获取临时数据库引擎
	PacsDB, err := model.NewTempDBEngine(hospitalConfig.PacsDBType.String, hospitalConfig.PacsDBConn.String)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	if len(data.AccessionNumber) > 3 {
		data.AccessionNumber = data.AccessionNumber[3:]
	}
	sql := "exec dbo.RCFY_CT_REPORT @Str_StudyInstanceUID = '" + data.AccessionNumber + "',"
	sql += "@Str_results = '" + data.ReportData.WYG + "',"
	sql += "@Str_finding = '" + data.ReportData.WYS + "',"
	sql += "@Str_reportdoc = '" + data.ReportData.Creater + "',"
	sql += "@Str_auditdoc = '" + data.ReportData.Approver + "',"
	sql += "@Str_WriteDateTime = " + "'" + data.ReportData.CreateDt + "',"
	sql += "@str_ReferringDate = " + "'" + data.ReportData.ApproveDt + "',"
	sql += "@Str_ReportStatus = " + "'3'"

	global.Logger.Debug("执行存储过程的sql: ", sql)
	_, err = PacsDB.Query(sql)
	if err != nil {
		global.Logger.Error("回写存储过程错误，err: ", err)
		return
	}
	global.Logger.Debug("回写报告，调用存储过程完成")
}

// (函数功能D)任城区妇保院报告回写
func ReportDataWriteBack(applyid string) {
	global.Logger.Debug("通过接口获取报告数据：", applyid)
	objdata := GetQYPacsApplyReportData(applyid)
	WriteBackProc(objdata.Data)
}

// (函数功能E)济宁附属医院远程诊断
func SendRemoteDiagnoseApplyData(applyid string) {
	global.Logger.Debug("通过接口获取远程诊断申请单信息：", applyid)
	// 获取区域PACS申请单信息
	objdata := GetQYPacsApplyData(applyid)
	// 获取申请医院的接口信息
	hospitalcfg, err := model.GetHospitalConfig(objdata.Data.HospitalId)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("申请医院信息：", hospitalcfg)
	// 获取上传中心医院的接口信息
	centerHospital, err := model.GetHospitalConfig(objdata.Data.CenterHospitalId)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("中心医院信息：", centerHospital)

	// 发送申请单数据到飞利浦PACS
	global.Logger.Debug("发送申请单数据到济宁附属医院: ", objdata)
	var FlpCheckItems []global.FLP_CheckItem
	for _, body := range objdata.Data.BodyPartList {
		for _, item := range body.ProjectList {
			flpitem := global.FLP_CheckItem{
				ProcedureCode: item.ProjectCode,
				CheckingItem:  item.ProjectName,
				ModalityType:  objdata.Data.ModalityName,
				Modality:      objdata.Data.DeviceName,
				RemoteRPID:    objdata.Data.RegisterId,
			}
			FlpCheckItems = append(FlpCheckItems, flpitem)
		}
	}
	obj := global.FLPPACSApplyData{
		SiteName:         objdata.Data.HospitalId,
		HospitalName:     hospitalcfg.HospitalName.String,
		PatientID:        objdata.Data.PatientCode,
		LocalName:        objdata.Data.PatientName,
		EnglishName:      objdata.Data.PatientSpellName,
		ReferenceNo:      objdata.Data.IdCardNumber,
		Birthday:         objdata.Data.Birthday,
		Gender:           objdata.Data.PatientSexName,
		Address:          objdata.Data.Address,
		Telephone:        objdata.Data.PhoneNumber,
		RemotePID:        objdata.Data.RequestNumber,
		Marriage:         "",
		AccNo:            objdata.Data.AccessionNumber,
		ApplyDept:        objdata.Data.RequestDepartmentName,
		ApplyDoctor:      objdata.Data.RequestDoctorName,
		StudyInstanceUID: objdata.Data.StudyInstanceUid,
		CardNo:           objdata.Data.MedicareCardNumber,
		InhospitalNo:     objdata.Data.ClinicNumber,
		ClinicNo:         objdata.Data.ClinicNumber,
		PatientType:      objdata.Data.PatientTypeName,
		Observation:      objdata.Data.ClinicalDiagnosis,
		HealthHistory:    objdata.Data.MedicalHistory,
		IsEmergency:      objdata.Data.MergencyStatus,
		BedNo:            objdata.Data.SickbedNumberName,
		CurrentAge:       strconv.Itoa(objdata.Data.Age) + objdata.Data.AgeUnitName,
		Registrar:        objdata.Data.RegisterDoctorCode,
		RegisterDt:       objdata.Data.RegisterTime,
		Technician:       objdata.Data.StudyDoctorCode,
		ExamineDt:        objdata.Data.StudyTime,
		Status:           1,
		CheckItems:       FlpCheckItems,
	}
	SyncFLPPacsApplyData(obj, centerHospital.PacsInterfaceURL.String)
}

// (函数功能F)济宁附属医院远程查看
func SendRemoteViewApplyData(applyid string) {
	global.Logger.Debug("通过接口获取远程查看申请单信息：", applyid)
	// 获取区域PACS申请单信息和报告
	objdata := GetQYPacsApplyReportData(applyid)
	// 获取申请医院的接口信息
	hospitalcfg, err := model.GetHospitalConfig(objdata.Data.HospitalId)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("申请医院信息：", hospitalcfg)

	// 获取上传中心医院的接口信息
	centerHospital, err := model.GetHospitalConfig(objdata.Data.CenterHospitalId)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("中心医院信息：", centerHospital)
	// 发送申请单数据到飞利浦PACS
	global.Logger.Debug("发送申请单数据到飞利浦PACS: ", objdata)
	var FlpCheckItems []global.FLP_CheckItem
	for _, body := range objdata.Data.BodyPartList {
		for _, item := range body.ProjectList {
			flpitem := global.FLP_CheckItem{
				ProcedureCode: item.ProjectCode,
				CheckingItem:  item.ProjectName,
				ModalityType:  objdata.Data.ModalityName,
				Modality:      objdata.Data.DeviceName,
				RemoteRPID:    objdata.Data.RegisterId,
			}
			FlpCheckItems = append(FlpCheckItems, flpitem)
		}
	}
	reportData := global.FLP_ReportData{
		ReportName: objdata.Data.ReportData.ReportName,
		WYSRTF:     objdata.Data.ReportData.WYSRTF,
		WYGRTF:     objdata.Data.ReportData.WYGRTF,
		WYS:        objdata.Data.ReportData.WYS,
		WYG:        objdata.Data.ReportData.WYG,
		IsPositive: objdata.Data.ReportData.IsPositive,
		Creater:    objdata.Data.ReportData.CreateDt,
		CreateDt:   objdata.Data.ReportData.CreateDt,
		Submitter:  objdata.Data.ReportData.Submitter,
		SubmitDt:   objdata.Data.ReportData.SubmitDt,
		Approver:   objdata.Data.ReportData.Approver,
		ApproveDt:  objdata.Data.ReportData.ApproveDt,
	}
	obj := global.FLPPACSApplyData{
		SiteName:         objdata.Data.HospitalId,
		HospitalName:     hospitalcfg.HospitalName.String,
		PatientID:        objdata.Data.PatientCode,
		LocalName:        objdata.Data.PatientName,
		EnglishName:      objdata.Data.PatientSpellName,
		ReferenceNo:      objdata.Data.IdCardNumber,
		Birthday:         objdata.Data.Birthday,
		Gender:           objdata.Data.PatientSexName,
		Address:          objdata.Data.Address,
		Telephone:        objdata.Data.PhoneNumber,
		RemotePID:        objdata.Data.RequestNumber,
		Marriage:         "",
		AccNo:            objdata.Data.AccessionNumber,
		ApplyDept:        objdata.Data.RequestDepartmentName,
		ApplyDoctor:      objdata.Data.RequestDoctorName,
		StudyInstanceUID: objdata.Data.StudyInstanceUid,
		CardNo:           objdata.Data.MedicareCardNumber,
		InhospitalNo:     objdata.Data.ClinicNumber,
		ClinicNo:         objdata.Data.ClinicNumber,
		PatientType:      objdata.Data.PatientTypeName,
		Observation:      objdata.Data.ClinicalDiagnosis,
		HealthHistory:    objdata.Data.MedicalHistory,
		IsEmergency:      objdata.Data.MergencyStatus,
		BedNo:            objdata.Data.SickbedNumberName,
		CurrentAge:       strconv.Itoa(objdata.Data.Age) + objdata.Data.AgeUnitName,
		Registrar:        objdata.Data.RegisterDoctorCode,
		RegisterDt:       objdata.Data.RegisterTime,
		Technician:       objdata.Data.StudyDoctorCode,
		ExamineDt:        objdata.Data.StudyTime,
		Status:           2,
		CheckItems:       FlpCheckItems,
		ReportData:       reportData,
	}
	SyncFLPPacsApplyData(obj, centerHospital.PacsInterfaceURL.String)
}

// 同步申请单到济宁附属医院飞利浦PACS
func SyncFLPPacsApplyData(data global.FLPPACSApplyData, url string) {
	global.Logger.Debug("开始执行同步申请单数据到飞利浦PACS", data)
	reqdata, err := json.Marshal(data)
	if err != nil {
		global.Logger.Error(err)
		return
	}

	global.Logger.Debug("操作的URL: ", url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqdata))
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer resp.Body.Close()
	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Info("resp.Body: ", string(resp_body))
}

// 获取区域PACS申请单数据
func GetQYPacsApplyData(registerid string) (data global.QYPacsApplyData) {
	global.Logger.Debug("开始获取区域PACS申请单数据：", registerid)
	url := global.SystemData.QYPacsInterfaceUrl
	url += "//"
	url += "ms-qypacs-data//v1//register//"
	url += registerid
	global.Logger.Debug("操作的URL: ", url)
	resp, err := http.Get(url)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer resp.Body.Close()
	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Info("resp.Body: ", string(resp_body))

	err = json.Unmarshal(resp_body, &data)
	if err != nil {
		global.Logger.Error("resp.Body err ", err)
		return
	}
	return
}

// 获取区域PACS申请单数据和报告
func GetQYPacsApplyReportData(registerid string) (data global.QYPacsApplyData) {
	global.Logger.Debug("开始获取区域PACS申请单和报告数据：", registerid)
	url := global.SystemData.QYPacsInterfaceUrl
	url += "//"
	url += "ms-qypacs-data//v1//register//report//"
	url += registerid
	global.Logger.Debug("操作的URL: ", url)
	resp, err := http.Get(url)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer resp.Body.Close()
	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Info("resp.Body: ", string(resp_body))

	err = json.Unmarshal(resp_body, &data)
	if err != nil {
		global.Logger.Error("resp.Body err ", err)
		return
	}
	return
}

func parsePasvResponse(response string) (string, int) {
	start, end := strings.Index(response, "("), strings.Index(response, ")")
	if start == -1 || end == -1 {
		return "", 0
	}

	addrParts := strings.Split(response[start+1:end], ",")
	addr := fmt.Sprintf("%s.%s.%s.%s", addrParts[0], addrParts[1], addrParts[2], addrParts[3])
	port := (toInt(addrParts[4]) << 8) + toInt(addrParts[5])

	return addr, port
}

func toInt(s string) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}
