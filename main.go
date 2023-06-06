package main

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	"WowjoyProject/DataAcceptanceSystem/internal/routers"
	"WowjoyProject/DataAcceptanceSystem/pkg/logger"
	"WowjoyProject/DataAcceptanceSystem/pkg/object"
	"WowjoyProject/DataAcceptanceSystem/pkg/setting"
	"WowjoyProject/DataAcceptanceSystem/pkg/workpattern"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

// @title PACS集成平台
// @version 1.0.0.0
// @description PACS集成平台
// @termsOfService https://github.com/jianghuxiaoloulou/DataAcceptanceSystem.git
func main() {
	readSetup()
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

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("General", &global.GeneralSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Object", &global.ObjectSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	lunberLogger := &lumberjack.Logger{
		Filename:  global.GeneralSetting.LogSavePath + "/" + global.GeneralSetting.LogFileName + global.GeneralSetting.LogFileExt,
		MaxSize:   global.GeneralSetting.LogMaxSize,
		MaxAge:    global.GeneralSetting.LogMaxAge,
		LocalTime: true,
	}
	global.Logger = logger.NewLogger(io.MultiWriter(lunberLogger, os.Stdout), "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupReadDBEngine() error {
	var err error
	global.ReadDBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupWriteDBEngine() error {
	var err error
	global.WriteDBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupOracleDBEngine() error {
	var err error
	global.OracleDBEngine, err = model.NewOracleDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func readSetup() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupReadDBEngine()
	if err != nil {
		log.Fatalf("init.setupReadDBEngine err: %v", err)
	}
	err = setupWriteDBEngine()
	if err != nil {
		log.Fatalf("init.setupWriteDBEngine err: %v", err)
	}
	err = setupOracleDBEngine()
	if err != nil {
		log.Fatalf("init.setupOracleDBEngine err: %v", err)
	}
}
