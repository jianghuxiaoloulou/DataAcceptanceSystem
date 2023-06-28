package global

// 服务编码
const (
	Server_TestConn     string = "0000" // 测试服务通信
	Server_UpdateDict   string = "0001" // 更新字段
	Server_ApplyStatus  string = "1000" // 申请单状态编码
	Server_UploadReport string = "1001" // 第三方PACS报告数据上传
	Server_GetApplyInfo string = "2001" // PACS获取申请单信息
)

// 测试服务通信
type DefaultParam struct {
	Bizno string `json:"bizno" binding:"required"` // 服务编码值
	Time  string `json:"time" binding:"required"`  // 请求时间
}

type DefaultResult struct {
	Bizno string  `json:"bizno"` // 服务编码值
	Time  string  `json:"time"`  // 服务响应时间
	Info  AckInfo `json:"ack_info"`
}
type AckInfo struct {
	Code int    `json:"code"` // 接口响应CODE
	Msg  string `json:"msg"`  // 接口响应消息
}
