package data

import (
	"fmt"
	"seckill/config"
	"seckill/model"
	"time"

	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMysql(conf config.AppConfig) *gorm.DB {
	fmt.Println("---Start initializing mysql connection---")
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbUser,
		conf.Password,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)), &gorm.Config{})

	if err != nil {
		fmt.Println("init mysql failed: ", err)
		panic(err)
	}
	sqlDB, _ := db.DB()

	// 配置数据库
	sqlDB.SetMaxIdleConns(10)                  // 设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxOpenConns(100)                 // 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 最大连接可复用时间

	// 自动迁移
	db.AutoMigrate(&model.User{}, &model.Goods{}, &model.Order{})

	fmt.Println("---Mysql connection is initialized---")
	return db
}
