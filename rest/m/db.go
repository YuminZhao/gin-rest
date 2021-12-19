package m

import (
	"database/sql"
	"fmt"
	"gin-rest/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var SqlDB *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Mysql.User, config.Mysql.Pass, config.Mysql.Host, config.Mysql.Port, config.Mysql.Database, config.Mysql.Charset)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	SqlDB, err = DB.DB()
	SqlDB.SetMaxIdleConns(config.Mysql.MaxIdleConns)
	SqlDB.SetMaxOpenConns(config.Mysql.MaxOpenConns)
	SqlDB.SetConnMaxLifetime(time.Second * config.Mysql.ConnMaxLifetime)
}
