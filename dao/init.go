package dao

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	_ = ormLogger	//虚拟分配。。不然报错啊hhhh
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true, //禁止datetime精度，mysql5.6前不支持
		DontSupportRenameIndex:    true, //重命名索引，需要把索引先删掉再重建，mysql5.7之前不支持
		DontSupportRenameColumn:   true, //用change重命名列，mysql8之前的数据库不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池
	sqlDB.SetMaxOpenConns(100) //打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	//主从配置
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(connWrite)},                      //write
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, //read
		Policy:   dbresolver.RandomPolicy{},
	}))
	Migration()
}

func NewDBclient (ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
	
}