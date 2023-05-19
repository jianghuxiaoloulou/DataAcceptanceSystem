package v1

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/pkg/object"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 服务测试通讯
func GetServerTime(c *gin.Context) {
	reqIP := c.ClientIP()
	global.Logger.Debug("请求的主机IP: ", reqIP)
	var testserver global.TestServer
	var result global.TestServerResult
	err := c.ShouldBind(&testserver)
	if err != nil {
		global.Logger.Error(reqIP, " bind error", testserver)
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  err.Error(),
		}
		result.Bizno = testserver.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	global.Logger.Debug(reqIP, " Server Test Data: ", testserver)
	// 返回结果
	ack_info := global.AckInfo{
		Code: 0,
		Msg:  "successful",
	}
	result.Bizno = testserver.Bizno
	result.Time = time.Now().Format("20060102150405")
	result.Info = ack_info
	c.JSON(http.StatusOK, result)
}

// 申请单状态
func ApplyFormStatus(c *gin.Context) {
	reqIP := c.ClientIP()
	global.Logger.Debug("请求的主机IP: ", reqIP)
	var applyformstatus global.ApplyFormStatusData
	var result global.ApplyFormStatusResult
	err := c.ShouldBind(&applyformstatus)
	if err != nil {
		global.Logger.Error(reqIP, " bind error", applyformstatus)
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  err.Error(),
		}
		result.Bizno = applyformstatus.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	global.Logger.Debug(reqIP, " Server Data: ", applyformstatus)
	// 处理数据
	global.ApplyFormStatusDataChan <- applyformstatus
	// 返回结果
	ack_info := global.AckInfo{
		Code: 0,
		Msg:  "successful",
	}
	result.Bizno = applyformstatus.Bizno
	result.Time = time.Now().Format("20060102150405")
	result.Info = ack_info
	c.JSON(http.StatusOK, result)
}

// 获取申请单信息
func ApplyFormInfo(c *gin.Context) {
	reqIP := c.ClientIP()
	global.Logger.Debug("请求的主机IP: ", reqIP)
	var applyforminfo global.ApplyFormInfoData
	var result global.ApplyFormInfoResult
	err := c.ShouldBind(&applyforminfo)
	if err != nil {
		global.Logger.Error(reqIP, " bind error", applyforminfo)
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  err.Error(),
		}
		result.Bizno = applyforminfo.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.PARAM = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	// 获取患者申请单数据信息
	var data []global.ApplyFormResultData
	data = object.GetApplyFormData(applyforminfo)

	// 返回结果
	ack_info := global.AckInfo{
		Code: 0,
		Msg:  "successful",
	}
	result.Bizno = applyforminfo.Bizno
	result.Time = time.Now().Format("20060102150405")
	result.PARAM = ack_info
	result.DATA = data
	c.JSON(http.StatusOK, result)
}
