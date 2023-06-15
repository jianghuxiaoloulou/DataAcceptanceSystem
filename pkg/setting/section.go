package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type GeneralSettingS struct {
	LogSavePath string
	LogFileName string
	LogFileExt  string
	LogMaxSize  int
	LogMaxAge   int
	MaxThreads  int
	MaxTasks    int
}

type DatabaseSettingS struct {
	// 数据库模块修改后的参数配置
	DBMaxIdleConns int
	DBMaxOpenConns int
	DBMaxLifetime  int
	PacsDBType     string
	PacsDBConn     string
	HISSqlType     int
	MZApplyDBType  string
	MZApplyDBConn  string
	ZYApplyDBType  string
	ZYApplyDBConn  string
	TJApplyDBType  string
	TJApplyDBConn  string
}

type ObjectSettingS struct {
	InterfaceSystemType int
	HISURL              string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
