package global

import (
	"electronic-album/pkg/setting"
	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	DatabaseSetting *setting.DatabaseSettings
	DBEngine        *gorm.DB
	JwtSetting      *setting.JwtSettingS
	OSSSetting      *setting.OSSSettingS
)
