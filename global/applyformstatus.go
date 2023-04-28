package global

// 申请单状态
const (
	Apply_Status_Canceled   int = 1 // 已取消
	Apply_Status_Registered int = 2 // 已报到
	Apply_Status_Checked    int = 3 // 已检查
	Apply_Status_Drafted    int = 4 // 已起草
	Apply_Status_WaitAudit  int = 5 // 待审核
	Apply_Status_Audited    int = 6 // 已审核
	Apply_Status_Other      int = 9 // 其它
)

// 申请单状态请求
type ApplyFormStatusData struct {
	Bizno string `json:"bizno" binding:"required"`
	Time  string `json:"time" binding:"required"`
	PARAM Param  `json:"req_info" binding:"required"`
}

type Param struct {
	ParamType  int    `json:"param_type"`
	ParamValue string `json:"param_value"`
}

// 申请单状态返回
type ApplyFormStatusResult struct {
	Bizno string  `json:"bizno"`
	Time  string  `json:"time"`
	Info  AckInfo `json:"ack_info"`
}

var (
	ApplyFormStatusDataChan chan ApplyFormStatusData
)
