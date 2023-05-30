package model

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"database/sql"
	"strconv"
)

// 华卓his数据包
// 住院申请单 v_hsbdi_requestinfo
type ZYApplyInfo struct {
	SQJLID sql.NullInt16  // 申请记录序号
	SQMXID sql.NullInt16  // 申请明细序号
	BRJZHM sql.NullString // 住院号
	BRZYID sql.NullInt16  // 住院唯一号
	BRSQHM sql.NullString // 病人申请号码
	BRDAXM sql.NullString // 姓名
	BRDAXB sql.NullString // 性别 1男2女
	BRCSRQ sql.NullTime   // 出生日期
	BRJZNL sql.NullInt16  // 病人接诊年龄
	BRNLDW sql.NullString // 病人年龄单位
	DQKSID sql.NullString // 科别
	DQBQID sql.NullString // 病区id
	DQBQMC sql.NullString // 病区名称
	BQCWHM sql.NullString // 床号
	JBZDMC sql.NullString // 临床诊断
	BRSQLX sql.NullString // 申请类型 1 检验 2 检查
	KDKSID sql.NullString // 申请科室
	SQKDRQ sql.NullTime   // 申请时间
	JCXMID sql.NullInt16  // 项目代码
	JCXMMC sql.NullString // 项目名称
	XMDYMC sql.NullString // 项目打印名称
	JCXMLX sql.NullString // 检查项目类型 检查用
	XMSQSL sql.NullInt16  // 数量
	SQZTBZ sql.NullInt16  // 申请状态 3 已退费 -1 撤销 3 确认 6 计费
	ZXKSID sql.NullString // 检验科室
	SFJZPB sql.NullInt16  // 是否急诊 1 急诊 0 否
	FYHJJE sql.NullInt16  // 检验项目金额
	ZLBWMC sql.NullString // 诊疗部位名称
	ZLBWID sql.NullInt16  // 诊疗部位ID
	BRDABH sql.NullString // 病人档案编号
	BRLYMC sql.NullString // 病人类型名称 VIP会员
	KDKSMC sql.NullString // 开单科室名称
	KDYSXM sql.NullString // 申请医生名称
	KDYSID sql.NullString // 申请医生ID
	BRTGJC sql.NullString // 体征
	BRZSNR sql.NullString // 症状
	BRXBS  sql.NullString // 病人现病史
	BRLXDZ sql.NullString // 病人联系地址
	BRSFZH sql.NullString // 病人身份证号
	BRLXDH sql.NullString // 病人联系电话
	XMFYNR sql.NullString // 项目内容
	BRGLXX sql.NullString // 隔离
	SFTWBG sql.NullInt16  // 图文报告标志 0否1是
	JPSYFS sql.NullString // 胶片标志
	BRZYSX sql.NullString // 注意事项
	BRSYZZ sql.NullString // 适应症
	BRZSY  sql.NullString // 病人病案号
}

// 门诊申请单 v_hsbci_requestinfo
type MZApplyInfo struct {
	SQJLID sql.NullInt16  // 申请记录序号
	SQMXID sql.NullInt16  // 申请明细序号
	BRSQHM sql.NullString // 申请号码
	BRJZHM sql.NullString // 住院(门诊)号
	BRJZXH sql.NullInt16  // 病人接诊序号（唯一号）
	BRDAXM sql.NullString // 姓名
	BRDAXB sql.NullString // 性别 1男2女
	BRCSRQ sql.NullTime   // 出生日期
	BRJZNL sql.NullInt16  // 病人接诊年龄
	BRNLDW sql.NullString // 病人年龄单位
	KDKSMC sql.NullString // 开单科室名称
	JBZDMC sql.NullString // 临床诊断
	BRSQLX sql.NullString // 申请类型 1 检验 2 检查
	KDKSID sql.NullInt16  // 开单科室ID
	JCXMLX sql.NullString // 检查项目类型 检查用
	KDYSID sql.NullString // 申请医生
	KDKSXM sql.NullString // 申请科室
	SQKDRQ sql.NullTime   // 申请时间
	BRTGJC sql.NullString // 体征
	BRZSNR sql.NullString // 症状
	BRXBS  sql.NullString // 病人现病史
	ZLJFJC sql.NullString // 空字段
	SFTWBG sql.NullInt16  // 图文报告标志 0否1是
	JPSYFS sql.NullString // 胶片标志
	BRLYMC sql.NullString // 病人类型名称 VIP会员
	JCXMID sql.NullInt16  // 项目代码
	JCXMMC sql.NullString // 项目名称
	XMSQSL sql.NullInt16  // 数量
	SQZTBZ sql.NullInt16  // 申请状态 1执行 0未执行 3 已退费 -1 撤销 3 确认 6 计费
	ZXKSID sql.NullString // 检验科室
	SFJZPB sql.NullInt16  // 是否急诊 1 急诊 0 否
	FYHJJE sql.NullInt16  // 检验项目金额
	ZLBWID sql.NullInt16  // 诊疗部位ID
	ZLBWMC sql.NullString // 诊疗部位名称
	BRDABH sql.NullString // 病人档案编号
	BRLXDZ sql.NullString // 病人联系地址
	BRSFZH sql.NullString // 病人身份证号
	BRLXDH sql.NullString // 病人联系电话
	XMFYNR sql.NullString // 项目内容
	ZLJFMC sql.NullString // 检查机房
	BRZYSX sql.NullString // 注意事项
	BRSYZZ sql.NullString // 适应症
	BRZSY  sql.NullString // 病人病案号
}

// 体检申请单 V_HSBTJ_REQUESTINFO（oracle版本）
type TJApplyInfoOracle struct {
	sqjlid     sql.NullString // 体检号
	tijianbm   sql.NullString // 检查编号
	xingming   sql.NullString // 姓名
	xingbie    sql.NullString // 性别（1=男，2=女，%=不详）
	ylxh       sql.NullString // 组合项目编号(多个项目逗号隔开，英文逗号)
	zuhexmmc   sql.NullString // 组合项目名称(多个项目逗号隔开，英文逗号)
	sfzh       sql.NullString // 身份证号
	jianchalx  sql.NullString // 检查类型(CT/US/MR/DR/MR/ES/MG/RF.....)
	TIJIANRQ   sql.NullTime   // 开单日期
	chushengrq sql.NullTime   // 出生日期
	lianxidh   sql.NullString // 联系电话
	sqmxid     sql.NullString // 为空
	danjia     sql.NullString // 检查项目单价，如‘0’
}

// 体检申请单 VIEW_TJPACSJK（sqlserver版本）
type TJApplyInfoSqlServer struct {
	SQDH        sql.NullString // 体检号
	TJBH        sql.NullString // 检查编号
	XM          sql.NullString // 姓名
	XB          sql.NullString // 性别（M=男，F=女）
	XMBH        sql.NullString // 组合项目编号(多个项目逗号隔开，英文逗号)
	XMMC        sql.NullString // 组合项目名称(多个项目逗号隔开，英文逗号)
	studyclass0 sql.NullString // 检查类型(CT/US/MR/DR/MR/ES/MG/RF.....)
	DJRQ        sql.NullTime   // 开单日期
	CSNY        sql.NullTime   // 出生日期
	DABH        sql.NullString // 病历号 可空
}

// 获取门诊申请单数据
func GetMZViewApply(param global.Param, param2 []global.Param2) (count int, data []global.ApplyFormResultData) {
	global.Logger.Debug("通过门诊号获取申请单数.....")
	var err error
	sql := `select "SQJLID","SQMXID","BRSQHM","BRJZHM","BRJZXH","BRDAXM","BRDAXB","BRCSRQ","BRJZNL",
	"BRNLDW","KDKSMC","JBZDMC","BRSQLX","KDKSID","JCXMLX","KDYSID","KDKSXM","SQKDRQ","BRTGJC","BRZSNR",
	"BRXBS","ZLJFJC","SFTWBG","JPSYFS","BRLYMC","JCXMID","JCXMMC","XMSQSL","SQZTBZ","ZXKSID","SFJZPB",
	"FYHJJE","ZLBWID","ZLBWMC","BRDABH","BRLXDZ","BRSFZH","BRLXDH","XMFYNR","ZLJFMC","BRZYSX","BRSYZZ","BRZSY" 
	from v_hsbci_requestinfo where 1 = 1`

	switch param.ParamType {
	case global.Apply_Param_JZKH:
		sql += " and \"BRZSY\" = " + param.ParamValue
	case global.Apply_Param_MZH:
		sql += " and \"BRJZHM\" = " + param.ParamValue
	case global.Apply_Param_MZSQDH:
		sql += " and \"BRSQHM\" = " + param.ParamValue
	case global.Apply_Param_SFZH:
		sql += " and \"BRSFZH\" = " + param.ParamValue
	default:
		sql += " and 1 = 1"
	}

	// 参数2
	param2len := len(param2)
	var param2str string
	for i := 0; i < param2len; i++ {
		if i > 0 {
			param2str += " or "
		}
		if (param2[i].StartDate != "") && (param2[i].EndDate != "") {
			param2str += " \"SQKDRQ\" between to_date('" + param2[i].StartDate + "','yyyy-mm-dd hh24:mi:ss') and to_date('" + param2[i].EndDate + "','yyyy-mm-dd hh24:mi:ss')"
		} else {
			param2str += "1=1"
		}
	}

	sql += " and ("
	sql += param2str
	sql += ")"

	global.Logger.Debug("执行的sql语句是: ", sql)

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
		key := MZApplyInfo{}
		rows.Scan(&key.SQJLID, &key.SQMXID, &key.BRSQHM, &key.BRJZHM, &key.BRJZXH, &key.BRDAXM, &key.BRDAXB, &key.BRCSRQ, &key.BRJZNL,
			&key.BRNLDW, &key.KDKSMC, &key.JBZDMC, &key.BRSQLX, &key.KDKSID, &key.JCXMLX, &key.KDYSID, &key.KDKSXM, &key.SQKDRQ, &key.BRTGJC, &key.BRZSNR,
			&key.BRXBS, &key.ZLJFJC, &key.SFTWBG, &key.JPSYFS, &key.BRLYMC, &key.JCXMID, &key.JCXMMC, &key.XMSQSL, &key.SQZTBZ, &key.ZXKSID, &key.SFJZPB,
			&key.FYHJJE, &key.ZLBWID, &key.ZLBWMC, &key.BRDABH, &key.BRLXDZ, &key.BRSFZH, &key.BRLXDH, &key.XMFYNR, &key.ZLJFMC, &key.BRZYSX, &key.BRSYZZ, &key.BRZSY)

		var sex int
		if key.BRDAXB.String == "男" {
			sex = 1
		} else if key.BRDAXB.String == "女" {
			sex = 2
		} else {
			sex = 9
		}

		var ageUnit int
		if key.BRNLDW.String == "岁" {
			ageUnit = 1
		} else if key.BRNLDW.String == "月" {
			ageUnit = 2
		} else if key.BRNLDW.String == "周" {
			ageUnit = 3
		} else if key.BRNLDW.String == "天" {
			ageUnit = 4
		} else if key.BRNLDW.String == "时" {
			ageUnit = 5
		} else if key.BRNLDW.String == "分" {
			ageUnit = 6
		} else if key.BRNLDW.String == "秒" {
			ageUnit = 7
		}

		patinfo := global.PatientInfo{
			Pat_id:         strconv.Itoa(int(key.BRJZXH.Int16)),
			Pat_idno:       key.BRSFZH.String,
			Pat_si_no:      "",
			Pat_name:       key.BRDAXM.String,
			Pat_spell_name: "",
			Pat_sex_code:   "",
			Pat_sex:        sex,
			Pat_brsdate:    key.BRCSRQ.Time.Format("2006-10-21"),
			Pat_age:        int(key.BRJZNL.Int16),
			Pat_age_unit:   ageUnit,
			Pat_weight:     0,
			Pat_mari_code:  0,
			Pat_tel:        key.BRLXDH.String,
			Pat_addr:       key.BRLXDZ.String,
		}

		var studytype int
		if key.JCXMLX.String == "CT" {
			studytype = 3
		} else if key.JCXMLX.String == "DR" {
			studytype = 2
		} else if key.JCXMLX.String == "MR" {
			studytype = 4
		} else if key.JCXMLX.String == "DSA" {
			studytype = 5
		} else if key.JCXMLX.String == "US" {
			studytype = 6
		} else if key.JCXMLX.String == "ES" {
			studytype = 7
		} else if key.JCXMLX.String == "PA" {
			studytype = 8
		} else if key.JCXMLX.String == "NM" {
			studytype = 9
		} else if key.JCXMLX.String == "PET" {
			studytype = 10
		} else {
			studytype = 99
		}

		film_flag, _ := strconv.Atoi(key.JPSYFS.String)
		var vip int
		if key.BRLYMC.String != "" {
			vip = 1
		} else {
			vip = 0
		}

		applyinfo := global.ApplyInfo{
			Apply_hospital_id:            "",
			Apply_id:                     key.BRSQHM.String,
			Apply_status:                 int(key.SQZTBZ.Int16),
			Apply_time:                   key.SQKDRQ.Time.Format("2006-05-02 11:22:33"),
			Apply_department_id:          "",
			Apply_department:             key.KDKSXM.String,
			Apply_doctor_id:              "",
			Apply_doctor:                 key.KDYSID.String,
			Apply_pat_type_code:          "OP",
			Apply_pat_type:               1,
			Apply_clinic_id:              key.BRJZHM.String,
			Apply_visit_card_no:          "",
			Apply_medical_record:         key.BRZSY.String,
			Apply_pat_body_sign:          key.BRTGJC.String,
			Apply_pat_symptoms:           key.BRZSNR.String,
			Apply_indications:            key.BRSYZZ.String,
			Apply_clinical_manifestation: "",
			Apply_clinical_diagnosis:     key.JBZDMC.String,
			Apply_chief_complaint:        "",
			Apply_illness_history:        key.BRXBS.String,
			Apply_report_flag:            int(key.SFTWBG.Int16),
			Apply_film_count:             int(key.XMSQSL.Int16),
			Apply_film_flag:              film_flag,
			Apply_fee:                    strconv.Itoa(int(key.FYHJJE.Int16)),
			Apply_vip_flag:               vip,
			Apply_isolation_flag:         0,
			Apply_mergency_status:        int(key.SFJZPB.Int16),
			Apply_study_type:             studytype,
			Apply_bodyparts_id:           strconv.Itoa(int(key.ZLBWID.Int16)),
			Apply_bodyparts_name:         key.ZLBWMC.String,
			Apply_checkitems_id:          strconv.Itoa(int(key.JCXMID.Int16)),
			Apply_checkitems_name:        key.JCXMMC.String,
			Apply_details_id:             key.JCXMMC.String,
			Apply_check_note:             key.BRZYSX.String,
			Apply_bodys:                  nil,
			Apply_check_room:             key.KDKSXM.String,
			Apply_section_id:             "",
			Apply_section:                "",
			Apply_sicked_index:           "",
		}
		obj := global.ApplyFormResultData{
			Apply_Info: applyinfo,
			Pat_Info:   patinfo,
		}
		data = append(data, obj)
	}
	return
}

// 获取住院申请单数据
func GetZYViewApply(param global.Param, param2 []global.Param2) (count int, data []global.ApplyFormResultData) {
	global.Logger.Debug("通过门诊号获取申请单数.....")
	var err error
	sql := `select "SQJLID","SQMXID","BRJZHM","BRZYID","BRSQHM","BRDAXM","BRDAXB","BRCSRQ","BRJZNL",
	"BRNLDW","DQKSID","DQBQID","DQBQMC","BQCWHM","JBZDMC","BRSQLX","KDKSID","SQKDRQ","JCXMID","JCXMMC",
	"XMDYMC","JCXMLX","XMSQSL","SQZTBZ","ZXKSID","SFJZPB","FYHJJE","ZLBWMC","ZLBWID","BRDABH","BRLYMC",
	"KDKSMC","KDYSXM","KDYSID","BRTGJC","BRZSNR","BRXBS","BRLXDZ","BRSFZH","BRLXDH","XMFYNR","BRGLXX",
	"SFTWBG","JPSYFS","BRZYSX","BRSYZZ","BRZSY" 
	from v_hsbdi_requestinfo where 1 = 1`

	switch param.ParamType {
	case global.Apply_Param_JZKH:
		sql += " and \"BRZSY\" = " + param.ParamValue
	case global.Apply_Param_ZYH:
		sql += " and \"BRJZHM\" = " + param.ParamValue
	case global.Apply_Param_ZYSQDH:
		sql += " and \"BRSQHM\" = " + param.ParamValue
	case global.Apply_Param_SFZH:
		sql += " and \"BRSFZH\" = " + param.ParamValue
	default:
		sql += " and 1 = 1"
	}

	// 参数2
	param2len := len(param2)
	var param2str string
	for i := 0; i < param2len; i++ {
		if i > 0 {
			param2str += " or "
		}
		if (param2[i].StartDate != "") && (param2[i].EndDate != "") {
			param2str += " \"SQKDRQ\" between to_date('" + param2[i].StartDate + "','yyyy-mm-dd hh24:mi:ss') and to_date('" + param2[i].EndDate + "','yyyy-mm-dd hh24:mi:ss')"
		} else {
			param2str += "1=1"
		}
	}

	sql += " and ("
	sql += param2str
	sql += ")"

	global.Logger.Debug("执行的sql语句是: ", sql)

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
		key := ZYApplyInfo{}
		rows.Scan(&key.SQJLID, &key.SQMXID, &key.BRJZHM, &key.BRZYID, &key.BRSQHM, &key.BRDAXM, &key.BRDAXB, &key.BRCSRQ, &key.BRJZNL,
			&key.BRNLDW, &key.DQKSID, &key.DQBQID, &key.DQBQMC, &key.BQCWHM, &key.JBZDMC, &key.BRSQLX, &key.KDKSID, &key.SQKDRQ, &key.JCXMID, &key.JCXMMC,
			&key.XMDYMC, &key.JCXMLX, &key.XMSQSL, &key.SQZTBZ, &key.ZXKSID, &key.SFJZPB, &key.FYHJJE, &key.ZLBWMC, &key.ZLBWID, &key.BRDABH, &key.BRLYMC,
			&key.KDKSMC, &key.KDYSXM, &key.KDYSID, &key.BRTGJC, &key.BRZSNR, &key.BRXBS, &key.BRLXDZ, &key.BRSFZH, &key.BRLXDH, &key.XMFYNR, &key.BRGLXX,
			&key.SFTWBG, key.JPSYFS, key.BRZYSX, key.BRSYZZ, key.BRZSY)
		var sex int
		if key.BRDAXB.String == "男" {
			sex = 1
		} else if key.BRDAXB.String == "女" {
			sex = 2
		} else {
			sex = 9
		}

		var ageUnit int
		if key.BRNLDW.String == "岁" {
			ageUnit = 1
		} else if key.BRNLDW.String == "月" {
			ageUnit = 2
		} else if key.BRNLDW.String == "周" {
			ageUnit = 3
		} else if key.BRNLDW.String == "天" {
			ageUnit = 4
		} else if key.BRNLDW.String == "时" {
			ageUnit = 5
		} else if key.BRNLDW.String == "分" {
			ageUnit = 6
		} else if key.BRNLDW.String == "秒" {
			ageUnit = 7
		}

		patinfo := global.PatientInfo{
			Pat_id:         strconv.Itoa(int(key.BRZYID.Int16)),
			Pat_idno:       key.BRSFZH.String,
			Pat_si_no:      "",
			Pat_name:       key.BRDAXM.String,
			Pat_spell_name: "",
			Pat_sex_code:   "",
			Pat_sex:        sex,
			Pat_brsdate:    key.BRCSRQ.Time.Format("2006-10-21"),
			Pat_age:        int(key.BRJZNL.Int16),
			Pat_age_unit:   ageUnit,
			Pat_weight:     0,
			Pat_mari_code:  0,
			Pat_tel:        key.BRLXDH.String,
			Pat_addr:       key.BRLXDZ.String,
		}

		var studytype int
		if key.JCXMLX.String == "CT" {
			studytype = 3
		} else if key.JCXMLX.String == "DR" {
			studytype = 2
		} else if key.JCXMLX.String == "MR" {
			studytype = 4
		} else if key.JCXMLX.String == "DSA" {
			studytype = 5
		} else if key.JCXMLX.String == "US" {
			studytype = 6
		} else if key.JCXMLX.String == "ES" {
			studytype = 7
		} else if key.JCXMLX.String == "PA" {
			studytype = 8
		} else if key.JCXMLX.String == "NM" {
			studytype = 9
		} else if key.JCXMLX.String == "PET" {
			studytype = 10
		} else {
			studytype = 99
		}

		film_flag, _ := strconv.Atoi(key.JPSYFS.String)
		var vip int
		if key.BRLYMC.String != "" {
			vip = 1
		} else {
			vip = 0
		}

		var brglxx int
		if key.BRGLXX.String != "" {
			brglxx = 1
		} else {
			brglxx = 0
		}

		applyinfo := global.ApplyInfo{
			Apply_hospital_id:            "",
			Apply_id:                     key.BRSQHM.String,
			Apply_status:                 int(key.SQZTBZ.Int16),
			Apply_time:                   key.SQKDRQ.Time.Format("2006-05-02 11:22:33"),
			Apply_department_id:          "",
			Apply_department:             key.KDKSID.String,
			Apply_doctor_id:              "",
			Apply_doctor:                 key.KDYSID.String,
			Apply_pat_type_code:          "IH",
			Apply_pat_type:               2,
			Apply_clinic_id:              key.BRJZHM.String,
			Apply_visit_card_no:          key.BRZSY.String,
			Apply_medical_record:         key.BRZSY.String,
			Apply_pat_body_sign:          key.BRTGJC.String,
			Apply_pat_symptoms:           key.BRZSNR.String,
			Apply_indications:            key.BRSYZZ.String,
			Apply_clinical_manifestation: "",
			Apply_clinical_diagnosis:     key.JBZDMC.String,
			Apply_chief_complaint:        "",
			Apply_illness_history:        key.BRXBS.String,
			Apply_report_flag:            int(key.SFTWBG.Int16),
			Apply_film_count:             int(key.XMSQSL.Int16),
			Apply_film_flag:              film_flag,
			Apply_fee:                    strconv.Itoa(int(key.FYHJJE.Int16)),
			Apply_vip_flag:               vip,
			Apply_isolation_flag:         brglxx,
			Apply_mergency_status:        int(key.SFJZPB.Int16),
			Apply_study_type:             studytype,
			Apply_bodyparts_id:           strconv.Itoa(int(key.ZLBWID.Int16)),
			Apply_bodyparts_name:         key.ZLBWMC.String,
			Apply_checkitems_id:          strconv.Itoa(int(key.JCXMID.Int16)),
			Apply_checkitems_name:        key.JCXMMC.String,
			Apply_details_id:             key.JCXMMC.String,
			Apply_check_note:             key.BRZYSX.String,
			Apply_bodys:                  nil,
			Apply_check_room:             key.ZXKSID.String,
			Apply_section_id:             key.DQBQID.String,
			Apply_section:                key.DQBQMC.String,
			Apply_sicked_index:           key.BQCWHM.String,
		}
		obj := global.ApplyFormResultData{
			Apply_Info: applyinfo,
			Pat_Info:   patinfo,
		}
		data = append(data, obj)
	}
	return
}

// 获取体检申请单数据
func GetTJViewApply(param global.Param, param2 []global.Param2) (count int, data []global.ApplyFormResultData) {
	global.Logger.Debug("通过门诊号获取申请单数.....")
	var err error
	sql := `select "sqjlid","tijianbm","xingming","xingbie","ylxh","zuhexmmc","sfzh","jianchalx","TIJIANRQ",
	"chushengrq","lianxidh","sqmxid","danjia" 
	from V_HSBTJ_REQUESTINFO where 1 = 1`

	switch param.ParamType {
	case global.Apply_Param_TJH:
		sql += " and \"sqjlid\" = " + param.ParamValue
	case global.Apply_Param_TJSQDH:
		sql += " and \"tijianbm\" = " + param.ParamValue
	case global.Apply_Param_SFZH:
		sql += " and \"sfzh\" = " + param.ParamValue
	default:
		sql += " and 1 = 1"
	}

	// 参数2
	param2len := len(param2)
	var param2str string
	for i := 0; i < param2len; i++ {
		if i > 0 {
			param2str += " or "
		}
		if (param2[i].StartDate != "") && (param2[i].EndDate != "") {
			param2str += " \"TIJIANRQ\" between to_date('" + param2[i].StartDate + "','yyyy-mm-dd hh24:mi:ss') and to_date('" + param2[i].EndDate + "','yyyy-mm-dd hh24:mi:ss')"
		} else {
			param2str += "1=1"
		}
	}

	sql += " and ("
	sql += param2str
	sql += ")"

	global.Logger.Debug("执行的sql语句是: ", sql)

	err = global.OracleDBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.OracleDBEngine, _ = NewOracleDBEngine(global.DatabaseSetting)
	}
	rows, err := global.OracleDBEngine.Query(sql, param.ParamValue)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		key := TJApplyInfoOracle{}
		rows.Scan(&key.sqjlid, &key.tijianbm, &key.xingming, &key.xingbie, &key.ylxh, &key.zuhexmmc, &key.sfzh, &key.jianchalx, &key.TIJIANRQ,
			&key.chushengrq, &key.lianxidh, &key.sqmxid, &key.danjia)
		var sex int
		if key.xingming.String == "男" {
			sex = 1
		} else if key.xingming.String == "女" {
			sex = 2
		} else {
			sex = 9
		}

		patinfo := global.PatientInfo{
			Pat_id:         "",
			Pat_idno:       key.sfzh.String,
			Pat_si_no:      "",
			Pat_name:       key.xingming.String,
			Pat_spell_name: "",
			Pat_sex_code:   "",
			Pat_sex:        sex,
			Pat_brsdate:    key.chushengrq.Time.Format("2006-10-21"),
			Pat_age:        -1,
			Pat_age_unit:   1,
			Pat_weight:     0,
			Pat_mari_code:  0,
			Pat_tel:        key.lianxidh.String,
			Pat_addr:       "",
		}

		var studytype int
		if key.jianchalx.String == "CT" {
			studytype = 3
		} else if key.jianchalx.String == "DR" {
			studytype = 2
		} else if key.jianchalx.String == "MR" {
			studytype = 4
		} else if key.jianchalx.String == "DSA" {
			studytype = 5
		} else if key.jianchalx.String == "US" {
			studytype = 6
		} else if key.jianchalx.String == "ES" {
			studytype = 7
		} else if key.jianchalx.String == "PA" {
			studytype = 8
		} else if key.jianchalx.String == "NM" {
			studytype = 9
		} else if key.jianchalx.String == "PET" {
			studytype = 10
		} else {
			studytype = 99
		}

		applyinfo := global.ApplyInfo{
			Apply_hospital_id:            "",
			Apply_id:                     key.tijianbm.String,
			Apply_status:                 0,
			Apply_time:                   key.TIJIANRQ.Time.Format("2006-05-02 11:22:33"),
			Apply_department_id:          "",
			Apply_department:             "",
			Apply_doctor_id:              "",
			Apply_doctor:                 "",
			Apply_pat_type_code:          "PE",
			Apply_pat_type:               3,
			Apply_clinic_id:              key.sqjlid.String,
			Apply_visit_card_no:          "",
			Apply_medical_record:         "",
			Apply_pat_body_sign:          "",
			Apply_pat_symptoms:           "",
			Apply_indications:            "",
			Apply_clinical_manifestation: "",
			Apply_clinical_diagnosis:     "",
			Apply_chief_complaint:        "",
			Apply_illness_history:        "",
			Apply_report_flag:            0,
			Apply_film_count:             0,
			Apply_film_flag:              0,
			Apply_fee:                    key.danjia.String,
			Apply_vip_flag:               0,
			Apply_isolation_flag:         0,
			Apply_mergency_status:        0,
			Apply_study_type:             studytype,
			Apply_bodyparts_id:           "",
			Apply_bodyparts_name:         "",
			Apply_checkitems_id:          key.ylxh.String,
			Apply_checkitems_name:        key.zuhexmmc.String,
			Apply_details_id:             "",
			Apply_check_note:             "",
			Apply_bodys:                  nil,
			Apply_check_room:             "",
			Apply_section_id:             "",
			Apply_section:                "",
			Apply_sicked_index:           "",
		}
		obj := global.ApplyFormResultData{
			Apply_Info: applyinfo,
			Pat_Info:   patinfo,
		}
		data = append(data, obj)
	}
	return
}
