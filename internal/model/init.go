package model

import (
	"electronic-gallery/global"
	"electronic-gallery/pkg/setting"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	var db *gorm.DB
	var err error
	if global.DatabaseSetting.DBType == "mysql" {
		db, err = gorm.Open(mysql.Open(args), &gorm.Config{})
	} else {
		if global.DatabaseSetting.SqliteDB == "" {
			global.DatabaseSetting.SqliteDB = "gallery.db"
		}
		db, err = gorm.Open(sqlite.Open(global.DatabaseSetting.SqliteDB), &gorm.Config{})
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SetupDBEngine() error {
	var err error
	global.DBEngine, err = NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	_ = global.DBEngine.AutoMigrate(&User{}) // 自动化更新
	_ = global.DBEngine.AutoMigrate(&Gallery{})
	_ = global.DBEngine.AutoMigrate(&Picture{})
	_ = global.DBEngine.AutoMigrate(&Friend{})
	_ = global.DBEngine.AutoMigrate(&Post{})
	_ = global.DBEngine.AutoMigrate(&Comment{})
	_ = global.DBEngine.AutoMigrate(&UserPost{})
	return nil
}
