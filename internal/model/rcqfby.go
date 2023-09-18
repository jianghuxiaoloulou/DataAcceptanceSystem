package model

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"database/sql"
	"strconv"
	"strings"
)

// 任城区妇保院申请单数据
func GetRcqfbyApplyData(db *sql.DB, sql, hospitalid string) (data []global.RcqfbtApplyData) {
	global.Logger.Debug("开始查询视图数据.....")
	rows, err := db.Query(sql)
	if err != nil {
		global.Logger.Error("QUery err: ", err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		key := global.RcqfbtApplyInfo{}
		err = rows.Scan(&key.HisApplyID, &key.HisApplyJLID, &key.PatientName, &key.PatientSpellName, &key.PatientTypeCode,
			&key.PatientTypeName, &key.MedicalRecordNumber, &key.SexCode, &key.SexName, &key.Age, &key.AgeUnit, &key.Birthday,
			&key.ModalityCode, &key.ProjectCode, &key.ProjectName, &key.ProjectFee, &key.ProjectNote, &key.ProjectDetailID,
			&key.BodypartCode, &key.BodyPart, &key.ProjectCount, &key.ClinicNumber, &key.VisitCardNumber, &key.PhoneNumber,
			&key.PatientSectionCode, &key.PatientSectionName, &key.SickbedNumber, &key.RequestTime, &key.IdCardNumber,
			&key.Address, &key.ClinicalDiagnosis, &key.MedicalHistory, &key.RequestDepartmentCode, &key.RequestDepartmentName,
			&key.RequestDoctorCode, &key.RequestDoctorName, &key.CheckNote, &key.FilmCount, &key.FilmType, &key.GraphicReport,
			&key.Emergency, &key.IsolationFlag, &key.GreenchanFlag, &key.Fee, &key.RmethodName, &key.AccessionNumber,
			&key.PatientCode, &key.HisPatientId, &key.RegisterStatus, &key.RegisterDoctorId, &key.RegisterDoctorCode,
			&key.RegisterDoctorName, &key.RegisterTime, &key.QueueNumber, &key.DeviceId, &key.DeviceCode, &key.DeviceName,
			&key.StudyDoctorId, &key.StudyDoctorCode, &key.StudyDoctorName, &key.AssistDoctorId, &key.AssistDoctorCode,
			&key.AssistDoctorName, &key.OperationDoctorId, &key.OperationDoctorCode, &key.OperationDoctorName)
		if err != nil {
			global.Logger.Error("Scan err: ", err.Error())
			return
		}
		global.Logger.Debug("查询到的数据：", key)
		// 获取性别编码
		var sexCode int
		if key.SexName.String != "" {
			sexCode = GetDictCode(key.SexName.String)
		} else {
			sexCode = global.OtherCode
		}

		// 获取年龄单位编码
		var ageUnitCode int
		if key.AgeUnit.String != "" {
			ageUnitCode = GetDictCode(key.AgeUnit.String)
		} else {
			ageUnitCode = global.OtherCode
		}

		// 获取就诊类型编码
		var patCode int
		if key.PatientTypeName.String != "" {
			patCode = GetDictCode(key.PatientTypeName.String)
		} else {
			patCode = global.OtherCode
		}

		// 获取检查类型编码
		var studyCode int
		if key.ModalityCode.String != "" {
			studyCode = GetDictCode(key.ModalityCode.String)
		} else {
			studyCode = global.OtherCode
		}

		// 检查项目检查部位处理 分隔符|*|
		checkBodysCode := strings.Split(key.BodypartCode.String, "|*|")
		checkBodysName := strings.Split(key.BodyPart.String, "|*|")
		checkItemsCode := strings.Split(key.ProjectCode.String, "|*|")
		checkItemsName := strings.Split(key.ProjectName.String, "|*|")
		checkItemsNote := strings.Split(key.ProjectNote.String, "|*|")
		checkItemsfee := strings.Split(key.ProjectFee.String, "|*|")
		requestdetailid := strings.Split(key.ProjectDetailID.String, "|*|")

		var bodymap = make(map[string]bool)
		var bodysarr []global.BodyPartStr

		for i := 0; i < len(checkBodysCode) && i < len(checkBodysName) && i < len(checkItemsCode) && i < len(checkItemsCode) && i < len(checkItemsName); i++ {
			var item global.ProjectItemStr
			var body global.BodyPartStr
			var itemsarr []global.ProjectItemStr
			bodycode := checkBodysCode[i]
			bodyname := checkBodysName[i]
			itemcode := checkItemsCode[i]
			itemname := checkItemsName[i]
			var itemnote string
			var itemdetailid string
			var itemfee int
			if i < len(checkItemsNote) {
				itemnote = checkItemsNote[i]
			}

			if i < len(checkItemsfee) {
				itemfee, _ = strconv.Atoi(checkItemsfee[i])
			}
			if i < len(requestdetailid) {
				itemdetailid = requestdetailid[i]
			}

			if bodymap[bodycode] {
				// 存在
				item.ProjectCode = itemcode
				item.ProjectName = itemname
				item.ProjectNote = itemnote
				item.RequestDetailId = itemdetailid
				item.Fee = float64(itemfee)
				for i := 0; i < len(bodysarr); i++ {
					if (bodysarr[i].BodyPartCode) == bodycode {
						bodysarr[i].ProjectList = append(bodysarr[i].ProjectList, item)
					}
				}

			} else {
				// 不存在
				bodymap[bodycode] = true
				item.ProjectCode = itemcode
				item.ProjectName = itemname
				item.ProjectNote = itemnote
				item.Fee = float64(itemfee)
				item.RequestDetailId = itemdetailid
				itemsarr = append(itemsarr, item)
				body.BodyPartCode = bodycode
				body.BodyPartName = bodyname
				body.ProjectList = itemsarr
				bodysarr = append(bodysarr, body)
			}
		}

		applyinfo := global.RcqfbtApplyData{
			HospitalId:            hospitalid,
			PatientCode:           key.PatientCode.String,
			HisPatientId:          key.HisPatientId.String,
			PatientName:           key.PatientName.String,
			PatientSpellName:      key.PatientSpellName.String,
			PatientSexCode:        sexCode,
			PatientSexName:        key.SexName.String,
			Age:                   int(key.Age.Int16),
			AgeUnitCode:           ageUnitCode,
			AgeUnitName:           key.AgeUnit.String,
			Birthday:              key.Birthday.String,
			IdCardNumber:          key.IdCardNumber.String,
			MedicareCardNumber:    key.VisitCardNumber.String,
			PhoneNumber:           key.PhoneNumber.String,
			Address:               key.Address.String,
			ChiefComplaint:        "",
			ClinicalManifestation: "",
			ClinicalDiagnosis:     key.ClinicalDiagnosis.String,
			MedicalHistory:        key.MedicalHistory.String,
			CheckMemo:             key.CheckNote.String,
			MergencyStatus:        int(key.Emergency.Int16),
			RequestId:             key.HisApplyJLID.String,
			RequestNumber:         key.HisApplyID.String,
			RequestTime:           key.RequestTime.String,
			ClinicNumber:          key.ClinicNumber.String,
			MedicalRecordNumber:   key.MedicalRecordNumber.String,
			PatientTypeCode:       patCode,
			PatientTypeName:       key.PatientTypeName.String,
			ModalityCode:          studyCode,
			ModalityName:          key.ModalityCode.String,
			RequestDoctorCode:     key.RequestDoctorCode.String,
			RequestDoctorName:     key.RequestDoctorName.String,
			RequestDepartmentCode: key.RequestDepartmentCode.String,
			RequestDepartmentName: key.RequestDepartmentName.String,
			PatientSectionCode:    key.PatientSectionCode.String,
			PatientSectionName:    key.PatientSectionName.String,
			SickbedNumberCode:     "",
			SickbedNumberName:     key.SickbedNumber.String,
			GraphicReport:         int(key.GraphicReport.Int16),
			FilmType:              int(key.FilmType.Int16),
			FilmNum:               int(key.FilmCount.Int16),
			IsolationSign:         strconv.Itoa(int(key.IsolationFlag.Int16)),
			IsGreenChannel:        int(key.GreenchanFlag.Int16),
			AccessionNumber:       key.AccessionNumber.String,
			RegisterStatus:        int(key.RegisterStatus.Int16),
			RegisterDoctorId:      key.RegisterDoctorId.String,
			RegisterDoctorCode:    key.RegisterDoctorCode.String,
			RegisterDoctorName:    key.RegisterDoctorName.String,
			RegisterTime:          key.RegisterTime.String,
			QueueNumber:           key.QueueNumber.String,
			DeviceId:              key.DeviceId.String,
			DeviceCode:            key.DeviceCode.String,
			TotalFee:              key.Fee.Float64,
			StudyDoctorId:         key.StudyDoctorId.String,
			StudyDoctorCode:       key.StudyDoctorCode.String,
			StudyDoctorName:       key.StudyDoctorName.String,
			AssistDoctorId:        key.AssistDoctorId.String,
			AssistDoctorCode:      key.AssistDoctorCode.String,
			AssistDoctorName:      key.AssistDoctorName.String,
			OperationDoctorId:     key.OperationDoctorId.String,
			OperationDoctorCode:   key.OperationDoctorCode.String,
			OperationDoctorName:   key.OperationDoctorName.String,
			BodyPartList:          bodysarr,
		}
		data = append(data, applyinfo)
	}
	return
}

// 任城区妇保院DICOM数据
func GetRcqfbyDicomData(db *sql.DB, sql, hospitalid string) {
	global.Logger.Debug("开始查询DICOM视图数据.....")
	rows, err := db.Query(sql)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		key := global.DicomDB{}
		rows.Scan(&key.AccessionNumber, &key.StudyInstanceUid, &key.SeriesInstanceUid, &key.SopInstanceUid,
			&key.FileType, &key.DicomFileName, &key.Host, &key.Port, &key.User, &key.Password, &key.UpdateTime)

		data := global.DicomInfo{
			HospitalID:        hospitalid,
			AccessionNumber:   key.AccessionNumber.String,
			StudyInstanceUid:  key.StudyInstanceUid.String,
			SeriesInstanceUid: key.SeriesInstanceUid.String,
			SopInstanceUid:    key.SopInstanceUid.String,
			FileType:          int(key.FileType.Int16),
			DicomFileName:     key.DicomFileName.String,
			Host:              key.Host.String,
			Port:              key.Port.String,
			User:              key.User.String,
			Password:          key.Password.String,
			UpdateTime:        key.UpdateTime.String,
		}
		// 工厂模式单独处理DICOM影像上传
		global.Logger.Debug("需要处理的DICOM 数据是：", data)
		global.DicomDataChan <- data
	}
}

// 更新医院数据上传的时间
func UpdateUploadTiem(time, hospitalid string) {
	global.Logger.Debug("更新申请单数据上传时间.")
	sql := `update sys_dict_hospital_config set upload_time = ? where hospital_id = ?;`
	err := global.DBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.DBEngine.Close()
		global.DBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	global.DBEngine.Exec(sql, time, hospitalid)
}
