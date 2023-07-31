package global

// 服务编码
const (
	Server_TestConn      string = "0000" // 通讯测试
	Server_UpdateDict    string = "0001" // 更新字典服务
	Server_ApplyStatus   string = "1000" // 区域PACS发送申请单状态
	Server_UploadReport  string = "1001" // 第三方PACS报告数据上传
	Server_ApplyAndDicom string = "1002" // 第三方PACS申请单和影像数据上传
	Server_GetApplyInfo  string = "2001" // 区域PACS获取申请单信息
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
