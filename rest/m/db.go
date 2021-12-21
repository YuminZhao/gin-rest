package m

import (
	"database/sql"
	"fmt"
	"gin-rest/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error
var SqlDB *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Mysql.User, config.Mysql.Pass, config.Mysql.Host, config.Mysql.Port, config.Mysql.Database, config.Mysql.Charset)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      false,
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	SqlDB, err = DB.DB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	SqlDB.SetMaxIdleConns(config.Mysql.MaxIdleConns)
	SqlDB.SetMaxOpenConns(config.Mysql.MaxOpenConns)
	SqlDB.SetConnMaxLifetime(time.Second * config.Mysql.ConnMaxLifetime)
}
