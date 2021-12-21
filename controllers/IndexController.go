package controllers

import (
	"gin-rest/rest/r"
	"gin-rest/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type indexController struct {
	Index  func(c *gin.Context)
	Create func(c *gin.Context)
	Show   func(c *gin.Context)
}

type create struct {
	Name string `form:"name" label:"姓名" validate:"required,username"`
}

var IndexController = indexController{
	Index: func(c *gin.Context) {
		user, err := services.UserService.GetUser()
		if err != nil {
			r.Failed(c, 11101, err.Error())
			return
		}
		r.Return(c, user)
	},
	Create: func(c *gin.Context) {
		var data create
		if err := r.Validate(c, &data); err.Err != nil {
			r.Error(c, 11101, err.Err.Error(), err.Data)
			return
		}
		err := services.UserService.CreateUser(data.Name)
		if err != nil {
			r.Failed(c, 11102, err.Error())
			return
		}
		r.Success(c, "添加用户成功")
	},
	Show: func(c *gin.Context) {
		pid := c.Param("id")
		id, err := strconv.Atoi(pid)
		if err != nil {
			r.Failed(c, 11101, "请输入正确的数字")
			return
		}
		user, err := services.UserService.GetUserById(uint(id))
		if err != nil {
			r.Failed(c, 11101, err.Error())
			return
		}
		r.Return(c, user)
	},
}
