package config

import (
	"gin-rest/rest/yaml"
)

var Mysql = yaml.MysqlType{
	Host:     yaml.Mysql.Host,
	Port:     yaml.Mysql.Port,
	User:     yaml.Mysql.User,
	Pass:     yaml.Mysql.Pass,
	Database: yaml.Mysql.Database,
	Charset:  yaml.Mysql.Charset,
}
