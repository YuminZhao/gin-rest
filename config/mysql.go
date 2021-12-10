package config

import (
	"gin-rest/rest/yaml"
)

var Mysql = map[string]yaml.MysqlType{
	"default": {
		User: yaml.Mysql["default"].User,
		Pass: yaml.Mysql["default"].Pass,
	},
	"account": {
		User: yaml.Mysql["account"].User,
		Pass: yaml.Mysql["account"].Pass,
	},
}
