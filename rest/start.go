package rest

import (
	"gin-rest/config"
	"gin-rest/rest/m"
	"gin-rest/rest/r"
	"gin-rest/route"
	"io/ioutil"
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
	gin.DefaultWriter = ioutil.Discard

	file, _ := os.Create(config.Server.LogFile + "/http.log")

	log.SetPrefix("[GIN-REST] ")
	log.SetOutput(file)
}

func Start() {
	defer m.SqlDB.Close()
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
