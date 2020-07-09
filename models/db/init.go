/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package db

import (
	"shortlink/config"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	_DB *gorm.DB
)

func Init() {
	var err error
	_DB, err = gorm.Open("mysql", config.Config().Mysql.Host)
	if err != nil {
		panic(err)
	}

	_DB.LogMode(false)
	_DB.DB().SetConnMaxLifetime(5 * time.Second)
	_DB.DB().SetMaxIdleConns(config.Config().Mysql.MaxIdleConn)
	_DB.DB().SetMaxOpenConns(config.Config().Mysql.MaxOpenConn)

	_DB.SingularTable(true)
	_DB.AutoMigrate(&TBShortLink{})
	_DB.AutoMigrate(&TBCustomizeShortLink{})
}
