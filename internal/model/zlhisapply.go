package model

// 中联his 数据包

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"database/sql"
	"strconv"
	"strings"
)

// 获取申请单数据
func GetZLHisViewApply(db *sql.DB, sql string) (data []global.ApplyFormResultData) {
	global.Logger.Debug("开始查询视图数据.....")
	rows, err := db.Query(sql)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		key := global.ZLHISApplyInfo{}
		rows.Scan(&key.HisApplyID, &key.HisApplyJLID, &key.PatientName, &key.PatientSpellName, &key.PatientTypeCode, &key.PatientTypeName, &key.MedicalRecordNumber,
			&key.SexCode, &key.SexName, &key.Age, &key.AgeUnit, &key.Birthday, &key.ModalityCode, &key.ProjectCode, &key.ProjectName, &key.ProjectFee, &key.ProjectNote,
			&key.ProjectDetailID, &key.BodypartCode, &key.BodyPart, &key.ProjectCount, &key.ClinicNumber, &key.VisitCardNumber, &key.PhoneNumber, &key.PatientSectionCode,
			&key.PatientSectionName, &key.SickbedNumber, &key.RequestTime, &key.IdCardNumber, &key.Address, &key.ClinicalDiagnosis, &key.MedicalHistory, &key.RequestDepartmentCode,
			&key.RequestDepartmentName, &key.RequestDoctorCode, &key.RequestDoctorName, &key.CheckNote, &key.FilmCount, &key.FilmType, &key.GraphicReport,
			&key.Emergency, &key.IsolationFlag, &key.GreenchanFlag, &key.Fee, &key.RmethodName)

		// 获取性别编码
		var sexCode int
		if key.SexName.String != "" {
			for _, dict := range global.DictDatas {
				if key.SexName.String == dict.Name {
					sexCode = dict.Code
					break
				}
			}
		} else {
			sexCode = global.OtherCode
		}

		// 获取年龄单位编码
		var ageUnitCode int
		if key.AgeUnit.String != "" {
			for _, dict := range global.DictDatas {
				if key.AgeUnit.String == dict.Name {
					ageUnitCode = dict.Code
					break
				}
			}
		} else {
			ageUnitCode = global.OtherCode
		}

		patinfo := global.PatientInfo{
			Pat_name:     key.PatientName.String,
			Pat_sex_code: key.SexCode.String,
			Pat_sex:      sexCode,
			Pat_age:      int(key.Age.Int16),
			Pat_age_unit: ageUnitCode,
			Pat_brsdate:  key.Birthday.String,
			Pat_tel:      key.PhoneNumber.String,
			Pat_idno:     key.IdCardNumber.String,
			Pat_addr:     key.Address.String,
		}

		// 获取就诊类型编码
		var patCode int
		if key.PatientTypeName.String != "" {
			for _, dict := range global.DictDatas {
				if key.PatientTypeName.String == dict.Name {
					patCode = dict.Code
					break
				}
			}
		} else {
			patCode = global.OtherCode
		}

		// 获取检查类型编码
		var studyCode int
		if key.ModalityCode.String != "" {
			for _, dict := range global.DictDatas {
				if key.ModalityCode.String == dict.Name {
					studyCode = dict.Code
					break
				}
			}
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
		var bodysarr []global.CheckBody

		for i := 0; i < len(checkBodysCode) && i < len(checkBodysName) && i < len(checkItemsCode) && i < len(checkItemsCode) && i < len(checkItemsName); i++ {
			var item global.CheckItem
			var body global.CheckBody
			var itemsarr []global.CheckItem
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
				item.Apply_check_item_code = itemcode
				item.Apply_check_item_name = itemname
				item.Apply_points_note = itemnote
				item.Apply_detail_id = itemdetailid
				item.Apply_checkitem_fee = itemfee
				for i := 0; i < len(bodysarr); i++ {
					if (bodysarr[i].Apply_bodypart_code) == bodycode {
						bodysarr[i].Apply_checkItems = append(bodysarr[i].Apply_checkItems, item)
					}
				}

			} else {
				// 不存在
				bodymap[bodycode] = true
				item.Apply_check_item_code = itemcode
				item.Apply_check_item_name = itemname
				item.Apply_points_note = itemnote
				item.Apply_checkitem_fee = itemfee
				item.Apply_detail_id = itemdetailid
				itemsarr = append(itemsarr, item)
				body.Apply_bodypart_code = bodycode
				body.Apply_bodypart_name = bodyname
				body.Apply_checkItems = itemsarr
				bodysarr = append(bodysarr, body)
			}
		}

		applyinfo := global.ApplyInfo{
			Apply_id:                 key.HisApplyID.String,
			Apply_pat_type_code:      key.PatientTypeCode.String,
			Apply_pat_type:           patCode,
			Apply_medical_record:     key.MedicalRecordNumber.String,
			Apply_study_type:         studyCode,
			Apply_checkitems_id:      key.ProjectCode.String,
			Apply_checkitems_name:    key.ProjectName.String,
			Apply_bodyparts_id:       key.BodypartCode.String,
			Apply_bodyparts_name:     key.BodyPart.String,
			Apply_bodys:              bodysarr,
			Apply_clinic_id:          key.ClinicNumber.String,
			Apply_visit_card_no:      key.VisitCardNumber.String,
			Apply_section_id:         key.PatientSectionCode.String,
			Apply_section:            key.PatientSectionName.String,
			Apply_sicked_index:       key.SickbedNumber.String,
			Apply_time:               key.RequestTime.String,
			Apply_details_id:         key.ProjectDetailID.String,
			Apply_clinical_diagnosis: key.ClinicalDiagnosis.String,
			Apply_illness_history:    key.MedicalHistory.String,
			Apply_department_id:      key.RequestDepartmentCode.String,
			Apply_department:         key.RequestDepartmentName.String,
			Apply_doctor_id:          key.RequestDoctorCode.String,
			Apply_doctor:             key.RequestDoctorName.String,
			Apply_check_note:         key.CheckNote.String,
			Apply_film_count:         int(key.FilmCount.Int16),
			Apply_film_flag:          int(key.FilmType.Int16),
			Apply_report_flag:        int(key.GraphicReport.Int16),
			Apply_isolation_flag:     int(key.IsolationFlag.Int16),
			Apply_greenchan_flag:     int(key.GreenchanFlag.Int16),
			Apply_mergency_status:    int(key.Emergency.Int16),
			Apply_fee:                key.Fee.String,
		}
		obj := global.ApplyFormResultData{
			Apply_Info: applyinfo,
			Pat_Info:   patinfo,
		}
		data = append(data, obj)
	}
	return
}

// 获取数据总数
func GetDataCount(db *sql.DB, sql string) (count int) {
	global.Logger.Debug("开始查询条件数据总数.....")
	rows, err := db.Query(sql)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		count++
	}
	return
}
