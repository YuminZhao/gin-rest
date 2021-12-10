package rest

import (
	"gin-rest/config"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Start() {

	gin.SetMode(config.Server.Mode)

	logFile, err := os.Create(config.Server.LogFile + "/http_" + ".log")
	if err != nil {
		log.Fatalln(err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(logFile)

	var cstZone = time.FixedZone("CST", config.Server.Zone*3600)
	time.Local = cstZone

	App := gin.Default()
	App.SetTrustedProxies(nil)
	App.GET("/", func(c *gin.Context) {
		log.Println(time.Now().Format("2006-01-02 15:04:05"))
		c.JSON(http.StatusOK, gin.H{
			"message": "hh",
			"data":    gin.H{"1": 212},
		})
	})
	log.Println("提供服务和监听于端口:" + strconv.Itoa(config.Server.Port))
	App.Run(":" + strconv.Itoa(config.Server.Port))
}
