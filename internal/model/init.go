package model

import (
	"electronic-album/global"
	"electronic-album/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(databaseSetting.DBType, args)
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)

	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

func SetupDBEngine() error {
	var err error
	global.DBEngine, err = NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	_ = global.DBEngine.AutoMigrate(&User{}) // 自动化更新
	_ = global.DBEngine.AutoMigrate(&Album{})
	_ = global.DBEngine.AutoMigrate(&Picture{})
	_ = global.DBEngine.AutoMigrate(&Friend{})
	return nil
}
