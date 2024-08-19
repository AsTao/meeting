package conf

import (
	"time"

	"github.com/AsTao/meeting/global"
	"github.com/AsTao/meeting/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() {
	mode := logger.Info
	if !viper.GetBool("mode.develop") {
		mode = logger.Error
	}
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "meeting",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(mode),
	})

	if err != nil {
		global.Logger.Error(err)
		panic(err.Error())
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&model.User{})

	global.Db = db

}
