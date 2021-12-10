package config

import "gin-rest/rest/yaml"

var Server = &yaml.ServerType{
	Port:    yaml.Server.Port,
	Mode:    yaml.Server.Mode,
	LogFile: yaml.Server.LogFile,
	Zone:    yaml.Server.Zone,
}
