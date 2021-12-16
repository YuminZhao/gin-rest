package controllers

import (
	"gin-rest/config"
	"gin-rest/rest/r"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Index func(c *gin.Context)
}

type Person struct {
	Age  int    `form:"age" binding:"required,gt=10"`
	Name string `form:"name" binding:"required"`
}

var IndexController = controller{
	Index: func(c *gin.Context) {
		var data Person
		if err := c.ShouldBind(data); err != nil {
			r.Forbidden(c)
			return
		}
		r.Return(c, config.Mysql)
	},
}
