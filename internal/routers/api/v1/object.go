package v1

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	"WowjoyProject/DataAcceptanceSystem/pkg/object"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 服务测试通讯
func GetServerTime(c *gin.Context) {
	reqIP := c.ClientIP()
	global.Logger.Debug("请求的主机IP: ", reqIP)
	var param global.DefaultParam
	var result global.DefaultResult
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Error(reqIP, " bind error", param)
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  err.Error(),
		}
		result.Bizno = param.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	if param.Bizno != global.Server_TestConn {
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  "交易编码错误，正确编码：" + global.Server_TestConn,
		}
		result.Bizno = param.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	global.Logger.Debug(reqIP, " Server Test Data: ", param)
	// 返回结果
	ack_info := global.AckInfo{
		Code: 0,
		Msg:  "successful",
	}
	result.Bizno = param.Bizno
	result.Time = time.Now().Format("20060102150405")
	result.Info = ack_info
	c.JSON(http.StatusOK, result)
}

// 重新获取字典数据
func UpdateDictData(c *gin.Context) {
	reqIP := c.ClientIP()
	global.Logger.Debug("请求的主机IP: ", reqIP)
	var param global.DefaultParam
	var result global.DefaultResult
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.Error(reqIP, " bind error", param)
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  err.Error(),
		}
		result.Bizno = param.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	if param.Bizno != global.Server_UpdateDict {
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  "交易编码错误，正确编码：" + global.Server_UpdateDict,
		}
		result.Bizno = param.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	global.Logger.Debug(reqIP, " get dict data: ", param)
	// 获取字典数据
	model.GetDictData()
	// 返回结果
	ack_info := global.AckInfo{
		Code: 0,
		Msg:  "successful",
	}
	result.Bizno = param.Bizno
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
		result.HospitalID = applyformstatus.HospitalID
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	if applyformstatus.Bizno != global.Server_ApplyStatus {
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  "交易编码错误，正确编码：" + global.Server_ApplyStatus,
		}
		result.Bizno = applyformstatus.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.HospitalID = applyformstatus.HospitalID
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	global.Logger.Debug(reqIP, " Server Data: ", applyformstatus)
	// 单独处理数据
	global.ApplyFormStatusDataChan <- applyformstatus
	// 返回结果
	ack_info := global.AckInfo{
		Code: 0,
		Msg:  "successful",
	}
	result.Bizno = applyformstatus.Bizno
	result.Time = time.Now().Format("20060102150405")
	result.HospitalID = applyformstatus.HospitalID
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
		result.HospitalID = applyforminfo.HospitalID
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	if applyforminfo.Bizno != global.Server_GetApplyInfo {
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  "交易编码错误，正确编码：" + global.Server_GetApplyInfo,
		}
		result.Bizno = applyforminfo.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.HospitalID = applyforminfo.HospitalID
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	// 获取患者申请单数据信息
	var data []global.ApplyFormResultData
	count, data := object.GetApplyFormData(applyforminfo)

	// 返回结果
	ack_info := global.AckInfo{
		Code: 0,
		Msg:  "successful",
	}
	result.Bizno = applyforminfo.Bizno
	result.Time = time.Now().Format("20060102150405")
	result.HospitalID = applyforminfo.HospitalID
	result.Info = ack_info
	result.DataCount = count
	result.DATA = data
	c.JSON(http.StatusOK, result)
}

// 第三方PACS报告上传接口
func UploadReportInfo(c *gin.Context) {
	reqIP := c.ClientIP()
	global.Logger.Debug("请求的主机IP: ", reqIP)
	var reportdata global.ReportData
	var result global.ApplyFormStatusResult
	err := c.ShouldBind(&reportdata)
	if err != nil {
		global.Logger.Error(reqIP, " bind error", reportdata)
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  err.Error(),
		}
		result.Bizno = reportdata.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}
	if reportdata.Bizno != global.Server_UploadReport {
		ack_info := global.AckInfo{
			Code: 1,
			Msg:  "交易编码错误，正确编码：" + global.Server_UploadReport,
		}
		result.Bizno = reportdata.Bizno
		result.Time = time.Now().Format("20060102150405")
		result.Info = ack_info
		c.JSON(http.StatusBadRequest, result)
		return
	}

	global.Logger.Debug(reqIP, " Server Data: ", reportdata)
	// 处理数据 调用后端接口插入报告数据

	// 返回结果
	ack_info := global.AckInfo{
		Code: 0,
		Msg:  "successful",
	}
	result.Bizno = reportdata.Bizno
	result.Time = time.Now().Format("20060102150405")
	result.Info = ack_info
	c.JSON(http.StatusOK, result)
}
