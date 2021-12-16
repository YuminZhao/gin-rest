package rest

import (
	"gin-rest/config"
	"gin-rest/rest/r"
	"gin-rest/route"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	cstZone := time.FixedZone("CST", config.Server.Zone*3600)
	time.Local = cstZone

	gin.SetMode(config.Server.Mode)

	logFile, err := os.Create(config.Server.LogFile + "/http_" + ".log")
	if err != nil {
		log.Fatalln(err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(logFile)
}

func Start() {
	App := gin.Default()
	App.SetTrustedProxies(nil)
	App.NoRoute(func(c *gin.Context) {
		r.NotFound(c)
	})

	route.ApiRoute(App)

	err := App.Run(":" + strconv.Itoa(config.Server.Port))
	//s := endless.NewServer(":"+strconv.Itoa(config.Server.Port), App)
	//err := s.ListenAndServe()
	if err != nil {
		log.Fatalln("ERROR:", err.Error())
	}
}
