package hzhis

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
)

// 华卓HIS包

type HZApplyRequest struct {
	ParamType string      `json:"paramsType"`
	ParamData HZApplyData `json:"paramsData"`
}

type HZApplyData struct {
	DetailId        string `json:"detailId"`        // 申请明细序号
	RequestId       string `json:"requestId"`       // 申请记录序号
	ReqStatus       int    `json:"reqStatus"`       // 申请单状态 状态 3：取消 5：已登记 6：已计费 9：已检查
	UpdateTime      string `json:"updateTime"`      // 操作时间
	StaffNum        string `json:"staffNum"`        // 操作医生工号
	PatiType        string `json:"patiType"`        // 病人就诊类别: 病人类型 OP:门诊 IH:住院等
	ReqType         string `json:"reqType"`         // 病人申请类型：1-检验2-检查3-病理4-手术5- 医技确认（诊疗项目）8- 通费单 9-治疗单 定值：2
	SampleId        string `json:"sampleId"`        // 注意：住院更显状态为9的时候和门诊更新状态为1的时候，sampleid为必传
	AccessionNumber string `json:"accessionNumber"` // pacs检查号
}

// 获取申请单数据
func ByHZHisViewGetApply(object global.ApplyFormInfoData) (data []global.ApplyFormResultData) {
	global.Logger.Debug("开始通过华卓HIS视图获取数据：")
	// 查询视图数据
	param1len := len(object.PARAM)

	// 参数1
	for i := 0; i < param1len; i++ {
		if i > 0 {
			data = append(data, data...)
		}
		switch object.PARAM[i].ParamType {
		case global.Apply_Param_JZKH:
			// 门诊
			result := model.GetMZViewApply(object.PARAM[i], object.PARAM2)
			data = append(data, result...)
			// 住院
			result = model.GetZYViewApply(object.PARAM[i], object.PARAM2)
			data = append(data, result...)
		case global.Apply_Param_MZH:
			// 门诊
			data = model.GetMZViewApply(object.PARAM[i], object.PARAM2)
		case global.Apply_Param_ZYH:
			// 住院
			data = model.GetZYViewApply(object.PARAM[i], object.PARAM2)
		case global.Apply_Param_KSID:
		case global.Apply_Param_TJH:
			// 体检
			data = model.GetTJViewApply(object.PARAM[i], object.PARAM2)
		case global.Apply_Param_MZSQDH:
			// 门诊
			data = model.GetMZViewApply(object.PARAM[i], object.PARAM2)
		case global.Apply_Param_ZYSQDH:
			// 住院
			data = model.GetZYViewApply(object.PARAM[i], object.PARAM2)
		case global.Apply_Param_TJSQDH:
			// 体检
			data = model.GetTJViewApply(object.PARAM[i], object.PARAM2)
		case global.Apply_Param_SFZH:
			// 门诊
			result := model.GetMZViewApply(object.PARAM[i], object.PARAM2)
			data = append(data, result...)
			// 住院
			result = model.GetZYViewApply(object.PARAM[i], object.PARAM2)
			data = append(data, result...)
			// 体检
			result = model.GetTJViewApply(object.PARAM[i], object.PARAM2)
			data = append(data, result...)
		default:
			// 门诊
			result := model.GetMZViewApply(object.PARAM[i], object.PARAM2)
			data = append(data, result...)
			// 住院
			result = model.GetZYViewApply(object.PARAM[i], object.PARAM2)
			data = append(data, result...)
			// 体检
			result = model.GetTJViewApply(object.PARAM[i], object.PARAM2)
			data = append(data, result...)
		}
	}
	return
}

// 更新申请单信息
func UpdateApplyInfo() {

}
