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

var IndexController = indexController{
	Index: func(c *gin.Context) {
		var s services.UserService
		user, err := s.GetUser()
		if err != nil {
			r.Failed(c, 11101, err.Error())
			return
		}
		r.Return(c, user)
	},
	Create: func(c *gin.Context) {
		var data services.UserCreate
		if err := r.Validate(c, &data); err.Err != nil {
			r.Error(c, 11101, err.Err.Error(), err.Data)
			return
		}
		var s services.UserService
		err := s.CreateUser(data.Name)
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
			r.NotFound(c)
			return
		}
		var s services.UserService
		user, err := s.GetUserById(uint(id))
		if err != nil {
			r.Failed(c, 11101, err.Error())
			return
		}
		r.Return(c, user)
	},
}
