package yaml

import (
	"gin-rest/rest/param"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type ConfigType struct {
	Server ServerType `ymal:"server"`
	Mysql  []ConnType `yaml:"mysql"`
}
type ServerType struct {
	Port    int    `yaml:"port"`
	Mode    string `yaml:"mode"`
	LogFile string `yaml:"log_file"`
	Zone    int    `yaml:"zone"`
}
type ConnType struct {
	Name string    `yaml:"name"`
	Conn MysqlType `yaml:"conn"`
}
type MysqlType struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

var Server *ServerType
var Mysql map[string]MysqlType

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

	log.Println("配置文件" + param.Cpath + "加载成功")

	Server = &config.Server
	mysql := make(map[string]MysqlType)
	for _, v := range config.Mysql {
		mysql[v.Name] = v.Conn
	}
	Mysql = mysql
}
