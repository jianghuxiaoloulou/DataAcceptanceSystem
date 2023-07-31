package routers

import (
	v1 "WowjoyProject/DataAcceptanceSystem/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 注册中间件
	apiv1 := r.Group("/api/v1")
	{
		// 通信测试
		apiv1.POST("/", v1.GetServerTime)
		// 更新字典服务
		apiv1.POST("/Update/DictData", v1.UpdateDictData)
		// 区域PACS发送申请单状态
		apiv1.POST("/PACS/ApplyFormStatus", v1.ApplyFormStatus)
		// 区域PACS获取申请单信息
		apiv1.POST("/PACS/ApplyFormInfo", v1.ApplyFormInfo)
		// 第三方PACS报告数据上传
		apiv1.POST("/PACS/Upload/ReportInfo", v1.UploadReportInfo)
		// 第三方PACS申请单和影像数据上传
		apiv1.POST("/PACS/Upload/ApplyAndDicomInfo", v1.UploadApplyAndDicomInfo)
	}
	return r
}
