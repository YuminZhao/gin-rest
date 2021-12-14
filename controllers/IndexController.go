package controllers

import (
	"gin-rest/config"
	"gin-rest/rest/r"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Index func(c *gin.Context)
}

var IndexController = controller{
	Index: func(c *gin.Context) {

		r.Return(c, r.D{
			Data: config.Mysql,
		})
	},
}
