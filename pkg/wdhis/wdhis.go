package wdhis

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// 万达信息HIS对接（宁波金唐软件）
type PostRequest struct {
	Body map[string]interface{} `json:"body"`
}

// 获取申请单数据返回的报文
type ApplyResponse struct {
	Body BodyS `json:"BODY"`
}

type BodyS struct {
	Head     ResHead   `json:"HEAD"`
	Response ResponseS `json:"RESPONSE"`
}

type ResponseS struct {
	ExamRequest []WDHISApplyInfo `json:"exam_request"`
}

// EX1002
type ExamInfoS struct {
	ExamInfo EX1002Request `json:"exam_info"`
}

// EX1003
type ReportInfoS struct {
	ReportInfo EX1003Request `json:"report_info"`
}

// EX1004
type AuditInfoS struct {
	AuditInfo EX1004Request `json:"audit_info"`
}

// 请求Head
type ReqHead struct {
	UserId   string `json:"userid"`
	PassWord string `json:"password"`
	TransNo  string `json:"trans_no"`
}

// EX1101
type RevokeInfoS struct {
	RevokeInfo EX1101Request `json:"revoke_info"`
}

// 返回Head
type ResHead struct {
	RetCode string `json:"RET_CODE"` // 交易返回值 为0时表示成功，其它表示交易失败
	RetInfo string `json:"RET_INFO"` // 交易返回内容
}

// 检查申请单信息
type WDHISApplyInfo struct {
	RequestNo     json.Number `json:"REQUEST_NO"`      // 申请单号 检查申请单的唯一内码
	BranchCode    string      `json:"BRANCH_CODE"`     // 申请医疗机构代码 医疗机构代码，HIS内部分配机构代码
	PatId         json.Number `json:"PAT_ID"`          // 病人号 病人号，病人在医院的唯一编号
	SourceType    json.Number `json:"SOURCE_TYPE"`     // 申请来源 申请来源(1=门诊/2=住院/3=体检/7=急诊/8=门诊留观/99=其它)
	RegId         json.Number `json:"REG_ID"`          // 就诊号 与申请来源对应的具体就诊号(住院=住院ID/门诊=门诊ID/体检=体检ID)
	Name          string      `json:"NAME"`            // 病人姓名
	Sex           json.Number `json:"SEX"`             // 性别 (1=男/2=女/3=不详/9=其它)
	DateOfBirth   string      `json:"DATE_OF_BIRTH"`   // 出生日期 格式：yyyy-MM-dd
	Addr          string      `json:"ADDR"`            // 病人地址
	IdCard        string      `json:"ID_CARD"`         // 身份证号
	Profession    string      `json:"PROFESSION"`      // 职业
	Married       string      `json:"MARRIED"`         // 婚姻状况
	BloodAbo      string      `json:"BLOOD_ABO"`       // ABO血型
	BloodRh       string      `json:"BLOOD_RH"`        // RH血型 1=Rh阴性/2=Rh阳性/3=不详
	Nationality   string      `json:"NATIONALITY"`     // 国籍
	Ethnic        string      `json:"ETHNIC"`          // 名族
	Email         string      `json:"EMAIL"`           // email
	MobilePhone   string      `json:"MOBILE_PHONE"`    // 移动电话
	BedNo         string      `json:"BED_NO"`          // 当前床位号
	InDiagCode    string      `json:"IN_DIAG_CODE"`    // 入院诊断代码
	InDiagName    string      `json:"IN_DIAG_NAME"`    // 入院诊断名称
	RegDeptCode   string      `json:"REG_DEPT_CODE"`   // 门诊挂号科室代码
	RegDeptName   string      `json:"REG_DEPT_NAME"`   // 门诊挂号科室名称
	State         json.Number `json:"STATE"`           // 检查申请状态 1-预开/2-审核/3-提交/5-修改/7-登记/11-报告/13-报告审核完成/90=作废/98-撤销
	IsEme         json.Number `json:"IS_EME"`          // 是否加急 1=加急/0=非加急
	Efid          string      `json:"EFID"`            // 电子定位码 条码/RFID等定位位
	ExamType      string      `json:"EXAM_TYPE"`       // 检查分类
	ExamTypeName  string      `json:"EXAM_TYPE_NAME"`  // 检查分类名称
	ExecDeptCode  string      `json:"EXEC_DEPT_CODE"`  // 执行科室代码，表明指定执行此申请的检查科室
	ExecDeptName  string      `json:"EXEC_DEPT_NAME"`  // 执行科室名称
	ReqWard       string      `json:"REQ_WARD"`        // 申请病区
	ReqDeptCode   string      `json:"REQ_DEPT_CODE"`   // 申请科室代码
	ReqDept       string      `json:"REQ_DEPT"`        // 申请科室
	ReqDoctor     string      `json:"REQ_DOCTOR"`      // 申请医生
	ReqDoctorId   json.Number `json:"REQ_DOCTOR_ID"`   // 申请医生工号
	ReqDoctorCode string      `json:"REQ_DOCTOR_CODE"` // 申请医生工号
	Diag          string      `json:"DIAG"`            // 临床诊断
	BriefMedH     string      `json:"BRIEF_MED_H"`     // 简要病史
	AllePastH     string      `json:"ALLE_PAST_H"`     // 过敏既往史
	RegEmpid      json.Number `json:"REG_EMPID"`       // 登记医生工号
	RegTime       string      `json:"REG_TIME"`        // 登记时间
	RegEmpCode    string      `json:"REG_EMP_CODE"`    // 登记医生工号
	// IsSkinTest       string      `json:"IS_SKINTEST"`        // 是否皮试 1=皮试/0=非皮试
	StResult         string      `json:"ST_RESULT"`          // 皮试结果 P=阳性/N=阴性/0=不需要做(即非皮试)/1=未做(即皮试未做)
	IsAnes           int         `json:"IS_ANES"`            // 是否无痛 1=无痛/0=非无痛
	IsEnhance        int         `json:"IS_ENHANCE"`         // 是否增强 1=增强/0=非增强
	NeedFilm         int         `json:"NEED_FILM"`          // 是否需要片子 0=不要,1=需要
	PrivacyLevel     json.Number `json:"PRIVACY_LEVEL"`      // 稳私级别 (0=普通/1=秘密/2=机密/3=绝密)
	PrnCardTime      string      `json:"PRN_CARD_TIME"`      // 条码打印时间 格式：yyyy-MM-dd hh24:mi:ss
	InvalidEmpid     json.Number `json:"INVALID_EMPID"`      // 作废人员
	InvalidTime      string      `json:"INVALID_TIME"`       // 作废时间 格式：yyyy-MM-dd hh24:mi:ss
	InvalidEmpCode   string      `json:"INVALID_EMP_CODE"`   // 作废人员
	InvalidReason    string      `json:"INVALID_REASON"`     // 作废原因
	ReqRemark        string      `json:"REQ_REMARK"`         // 申请备注
	Note             string      `json:"NOTE"`               // 注意事项
	Remark           string      `json:"REMARK"`             // 其它备注
	SubmitEmpid      json.Number `json:"SUBMIT_EMPID"`       // 提交申请人员
	SubmitTime       string      `json:"SUBMIT_TIME"`        // 提交申请时间
	SubmitEmpCode    string      `json:"SUBMIT_EMP_CODE"`    // 提交申请人员
	RevokeEmpid      json.Number `json:"REVOKE_EMPID"`       // 撤销人员
	RevokeTime       string      `json:"REVOKE_TIME"`        // 撤销时间
	RevokeEmpCode    string      `json:"REVOKE_EMP_CODE"`    // 撤销人员
	ChargeNo         json.Number `json:"CHARGE_NO"`          // 收费编号 如果为空或为0表示未记费
	IsResv           json.Number `json:"IS_RESV"`            // 是否预约 1=预约/0=非预约
	RequestFee       json.Number `json:"REQUEST_FEE"`        // 申请单费用
	ResvId           json.Number `json:"RESV_ID"`            // 预约ID 预约的唯一内码
	ResvDate         string      `json:"RESV_DATE"`          // 预约日期 格式：yyyy-MM-dd
	ResvTimePeriod   json.Number `json:"RESV_TIME_PERIOD"`   // 预约时段 (1=全天/2=上午/3=下午)
	ResvNum          json.Number `json:"RESV_NUM"`           // 预约号码
	ResvTimeLimit    string      `json:"RESV_TIME_LIMIT"`    // 预约参考时间段 预约参考时间段，固定格式(08:10-08:30)
	ResvtimeDescribe string      `json:"RESV_TIME_DESCRIBE"` // 预约参考时间描述
	IsArrange        int         `json:"IS_ARRANGE"`         // 是否安排预约 1=安排/0=非安排
	ArrangeDate      string      `json:"ARRANGE_DATE"`       // 预约安排日期 格式：yyyy-MM-dd
	ArrangeNo        json.Number `json:"ARRANGE_NO"`         // 预约安排序号 PACS预约安排回写
	ArrangeDescribe  string      `json:"ARRANGE_DESCRIBE"`   // 预约安排时间描述
	ArrangeTime      string      `json:"ARRANGE_TIME"`       // 预约安排时间时间段
	ArrangeRoom      string      `json:"ARRANGE_ROOM"`       // 预约安排房间
	CultureLevel     string      `json:"CULTURE_LEVEL"`      // 文化程度
	Barcode          string      `json:"BARCODE"`            // 条码号
	FeeNature        string      `json:"FEE_NATURE"`         // 费别类型
	DoctorPhone      string      `json:"DOCTOR_PHONE"`       // 移动电话
	PatientCardNo    string      `json:"PATIENT_CARDNO"`     // 卡号
	PatientCardType  int         `json:"PATIENT_CARDTYPE"`   // 卡类型
	Physical         string      `json:"PHYSICAL"`           // 体格
	IsMdro           json.Number `json:"IS_MDRO"`            // 多重耐药菌标志 多重耐药菌标志(0=否,1=是)
	DiscountTypeId   string      `json:"DISCOUNT_TYPE_ID"`   // 优惠类型代码
	DiscountTypeName string      `json:"DISCOUNT_TYPE_NAME"` // 优惠类型名称
	FilmFees         []FilmFeeS  `json:"film_fee"`           // 检查胶片费用明细
	ExamItems        []ExamItemS `json:"exam_item"`          // 多个检查项目循环开始
}

// 检查胶片费用明细
type FilmFeeS struct {
	Seq     json.Number `json:"SEQ"`      // 费用序号
	Name    string      `json:"NAME"`     // 费用名称
	Num     json.Number `json:"NUM"`      // 数量
	Price   json.Number `json:"PRICE"`    // 单价
	FeeCode string      `json:"FEE_CODE"` // 费用代码
}

// 多个检查项目循环开始
type ExamItemS struct {
	RequestNo json.Number `json:"REQUEST_NO"` // 申请单号 检查申请单的唯一内码
	ItemSeq   json.Number `json:"ITEM_SEQ"`   // 项目序号 根据申请单项目数累增
	ItemCode  string      `json:"ITEM_CODE"`  // 项目代码
	ItemName  string      `json:"ITEM_NAME"`  // 项目名称
	ItemFee   float64     `json:"ITEM_FEE"`   // 项目费用
	State     json.Number `json:"STATE"`      // 项目状态 7=未检查/9-已检查/99-取消
	ExamParts []ExamPartS `json:"exam_part"`  // 单个检查项目下的多个检查部位循环开始
	// ExamTime     string      `json:"EXAM_TIME"`    // 检查时间，预约检查的时间(PACS) 格式：yyyy-MM-ddThh24:mi:ss
	// ExamNo       string      `json:"EXAM_NO"`      // 检查号
	// ExamEmpid    string      `json:"EXAM_EMPID"`   // 检查医生工号(PACS)
	// InvalidEmpid string      `json:"NVALID_EMPID"` // 项目取消人员
	// InvalideTime string      `json:"INVALID_TIME"` // 项目取消时间
}

type ExamPartS struct {
	PartSeq      json.Number `json:"ORDER_SEQ"`      // 部位序号
	PartCode     string      `json:"PART_CODE"`      // 部位代码
	PartName     string      `json:"PART_NAME"`      // 部位名称
	StdOrderCode string      `json:"STD_ORDER_CODE"` // 中心码
	StdOrderName string      `json:"STD_ORDER_NAME"` // 中心名称
}

// 通过申请单号或条码获取检查申请单信息
type EX0001Request struct {
	QueryWay    string `json:"query_way"`     // 查询方式 （1=申请单号2=条码号3=姓名4=门诊号,5=住院号,6=体检号7=发票号8=病人编号9=身份证号10=就诊卡号）
	QueryValue  string `json:"query_value"`   // 查询 查询值(申请单号/条码号)
	BranchCode  string `json:"branch_code"`   // 机构代码 机构代码（HIS定义每家机构有独立的代码)
	ExamClass   string `json:"exam_class"`    // 检查分类 检查的分类代码
	GetAllState string `json:"get_all_state"` // 是否获取未提交的申请单(false)
}

// 根据申请时间查询的检查列表
type EX0002Request struct {
	ReqTimeBegin string `json:"req_time_begin"` // 申请单开始日期
	ReqTimeEnd   string `json:"req_time_end"`   // 申请单结束日期
	ExamClase    string `json:"exam_class"`     // 检查分类
	DeptCode     string `json:"dept_code"`      // 执行科室
	BranchCode   string `json:"branch_code"`    // 医疗机构代码
	SourceType   string `json:"source_type"`    // 来源 1=门诊/2=住院/3=体检/7=急诊/8=门诊留观/99=其它，为空时默认不根据来源过滤
}

// 通过申请单号或条码获取检查申请单信息（多部位）
type EX0013Request struct {
	QueryWay    string `json:"query_way"`     // 查询方式
	QueryValue  string `json:"query_value"`   // 查询
	BranchCode  string `json:"branch_code"`   // 机构代码
	ExamClass   string `json:"exam_class"`    // 检查分类 多个用逗号分割
	GetAllState string `json:"get_all_state"` // 是否获取未提交的申请单(false)
}

// 根据申请时间查询的检查列表（多部位）
type EX0014Request struct {
	ReqTimeBegin string `json:"req_time_begin"` // 申请单开始日期
	ReqTimeEnd   string `json:"req_time_end"`   // 申请单结束日期
	ExamClase    string `json:"exam_class"`     // 检查分类
	DeptCode     string `json:"dept_code"`      // 执行科室
	BranchCode   string `json:"branch_code"`    // 医疗机构代码
}

// 检查登记
type EX1001Request struct {
	RequestNo string `json:"request_no"` // 申请单编号
	RegEmpid  string `json:"reg_empid"`  // 登记人员
	RegTime   string `json:"reg_time"`   // 登记时间 yyyy-MM-dd hh24:mi:ss
}

// 操作回退
type EX1101Request struct {
	RequestNo     string `json:"request_no"`      // 申请单编号
	ItemSeq       string `json:"item_seq"`        // 项目代码
	RevokeTransNo string `json:"revoke_trans_no"` // 撤销交易码 支持EX1001,EX1002,EX1003,EX1004等交易的回退
	RevokeEmpid   string `json:"revoke_empid"`    // 撤销人员工号
	RevokeTime    string `json:"revoke_time"`     // 撤销时间 格式：yyyy-MM-dd hh24:mi:ss
}

// 检查
type EX1002Request struct {
	RequestNo   string `json:"request_no"`   // 申请单编号
	ItemCode    string `json:"item_code"`    // 项目代码
	ItemSeq     string `json:"item_seq"`     // 项目序号
	ExamNo      string `json:"exam_no"`      // 检查号
	ExamEmpid   string `json:"exam_empid"`   // 检查人员
	ExamEmpname string `json:"exam_empname"` // 检查医生姓名
	ExamTime    string `json:"exam_time"`    // 检查时间 格式：yyyy-MM-dd hh24:mi:ss
	Mani        string `json:"mani"`         // 诊断表现
	Conclusion  string `json:"conclusion"`   // 诊断结论
	WebUrl      string `json:"web_url"`      // WEB调用地址
}

// 检查报告
type EX1003Request struct {
	RequestNo     string `json:"request_no"`     // 申请单编号
	ExamNo        string `json:"exam_no"`        // 项目序号
	ItemSeq       string `json:"item_seq"`       // 检查号
	ExamEmpid     string `json:"exam_empid"`     // 检查人员
	ExamEmpanme   string `json:"exam_empname"`   // 检查人员姓名
	ExamTime      string `json:"exam_time"`      // 检查时间
	ReportEmpid   string `json:"report_empid"`   // 报告医生工号
	ReportEmpname string `json:"report_empname"` // 报告人员姓名
	ReportTime    string `json:"report_time"`    // 报告时间
	Mani          string `json:"mani"`           // 诊断表现
	Conclusion    string `json:"conclusion"`     // 诊断结论
	WebUrl        string `json:"web_url"`        // WEB调用地址
	AbnormalFlag  string `json:"abnormal_flag"`  // 阴阳性结果
	AbnormalSign  string `json:"abnormal_sign"`  // 异常标志
	RefRange      string `json:"ref_range"`      // 正常参考值范围
}

// 报告审核
type EX1004Request struct {
	RequestNo     string `json:"request_no"`     // 申请单编号
	ItemSeq       string `json:"item_seq"`       // 项目代码
	AuditEmpid    string `json:"audit_empid"`    // 审核医生工号
	AuditEmpname  string `json:"audit_empname"`  // 审核人员姓名
	AuditTime     string `json:"audit_time"`     // 审核时间
	ReportTime    string `json:"report_time"`    // 报告时间
	ReportEmpid   string `json:"report_empid"`   // 报告医生工号
	ReportEmpname string `json:"report_empname"` // 报告医生工号
	Mani          string `json:"mani"`           // 诊断表现
	Conclusion    string `json:"conclusion"`     // 诊断结论
	WebUrl        string `json:"web_url"`        // WEB调用地址
	AbnormalFlag  string `json:"abnormal_flag"`  // 阴阳性结果
	AbnormalSign  string `json:"abnormal_sign"`  // 异常标志
	RefRange      string `json:"ref_range"`      // 正常参考值范围
}

// 申请单撤销
type EX1006Request struct {
	RequestNo    string `json:"request_no"`    // 申请单编号
	RevokeEmpid  string `json:"revoke_empid"`  // 撤销人员
	RevokeTime   string `json:"revoke_time"`   // 撤销时间
	RevokeReason string `json:"revoke_reason"` // 撤销原因
	ItemSeq      string `json:"item_seq"`      // 项目代码
}

// 危急情况通知
type EX2000Request struct {
	MessageType    string `json:"message_type"`    // 消息类型 1(危急)/2(普通)
	RequestNo      string `json:"request_no"`      // 申请单号
	MessageContent string `json:"message_content"` // 消息内容
	OperTime       string `json:"oper_time"`       // 操作时间 格式：yyyy-MM-dd hh24:mi:ss
	OperId         string `json:"oper_id"`         // 操作人员工号
	OperName       string `json:"oper_name"`       // 操作人员
	SourceId       string `json:"source_id"`       // 来源ID
	BranchCode     string `json:"branch_code"`     // 机构号 机构代码（HIS分配）
}

// 获取申请单信息
func GetApplyData(hospital global.HospitalConfig, object global.ApplyFormInfoData) (count int, data []global.ApplyFormResultData) {
	// 1.获取HIS配置信息
	hisconfig, err := model.GetHisConfig(int(hospital.HISType.Int16))
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Debug("获取到的HIS厂商信息：", hisconfig)
	// 就诊类型
	var SourceType []int
	for _, value := range object.PatientType {
		var his int
		switch value {
		case 1001: // 门诊
			his = 1
		case 1002: // 住院
			his = 2
		case 1003: // 体检
			his = 3
		}
		SourceType = append(SourceType, his)
	}

	// 急诊状态
	if object.MergencySta == global.Mergency_Type_True {
		SourceType = append(SourceType, 7)
	}

	// 检查类型
	var HisTypeCode []string
	for _, value := range object.StudyType {
		// 通过字典获取数据(通过检查类型去关联的HIS对照表获取HIS的检查分类)
		hiscode := model.ByCodeGetHisCode(value)
		if hiscode != "" {
			HisTypeCode = append(HisTypeCode, hiscode)
		}
	}

	// 查询参数1
	paramllen := len(object.PARAM)
	for i := 0; i < paramllen; i++ {
		var queryway, queryvalue string
		global.Logger.Debug("开始通过EX0001 开始获取申请单数据")
		switch object.PARAM[i].ParamType {
		case global.Apply_Param_JZKH:
			queryway = "10"
		case global.Apply_Param_MZH:
			queryway = "4"
		case global.Apply_Param_ZYH:
			queryway = "5"
		case global.Apply_Param_BLH:
			queryway = "8"
		case global.Apply_Param_TJH:
			queryway = "6"
		case global.Apply_Param_MZSQDH, global.Apply_Param_ZYSQDH, global.Apply_Param_TJSQDH:
			queryway = "1"
		case global.Apply_Param_SFZH:
			queryway = "9"
		case global.Apply_Param_XM:
			queryway = "3"
		}
		queryvalue = object.PARAM[i].ParamValue

		if queryvalue == "" {
			global.Logger.Debug("获取单条申请单param_value 不能为空")
			break
		}

		var examclass string
		for j := 0; j < len(HisTypeCode); j++ {
			if j > 0 {
				examclass += ","
			}
			examclass += HisTypeCode[j]
		}

		ex0001 := EX0001Request{
			QueryWay:    queryway,
			QueryValue:  queryvalue,
			BranchCode:  hisconfig.HISTypeName.String,
			ExamClass:   examclass,
			GetAllState: "false",
		}
		wdhisvalue := FuncEX0001(ex0001, hisconfig.HISInterfaceURL.String)
		obj := HisApplyToPacsApply(object.HospitalID, wdhisvalue)
		data = append(data, obj...)
	}

	// 查询参数2 时间
	param2len := len(object.PARAM2)
	for i := 0; i < param2len; i++ {
		start := object.PARAM2[i].StartDate
		if strings.Contains(start, " ") {
			start = start[:strings.Index(start, " ")]
		}
		end := object.PARAM2[i].EndDate
		if strings.Contains(end, " ") {
			end = end[:strings.Index(end, " ")]
		}
		if start == "" || end == "" {
			global.Logger.Debug("获取申请单时间范围内的申请单,开始或结束时间不能为空")
			break
		}
		// 检查分类代码
		if len(HisTypeCode) == 0 {
			for _, dict := range global.DictDatas {
				if 2000 == dict.Type {
					hiscode := dict.HisCode
					if hiscode != "" {
						HisTypeCode = append(HisTypeCode, hiscode)
					}
				}
			}
		}
		var examclass string
		for j := 0; j < len(HisTypeCode); j++ {
			if j > 0 {
				examclass += ","
			}
			examclass += HisTypeCode[j]
		}

		ex0002 := EX0002Request{
			ReqTimeBegin: start,
			ReqTimeEnd:   end,
			ExamClase:    examclass,
			BranchCode:   hisconfig.HISTypeName.String,
		}
		wdhisvalue := FuncEX0002(ex0002, hisconfig.HISInterfaceURL.String)
		obj := HisApplyToPacsApply(object.HospitalID, wdhisvalue)
		data = append(data, obj...)
	}
	return len(data), data
}

// HIS返回数据转换未PACS数据
func HisApplyToPacsApply(hospitalid string, wdhisvalue []WDHISApplyInfo) (data []global.ApplyFormResultData) {
	for _, hisvalue := range wdhisvalue {
		var sexCode int
		var sexName string
		if hisvalue.Sex == "1" {
			sexCode = 3001
			sexName = "男"
		} else if hisvalue.Sex == "2" {
			sexCode = 3002
			sexName = "女"
		} else {
			sexCode = 9999
			sexName = "其他"
		}
		patinfo := global.PatientInfo{
			Pat_id:       hisvalue.PatId.String(),
			Pat_idno:     hisvalue.IdCard,
			Pat_si_no:    hisvalue.PatientCardNo,
			Pat_name:     hisvalue.Name,
			Pat_sex_code: sexName,
			Pat_sex:      sexCode,
			Pat_brsdate:  hisvalue.DateOfBirth[:strings.Index(hisvalue.DateOfBirth, " ")],
			Pat_tel:      hisvalue.MobilePhone,
			Pat_addr:     hisvalue.Addr,
		}

		// 获取就诊类型编码
		sourcetype := hisvalue.SourceType
		var pattype int
		var patcode string
		switch sourcetype {
		case "1":
			pattype = 1001
			patcode = "OP"
		case "2":
			pattype = 1002
			patcode = "IH"
		case "3":
			pattype = 1003
			patcode = "PE"
		default:
			pattype = 9999
		}
		emestatus, _ := hisvalue.IsEme.Int64()

		var bodysarr []global.CheckBody

		for _, examitem := range hisvalue.ExamItems {
			var flag bool
			var itemsarr []global.CheckItem
			item := global.CheckItem{
				Apply_detail_id:       examitem.ItemSeq.String(),
				Apply_check_item_code: examitem.ItemCode,
				Apply_check_item_name: examitem.ItemName,
			}
			itemsarr = append(itemsarr, item)
			for _, exampart := range examitem.ExamParts {
				body := global.CheckBody{
					Apply_bodypart_code: exampart.PartCode,
					Apply_bodypart_name: exampart.PartName,
					Apply_bodypart_id:   exampart.PartSeq.String(),
				}
				for _, v := range bodysarr {
					if v.Apply_bodypart_code == body.Apply_bodypart_code {
						v.Apply_checkItems = append(v.Apply_checkItems, itemsarr...)
						flag = true
					}
				}
				if !flag {
					body.Apply_checkItems = itemsarr
					bodysarr = append(bodysarr, body)
				}
			}
		}
		fee := hisvalue.RequestFee.String()
		if fee == "" {
			fee = "0"
		}
		applyinfo := global.ApplyInfo{
			Apply_hospital_id:        hospitalid,
			Apply_id:                 hisvalue.RequestNo.String(),
			Apply_time:               hisvalue.SubmitTime,
			Apply_fee:                fee,
			Apply_department_id:      hisvalue.ReqDeptCode,
			Apply_department:         hisvalue.ReqDept,
			Apply_doctor_id:          hisvalue.ReqDoctorId.String(),
			Apply_doctor:             hisvalue.ReqDoctor,
			Apply_pat_type_code:      patcode,
			Apply_pat_type:           pattype,
			Apply_clinic_id:          hisvalue.RegId.String(),
			Apply_visit_card_no:      hisvalue.PatientCardNo,
			Apply_medical_record:     hisvalue.Barcode,
			Apply_pat_body_sign:      hisvalue.Physical,
			Apply_clinical_diagnosis: hisvalue.Diag,
			Apply_illness_history:    hisvalue.BriefMedH,
			Apply_mergency_status:    int(emestatus),
			Apply_study_type:         model.ByHisCodeGetCode(hisvalue.ExamType),
			Apply_check_note:         hisvalue.Note,
			Apply_check_room:         hisvalue.ExecDeptName,
			Apply_sicked_index:       hisvalue.BedNo,
			Apply_bodys:              bodysarr,
		}
		obj := global.ApplyFormResultData{
			Apply_Info: applyinfo,
			Pat_Info:   patinfo,
		}
		data = append(data, obj)
	}
	return
}

// 获取申请单信息(EX0001)（具体申请单信息）
func FuncEX0001(ex0001 EX0001Request, url string) (data []WDHISApplyInfo) {
	reqdata := make(map[string]interface{})
	reqdata["head"] = ReqHead{
		UserId:   "LDPACS",
		PassWord: "LD123",
		TransNo:  "EX0001",
	}
	reqdata["request"] = ex0001
	request := PostRequest{
		Body: reqdata,
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		global.Logger.Error("json Marshal err ", err.Error())
		return data
	}
	global.Logger.Debug("执行代码:EX0001,通过申请单号或条码获取检查申请单信息", string(requestData))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if err != nil {
		global.Logger.Error("http NewRequest err ", err.Error())
		return data
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		global.Logger.Error("http Do err ", err.Error())
		return data
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		global.Logger.Error("ioutil.ReadAll err ", err.Error())
		return data
	}
	global.Logger.Debug("EX0001：", string(body))
	resdata := ApplyResponse{}
	err = json.Unmarshal(body, &resdata)
	if err != nil {
		global.Logger.Error("resp.Body err ", err)
		return data
	}

	if resdata.Body.Head.RetCode != "0" {
		global.Logger.Debug("EX0001 获取数据是吧，", resdata.Body.Head.RetInfo)
		return data
	}
	data = append(data, resdata.Body.Response.ExamRequest...)
	return data
}

// 获取申请单列表(EX0002)（时间段中申请单列表）
func FuncEX0002(ex0002 EX0002Request, url string) (data []WDHISApplyInfo) {
	reqdata := make(map[string]interface{})
	reqdata["head"] = ReqHead{
		UserId:   "LDPACS",
		PassWord: "LD123",
		TransNo:  "EX0002",
	}
	reqdata["request"] = ex0002
	request := PostRequest{
		Body: reqdata,
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		global.Logger.Error("json Marshal err ", err.Error())
		return data
	}
	global.Logger.Debug("执行代码:EX0002,根据申请时间查询的检查列表", string(requestData))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if err != nil {
		global.Logger.Error("http NewRequest err ", err.Error())
		return data
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		global.Logger.Error("http Do err ", err.Error())
		return data
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		global.Logger.Error("ioutil.ReadAll err ", err.Error())
		return data
	}
	global.Logger.Debug("EX0002：", string(body))
	resdata := ApplyResponse{}
	err = json.Unmarshal(body, &resdata)
	if err != nil {
		global.Logger.Error("resp.Body err ", err)
		return data
	}

	if resdata.Body.Head.RetCode != "0" {
		global.Logger.Debug("EX0002 获取数据是吧，", resdata.Body.Head.RetInfo)
		return data
	}
	data = append(data, resdata.Body.Response.ExamRequest...)
	return data
}

// 获取申请单信息(EX0013)（具体申请单信息，用于多部位登记）
func FuncEX0013(ex0013 EX0013Request, url string) (data []WDHISApplyInfo) {
	reqdata := make(map[string]interface{})
	reqdata["head"] = ReqHead{
		UserId:   "LDPACS",
		PassWord: "LD123",
		TransNo:  "EX0013",
	}
	reqdata["request"] = ex0013
	request := PostRequest{
		Body: reqdata,
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		global.Logger.Error("json Marshal err ", err.Error())
		return data
	}
	global.Logger.Debug("执行代码:EX0013,用于多部位登记通过申请单号或条码获取检查申请单信息", requestData)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if err != nil {
		global.Logger.Error("http NewRequest err ", err.Error())
		return data
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		global.Logger.Error("http Do err ", err.Error())
		return data
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		global.Logger.Error("ioutil.ReadAll err ", err.Error())
		return data
	}
	global.Logger.Debug("EX0013：", string(body))
	resdata := ApplyResponse{}
	err = json.Unmarshal(body, &resdata)
	if err != nil {
		global.Logger.Error("resp.Body err ", err)
		return data
	}

	if resdata.Body.Head.RetCode != "0" {
		global.Logger.Debug("EX0013 获取数据是吧，", resdata.Body.Head.RetInfo)
		return data
	}
	data = append(data, resdata.Body.Response.ExamRequest...)
	return data
}

// 获取申请单列表(EX0014)（时间段中申请单列表，用于多部位登记）
func FuncEX0014(ex0014 EX0014Request, url string) (data []WDHISApplyInfo) {
	reqdata := make(map[string]interface{})
	reqdata["head"] = ReqHead{
		UserId:   "LDPACS",
		PassWord: "LD123",
		TransNo:  "EX0014",
	}
	reqdata["request"] = ex0014
	request := PostRequest{
		Body: reqdata,
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		global.Logger.Error("json Marshal err ", err.Error())
		return data
	}
	global.Logger.Debug("执行代码:EX0014,时间段中申请单列表，用于多部位登记", string(requestData))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if err != nil {
		global.Logger.Error("http NewRequest err ", err.Error())
		return data
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		global.Logger.Error("http Do err ", err.Error())
		return data
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		global.Logger.Error("ioutil.ReadAll err ", err.Error())
		return data
	}
	global.Logger.Debug("EX0014：", string(body))
	resdata := ApplyResponse{}
	err = json.Unmarshal(body, &resdata)
	if err != nil {
		global.Logger.Error("resp.Body err ", err)
		return data
	}

	if resdata.Body.Head.RetCode != "0" {
		global.Logger.Debug("EX0014 获取数据是吧，", resdata.Body.Head.RetInfo)
		return data
	}
	data = append(data, resdata.Body.Response.ExamRequest...)
	return data
}

// (函数功能A)检查登记（EX1001）value,登记后PACS的产生的申请单唯一ID
func FuncEX1001(value string) {
	qydata := model.GetQYPACSRegisterInfo(value)
	data := global.QYPACSRegisterInfo{}
	if len(qydata) > 0 {
		data = qydata[0]
	}
	if data.ApplydetailId.String == "" {
		global.Logger.Debug("申请单明细为空,不是HIS过来的数据,不需要回写HIS", data)
		return
	}
	// 1.通过HospitalID 获取医院相关数据库连接信息
	hospitalConfig, _ := model.GetHospitalConfig(data.HospitalID.String)
	// 1.获取HIS配置信息
	hisconfig, _ := model.GetHisConfig(int(hospitalConfig.HISType.Int16))

	// 获取HIS登记医生ID
	hisid := model.GetHisPersonID(data.RegisterDoctorId.String, data.HospitalID.String)

	// 获取回写his的申请单编号，登记人员，登记时间
	ex1001 := EX1001Request{
		RequestNo: data.ApplyId.String,
		RegEmpid:  hisid,
		RegTime:   data.RegisterTime.String,
	}

	reqdata := make(map[string]interface{})
	reqdata["head"] = ReqHead{
		UserId:   "LDPACS",
		PassWord: "LD123",
		TransNo:  "EX1001",
	}
	reqdata["request"] = ex1001
	request := PostRequest{
		Body: reqdata,
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		global.Logger.Error("json Marshal err ", err.Error())
		return
	}
	global.Logger.Debug("执行代码:EX1001,PACS登记一条申请单", string(requestData))

	client := &http.Client{}
	req, err := http.NewRequest("POST", hisconfig.HISInterfaceURL.String, bytes.NewBuffer(requestData))
	if err != nil {
		global.Logger.Error("http NewRequest err ", err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		global.Logger.Error("http Do err ", err.Error())
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		global.Logger.Error("ioutil.ReadAll err ", err.Error())
		return
	}
	global.Logger.Debug("EX1001：", string(body))
}

// (函数功能B)检查（EX1002）
func FuncEX1002(value string) {
	qydata := model.GetQYPACSRegisterInfo(value)
	for _, data := range qydata {

		if data.ApplydetailId.String == "" {
			global.Logger.Debug("申请单明细为空,不是HIS过来的数据,不需要回写HIS", data)
			continue
		}
		// 1.通过HospitalID 获取医院相关数据库连接信息
		hospitalConfig, _ := model.GetHospitalConfig(data.HospitalID.String)
		// 1.获取HIS配置信息
		hisconfig, _ := model.GetHisConfig(int(hospitalConfig.HISType.Int16))

		// 获取HIS登记医生ID
		hisid := model.GetHisPersonID(data.StudyDoctorId.String, data.HospitalID.String)

		// 获取回写his的申请单编号，登记人员，登记时间
		EX1002 := EX1002Request{
			RequestNo:   data.ApplyId.String,
			ItemCode:    data.ProjectCode.String,
			ItemSeq:     data.ApplydetailId.String,
			ExamNo:      data.AccessionNumber.String,
			ExamEmpid:   hisid,
			ExamEmpname: data.StudyDoctorName.String,
			ExamTime:    data.StudyTime.String,
		}
		reqdata := make(map[string]interface{})
		reqdata["head"] = ReqHead{
			UserId:   "LDPACS",
			PassWord: "LD123",
			TransNo:  "EX1002",
		}
		reqdata["request"] = ExamInfoS{
			ExamInfo: EX1002,
		}
		request := PostRequest{
			Body: reqdata,
		}
		requestData, err := json.Marshal(request)
		if err != nil {
			global.Logger.Error("json Marshal err ", err.Error())
			return
		}
		global.Logger.Debug("执行代码:EX1002,检查 ", string(requestData))

		client := &http.Client{}
		req, err := http.NewRequest("POST", hisconfig.HISInterfaceURL.String, bytes.NewBuffer(requestData))
		if err != nil {
			global.Logger.Error("http NewRequest err ", err.Error())
			return
		}
		req.Header.Add("Content-Type", "application/json;charset=utf-8")

		res, err := client.Do(req)
		if err != nil {
			global.Logger.Error("http Do err ", err.Error())
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			global.Logger.Error("ioutil.ReadAll err ", err.Error())
			return
		}
		global.Logger.Debug("EX1002：", string(body))

	}
}

// (函数功能C)检查报告（EX1003）
func FuncEX1003(value string) {
	qydata := model.GetQYPACSRegisterInfo(value)
	qyreport := model.GetQYPACSReportInfo(value)
	// 阴阳性
	var abnormalflag string
	if qyreport.PositiveNegativeStatus.Valid && qyreport.PositiveNegativeStatus.Int16 == 1 {
		abnormalflag = "阳性"
	} else {
		abnormalflag = "阴性"
	}
	// 1.通过HospitalID 获取医院相关数据库连接信息
	hospitalConfig, _ := model.GetHospitalConfig(qyreport.HospitalID.String)
	// 1.获取HIS配置信息
	hisconfig, _ := model.GetHisConfig(int(hospitalConfig.HISType.Int16))
	// 获取HIS报告医生ID
	reporthisid := model.GetHisPersonID(qyreport.ReportDoctorId.String, qyreport.HospitalID.String)
	for _, data := range qydata {
		global.Logger.Debug("获取的申请单数据：", data)
		if data.ApplydetailId.String == "" {
			global.Logger.Debug("申请单明细为空,不是HIS过来的数据,不需要回写HIS", data)
			continue
		}
		// 获取HIS检查医生ID
		examhisid := model.GetHisPersonID(data.StudyDoctorId.String, data.HospitalID.String)

		ex1003 := EX1003Request{
			RequestNo:     data.ApplyId.String,
			ItemSeq:       data.ApplydetailId.String,
			ExamNo:        data.AccessionNumber.String,
			ExamEmpid:     examhisid,
			ExamEmpanme:   data.StudyDoctorName.String,
			ExamTime:      data.StudyTime.String,
			ReportEmpid:   reporthisid,
			ReportEmpname: qyreport.ReportDoctorName.String,
			ReportTime:    qyreport.ReportTime.String,
			Mani:          qyreport.Finding.String,
			Conclusion:    qyreport.Conclusion.String,
			AbnormalFlag:  abnormalflag,
		}
		reqdata := make(map[string]interface{})
		reqdata["head"] = ReqHead{
			UserId:   "LDPACS",
			PassWord: "LD123",
			TransNo:  "EX1003",
		}
		reqdata["request"] = ReportInfoS{
			ReportInfo: ex1003,
		}
		request := PostRequest{
			Body: reqdata,
		}
		requestData, err := json.Marshal(request)
		if err != nil {
			global.Logger.Error("json Marshal err ", err.Error())
			return
		}
		global.Logger.Debug("执行代码:EX1003,检查报告 ", string(requestData))

		client := &http.Client{}
		req, err := http.NewRequest("POST", hisconfig.HISInterfaceURL.String, bytes.NewBuffer(requestData))
		if err != nil {
			global.Logger.Error("http NewRequest err ", err.Error())
			return
		}
		req.Header.Add("Content-Type", "application/json;charset=utf-8")

		res, err := client.Do(req)
		if err != nil {
			global.Logger.Error("http Do err ", err.Error())
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			global.Logger.Error("ioutil.ReadAll err ", err.Error())
			return
		}
		global.Logger.Debug("EX1003：", string(body))
	}
}

// (函数功能H)报告审核（EX1004）
func FuncEX1004(value string) {
	qydata := model.GetQYPACSRegisterInfo(value)
	qyreport := model.GetQYPACSReportInfo(value)
	// 阴阳性
	var abnormalflag string
	if qyreport.PositiveNegativeStatus.Valid && qyreport.PositiveNegativeStatus.Int16 == 1 {
		abnormalflag = "阳性"
	} else {
		abnormalflag = "阴性"
	}

	// 1.通过HospitalID 获取医院相关数据库连接信息
	hospitalConfig, _ := model.GetHospitalConfig(qyreport.HospitalID.String)
	// 1.获取HIS配置信息
	hisconfig, _ := model.GetHisConfig(int(hospitalConfig.HISType.Int16))
	// 获取HIS报告医生ID
	reporthisid := model.GetHisPersonID(qyreport.ReportDoctorId.String, qyreport.HospitalID.String)
	audithisid := model.GetHisPersonID(qyreport.AuditDoctorId.String, qyreport.HospitalID.String)
	for _, data := range qydata {
		if data.ApplydetailId.String == "" {
			global.Logger.Debug("申请单明细为空,不是HIS过来的数据,不需要回写HIS", data)
			return
		}
		ex1004 := EX1004Request{
			RequestNo:     data.ApplyId.String,
			ItemSeq:       data.ApplydetailId.String,
			AuditEmpid:    audithisid,
			AuditEmpname:  qyreport.AuditDoctorName.String,
			AuditTime:     qyreport.AuditTime.String,
			ReportEmpid:   reporthisid,
			ReportEmpname: qyreport.ReportDoctorName.String,
			ReportTime:    qyreport.ReportTime.String,
			Mani:          qyreport.Finding.String,
			Conclusion:    qyreport.Conclusion.String,
			AbnormalFlag:  abnormalflag,
		}
		reqdata := make(map[string]interface{})
		reqdata["head"] = ReqHead{
			UserId:   "LDPACS",
			PassWord: "LD123",
			TransNo:  "EX1004",
		}
		reqdata["request"] = AuditInfoS{
			AuditInfo: ex1004,
		}
		request := PostRequest{
			Body: reqdata,
		}
		requestData, err := json.Marshal(request)
		if err != nil {
			global.Logger.Error("json Marshal err ", err.Error())
			return
		}
		global.Logger.Debug("执行代码:ex1004,报告审核 ", string(requestData))

		client := &http.Client{}
		req, err := http.NewRequest("POST", hisconfig.HISInterfaceURL.String, bytes.NewBuffer(requestData))
		if err != nil {
			global.Logger.Error("http NewRequest err ", err.Error())
			return
		}
		req.Header.Add("Content-Type", "application/json;charset=utf-8")

		res, err := client.Do(req)
		if err != nil {
			global.Logger.Error("http Do err ", err.Error())
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			global.Logger.Error("ioutil.ReadAll err ", err.Error())
			return
		}
		global.Logger.Debug("EX1004：", string(body))
	}
}

// (函数功能)操作回退（EX1101）
func FuncEX1101(ex1101 EX1101Request, url string) {
	reqdata := make(map[string]interface{})
	reqdata["head"] = ReqHead{
		UserId:   "LDPACS",
		PassWord: "LD123",
		TransNo:  "EX1101",
	}
	reqdata["request"] = RevokeInfoS{
		RevokeInfo: ex1101,
	}
	request := PostRequest{
		Body: reqdata,
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		global.Logger.Error("json Marshal err ", err.Error())
		return
	}
	global.Logger.Debug("执行代码:EX1101,操作回退", string(requestData))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if err != nil {
		global.Logger.Error("http NewRequest err ", err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		global.Logger.Error("http Do err ", err.Error())
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		global.Logger.Error("ioutil.ReadAll err ", err.Error())
		return
	}
	global.Logger.Debug("EX1101：", string(body))
}

// (函数功能I)申请单撤销（EX1006）
func FuncEX1006(value string) {
	qydata := model.GetQYPACSRegisterInfo(value)
	for _, data := range qydata {
		if data.ApplydetailId.String == "" {
			global.Logger.Debug("申请单明细为空,不是HIS过来的数据,不需要回写HIS", data)
			continue
		}
		// 1.通过HospitalID 获取医院相关数据库连接信息
		hospitalConfig, _ := model.GetHospitalConfig(data.HospitalID.String)
		// 1.获取HIS配置信息
		hisconfig, _ := model.GetHisConfig(int(hospitalConfig.HISType.Int16))

		// 获取HIS登记医生ID
		hisid := model.GetHisPersonID(data.RegisterDoctorId.String, data.HospitalID.String)

		// 获取回写his的申请单编号，登记人员，登记时间
		ex1006 := EX1006Request{
			RequestNo:   data.ApplyId.String,
			RevokeEmpid: hisid,
			RevokeTime:  data.UpdateTime.String,
			ItemSeq:     data.ApplydetailId.String,
		}
		reqdata := make(map[string]interface{})
		reqdata["head"] = ReqHead{
			UserId:   "LDPACS",
			PassWord: "LD123",
			TransNo:  "EX1006",
		}
		reqdata["request"] = ex1006
		request := PostRequest{
			Body: reqdata,
		}
		requestData, err := json.Marshal(request)
		if err != nil {
			global.Logger.Error("json Marshal err ", err.Error())
			return
		}
		global.Logger.Debug("执行代码:EX1006,申请单撤销 ", string(requestData))

		client := &http.Client{}
		req, err := http.NewRequest("POST", hisconfig.HISInterfaceURL.String, bytes.NewBuffer(requestData))
		if err != nil {
			global.Logger.Error("http NewRequest err ", err.Error())
			return
		}
		req.Header.Add("Content-Type", "application/json;charset=utf-8")

		res, err := client.Do(req)
		if err != nil {
			global.Logger.Error("http Do err ", err.Error())
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			global.Logger.Error("ioutil.ReadAll err ", err.Error())
			return
		}
		global.Logger.Debug("EX1006：", string(body))
	}
}

// (函数功能J)危急情况通知（EX2000）
func FuncEX2000(value string) {
	qydata := model.GetQYPACSRegisterInfo(value)
	qyreport := model.GetQYPACSReportInfo(value)
	// 危急值
	if qyreport.CrisisStatus.Valid && qyreport.CrisisStatus.Int16 == 1 {
		global.Logger.Debug("危急值阳性，通知HIS")
		qycrisis := model.GetQYPACSCrisisInfo(value)
		// 1.通过HospitalID 获取医院相关数据库连接信息
		hospitalConfig, _ := model.GetHospitalConfig(qyreport.HospitalID.String)
		// 1.获取HIS配置信息
		hisconfig, _ := model.GetHisConfig(int(hospitalConfig.HISType.Int16))
		// 获取HIS登记医生ID
		hisid := model.GetHisPersonID(qycrisis.RequestDoctorCode.String, qyreport.HospitalID.String)

		for _, data := range qydata {
			global.Logger.Debug("获取的申请单数据：", data)
			ex2000 := EX2000Request{
				MessageType:    "1",
				RequestNo:      data.ApplyId.String,
				MessageContent: qycrisis.CrisisContent.String,
				OperTime:       qycrisis.ProcessTime.String,
				OperId:         hisid,
				OperName:       qycrisis.RequestDoctorName.String,
				BranchCode:     hisconfig.HISTypeName.String,
			}
			reqdata := make(map[string]interface{})
			reqdata["head"] = ReqHead{
				UserId:   "LDPACS",
				PassWord: "LD123",
				TransNo:  "EX2000",
			}
			reqdata["request"] = ex2000
			request := PostRequest{
				Body: reqdata,
			}
			requestData, err := json.Marshal(request)
			if err != nil {
				global.Logger.Error("json Marshal err ", err.Error())
				return
			}
			global.Logger.Debug("执行代码:EX2000,危急情况通知 ", string(requestData))

			client := &http.Client{}
			req, err := http.NewRequest("POST", hisconfig.HISInterfaceURL.String, bytes.NewBuffer(requestData))
			if err != nil {
				global.Logger.Error("http NewRequest err ", err.Error())
				return
			}
			req.Header.Add("Content-Type", "application/json;charset=utf-8")

			res, err := client.Do(req)
			if err != nil {
				global.Logger.Error("http Do err ", err.Error())
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				global.Logger.Error("ioutil.ReadAll err ", err.Error())
				return
			}
			global.Logger.Debug("EX2000：", string(body))
		}
	} else {
		global.Logger.Debug("危急值阴性，终止通知HIS")
		return
	}
}
