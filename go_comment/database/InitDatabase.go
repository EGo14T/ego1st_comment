package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var db *gorm.DB
var err error

var mysqlInstance *ConnectPool
var mysqlOnce sync.Once

type ConnectPool struct {
}

func InitDB() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/xinmusic_new?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info), // gorm日志模式：silent
		DisableForeignKeyConstraintWhenMigrating: true,                                  // 外键约束
		SkipDefaultTransaction:                   true,                                  // 禁用默认事务（提高运行速度）
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

func GetMysqlInstance() *ConnectPool {
	mysqlOnce.Do(func() {
		mysqlInstance = &ConnectPool{}
	})
	return mysqlInstance
}

/*
* @fuc  对外获取数据库连接对象db
 */
func (m *ConnectPool) GetMysqlPool() *gorm.DB {
	//db.LogMode(true)
	return db
}

func GetMysqlDb() (db *gorm.DB) {
	return GetMysqlInstance().GetMysqlPool()
}
