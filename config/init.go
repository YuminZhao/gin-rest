package config

import "time"

func init() {
	cstZone := time.FixedZone("CST", Server.Zone*3600)
	time.Local = cstZone
}
