package global

// 服务编码
const (
	Server_ApplyStatus string = "1000" // 申请单状态编码
	Server_ApplyInfo   string = "2001" // 申请单信息编码
)

// 测试服务通信
type TestServer struct {
	Bizno string `json:"bizno" binding:"required"` // 服务编码值
	Time  string `json:"time" binding:"required"`  // 请求时间
}

type TestServerResult struct {
	Bizno string  `json:"bizno"` // 服务编码值
	Time  string  `json:"time"`  // 服务响应时间
	Info  AckInfo `json:"ack_info"`
}
type AckInfo struct {
	Code int    `json:"code"` // 接口响应CODE
	Msg  string `json:"msg"`  // 接口响应消息
}
