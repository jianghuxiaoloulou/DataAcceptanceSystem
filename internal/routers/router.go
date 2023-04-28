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
		// 申请单状态接口
		apiv1.POST("/PACS/ApplyFormStatus", v1.ApplyFormStatus)
		// 获取患者申请单信息
		apiv1.POST("/PACS/ApplyFormInfo", v1.ApplyFormInfo)
	}
	return r
}
