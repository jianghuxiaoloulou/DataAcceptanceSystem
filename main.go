package main

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	initialize "WowjoyProject/DataAcceptanceSystem/internal/init"
	"WowjoyProject/DataAcceptanceSystem/internal/routers"
	"WowjoyProject/DataAcceptanceSystem/pkg/object"
	"WowjoyProject/DataAcceptanceSystem/pkg/workpattern"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title PACS集成平台
// @version 1.0.0.0
// @description PACS集成平台
// @termsOfService https://github.com/jianghuxiaoloulou/DataAcceptanceSystem.git
func main() {
	initialize.ReadSetup()
	global.Logger.Info("***开始运行PACS集成平台服务***")
	global.ApplyFormStatusDataChan = make(chan global.ApplyFormStatusData, global.GeneralSetting.MaxThreads)
	// 注册工作池，传入任务
	// 参数1 初始化worker(工人)设置最大线程数
	wokerPool := workpattern.NewWorkerPool(global.GeneralSetting.MaxThreads)
	// 有任务就去做，没有就阻塞，任务做不过来也阻塞
	wokerPool.Run()
	// 处理任务
	go func() {
		for {
			select {
			case data := <-global.ApplyFormStatusDataChan:
				sc := &Dosomething{key: data}
				wokerPool.JobQueue <- sc
			}
		}
	}()
	// TestData()
	web()
}

type Dosomething struct {
	key global.ApplyFormStatusData
}

func (d *Dosomething) Do() {
	global.Logger.Info("正在处理的数据是：", d.key)
	//处理封装对象
	switch d.key.Bizno {
	case global.Server_ApplyStatus:
		// 申请单状态处理
		object.ApplyFormStatusNotity(d.key)
	case global.Server_ApplyInfo:
		// obj.DownObject()
	}
}

func web() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	ser := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	ser.ListenAndServe()
}
