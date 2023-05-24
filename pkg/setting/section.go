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
	DBConn       string
	DBType       string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  int
	HISSqlType   int
	OracleDBType string
	OracleDBConn string
	OracleDBView string
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
