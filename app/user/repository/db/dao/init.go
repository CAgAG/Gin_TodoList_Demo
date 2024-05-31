package dao

import (
	"TodoList_demo/conf"
	"context"
	logging "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
)

var db *gorm.DB

func NewMysqlDB() error {
	// dsn := "test_root:123456@tcp(127.0.0.1:3306)/gin_test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := strings.Join([]string{conf.DbUser, ":", conf.DbPassWord, "@tcp(", conf.DbHost, ":", conf.DbPort, ")/", conf.DbName, "?charset=utf8&parseTime=true"}, "")
	var db_logger logger.Interface = logger.Default.LogMode(logger.Info)

	db_mysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "todo_list_demo_",
			SingularTable: true,
		},
		// 打印对应的 sql 语句
		Logger: db_logger,
	})
	if err != nil {
		logging.Info("数据库连接失败")
		logging.Info(err)
		return err
	}
	// db.AutoMigrate(&User{}) // 自动同步
	db = db_mysql
	migrate()
	return nil
}

// change current instance db's context to ctx
func NewDBClient(ctx context.Context) *gorm.DB {
	_db := db
	return _db.WithContext(ctx)
}
