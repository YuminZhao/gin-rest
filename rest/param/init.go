package param

import "flag"

var Cpath string

func init() {
	flag.StringVar(&Cpath, "c", "./rest.yaml", "配置文件地址")
	flag.Parse()
}
