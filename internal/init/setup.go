package init

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/internal/model"
	"WowjoyProject/DataAcceptanceSystem/pkg/logger"
	"WowjoyProject/DataAcceptanceSystem/pkg/setting"
	"io"
	"log"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func InitSetup() {
	ReadSetup()
}

func SetupSetting() error {
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

func SetupLogger() error {
	lunberLogger := &lumberjack.Logger{
		Filename:  global.GeneralSetting.LogSavePath + "/" + global.GeneralSetting.LogFileName + global.GeneralSetting.LogFileExt,
		MaxSize:   global.GeneralSetting.LogMaxSize,
		MaxAge:    global.GeneralSetting.LogMaxAge,
		LocalTime: true,
	}
	global.Logger = logger.NewLogger(io.MultiWriter(lunberLogger, os.Stdout), "", log.LstdFlags).WithCaller(2)
	return nil
}

func SetupPacsDBEngine() error {
	var err error
	global.PacsDBEngine, err = model.NewPacsDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func SetupMZApplyDBEngine() error {
	var err error
	global.MZApplyDBEngine, err = model.NewMZApplyDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func SetupZYApplyDBEngine() error {
	var err error
	global.ZYApplyDBEngine, err = model.NewZYApplyDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func SetupTJApplyDBEngine() error {
	var err error
	global.TJApplyDBEngine, err = model.NewTJApplyDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func ReadSetup() {
	err := SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = SetupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = SetupPacsDBEngine()
	if err != nil {
		log.Fatalf("init.setupReadDBEngine err: %v", err)
	}
	err = SetupMZApplyDBEngine()
	if err != nil {
		log.Fatalf("init.setupWriteDBEngine err: %v", err)
	}
	err = SetupZYApplyDBEngine()
	if err != nil {
		log.Fatalf("init.setupOracleDBEngine err: %v", err)
	}
	err = SetupTJApplyDBEngine()
	if err != nil {
		log.Fatalf("init.setupOracleDBEngine err: %v", err)
	}
}
