package database

import (
	"fmt"
	"gin_api_frame/config"
	"gin_api_frame/internal/model"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

type Database struct {
	Mysql *gorm.DB
}

func NewDatabase(config *config.Config,gl *GormLogger) *Database {
	dbUser := config.Mysql.DbUser
	dbPassWord := config.Mysql.DbPassWord
	dbHost := config.Mysql.DbHost
	dbPort := config.Mysql.DbPort
	dbName := config.Mysql.DbName

	var database Database

	pathRead := strings.Join([]string{dbUser, ":", dbPassWord, "@tcp(", dbHost, ":", dbPort, ")/", dbName, "?charset=utf8&parseTime=true"}, "")
	pathWrite := strings.Join([]string{dbUser, ":", dbPassWord, "@tcp(", dbHost, ":", dbPort, ")/", dbName, "?charset=utf8&parseTime=true"}, "")

	gl.Debug = true
	gl.SkipErrRecordNotFound = true

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       pathRead, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{
		Logger: gl,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	database.Mysql = db
	_ = database.Mysql.Use(dbresolver.
		Register(dbresolver.Config{
			// `db2` 作为 sources，`db3`、`db4` 作为 replicas
			Sources:  []gorm.Dialector{mysql.Open(pathRead)},                         // 写操作
			Replicas: []gorm.Dialector{mysql.Open(pathWrite), mysql.Open(pathWrite)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},                                      // sources/replicas 负载均衡策略
		}))
	err = database.Mysql.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
		)
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
	return &database

}


