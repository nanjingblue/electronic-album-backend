package global

import (
	"electronic-album/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	DatabaseSetting *setting.DatabaseSettings
	DBEngine        *gorm.DB
)
