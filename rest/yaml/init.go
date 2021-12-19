package yaml

import (
	"gin-rest/rest/param"
	"io/ioutil"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type ConfigType struct {
	Server ServerType `ymal:"server"`
	Mysql  MysqlType  `yaml:"mysql"`
}
type ServerType struct {
	Port     int    `yaml:"port"`
	Mode     string `yaml:"mode"`
	LogFile  string `yaml:"log_file"`
	Zone     int    `yaml:"zone"`
	Language string `yaml:"language"`
}
type MysqlType struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Pass            string        `yaml:"pass"`
	Database        string        `yaml:"database"`
	Charset         string        `yaml:"charset"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

var (
	Server *ServerType
	Mysql  *MysqlType
)

func init() {
	yamlFile, err := ioutil.ReadFile(param.Cpath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var config *ConfigType
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	Server = &config.Server
	Mysql = &config.Mysql
}
