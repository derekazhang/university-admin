package common

import (
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDb *gorm.DB
var gormdbOnce sync.Once

func GetGorm() *gorm.DB {
	gormdbOnce.Do(func() {
		conf := &gorm.Config{}
		if GetConfig().Runmode == "dev" {
			conf = &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			}
		}
		db, err := gorm.Open(mysql.Open(GetConfig().Mysql.Master.Addr), conf)
		if err != nil {
			panic(err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)

		gormDb = db
	})
	return gormDb
}
