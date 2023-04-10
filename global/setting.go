package global

import (
	"WowjoyProject/DataAcceptanceSystem/pkg/logger"
	"WowjoyProject/DataAcceptanceSystem/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	GeneralSetting  *setting.GeneralSettingS
	DatabaseSetting *setting.DatabaseSettingS
	ObjectSetting   *setting.ObjectSettingS
	Logger          *logger.Logger
)
