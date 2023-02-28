package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(conn string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,
		DefaultStringSize:         256,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	fmt.Println("here")
	if err != nil {
		fmt.Println(conn)
		fmt.Println("SSSSS22222")
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)                 //设置链接数
	sqlDB.SetMaxOpenConns(1000)                //打开链接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) //链接生命周期

	_db = db

	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(conn)},
		Replicas: []gorm.Dialector{mysql.Open(conn)},
		Policy:   dbresolver.RandomPolicy{},
	}))

	Migrate()
}

// 复用
func NewDbClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
