package db

import (
	"fmt"
	"steward/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// DB 获取 mysql 连接
func DB() *gorm.DB {
	if db == nil {
		initMySQL()
	}
	return db
}

// InitMySQL 初始化 mysql
func InitMySQL() {
	initMySQL()
}

// initMySQL 初始化 mysql
func initMySQL() {
	c := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		c.Mysql.Username,
		c.Mysql.Password,
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.DB,
		c.Mysql.Charset,
		c.Mysql.ParseTime,
		c.Mysql.Loc)
	d, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         uint(c.Mysql.DefaultStringSize),
		DisableDatetimePrecision:  c.Mysql.DisableDatetimePrecision,
		DontSupportRenameIndex:    c.Mysql.DontSupportRenameIndex,
		DontSupportRenameColumn:   c.Mysql.DontSupportRenameColumn,
		SkipInitializeWithVersion: c.Mysql.SkipInitializeWithVersion,
	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("init mysql error: %w", err))
	}
	db = d
}
