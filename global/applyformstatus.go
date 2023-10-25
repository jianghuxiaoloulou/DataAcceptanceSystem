package global

// 申请单状态变量

// 申请单状态
const (
	Apply_Status_Canceled   int = 101 // 已取消
	Apply_Status_Registered int = 102 // 已报到
	Apply_Status_Checked    int = 103 // 已检查
	Apply_Status_Drafted    int = 104 // 已起草
	Apply_Status_WaitAudit  int = 105 // 待审核
	Apply_Status_Audited    int = 106 // 已审核
	Apply_Status_Diagnose   int = 107 // 远程诊断
	Apply_Status_ViewRemote int = 108 // 远程查看
	Apply_Status_Charging   int = 199 // 主动计费
)

// 申请单功能模块
const (
	Apply_Func_A string = "A" // 中联HIS申请单取消报到回写
	Apply_Func_B string = "B" // 中联HIS申请单取消报到回写
	Apply_Func_C string = "C" // 中联HIS申请单取消报到回写
	Apply_Func_D string = "D" // 任城区妇保院报告回写
	Apply_Func_E string = "E" // 济宁附属医院远程诊断
	Apply_Func_F string = "F" // 济宁附属医院远程查看
)

// 申请单状态请求
type ApplyFormStatusData struct {
	Bizno      string `json:"bizno" binding:"required"`
	Time       string `json:"time" binding:"required"`
	HospitalID string `json:"hospital_id" binding:"required"`
	PARAM      Param  `json:"req_info" binding:"required"`
}

type Param struct {
	ParamType  int    `json:"param_type"`
	ParamValue string `json:"param_value"`
}

// 第三方PACS申请单和影像数据上传
type ApplyDicomData struct {
	Bizno      string `json:"bizno" binding:"required"`
	Time       string `json:"time" binding:"required"`
	HospitalID string `json:"hospital_id" binding:"required"`
	ApplyID    string `json:"apply_id"`
	PARAM      Param2 `json:"req_info"`
}

// 申请单状态返回
type ApplyFormStatusResult struct {
	Bizno      string  `json:"bizno"`
	Time       string  `json:"time"`
	HospitalID string  `json:"hospital_id"`
	Info       AckInfo `json:"ack_info"`
}
