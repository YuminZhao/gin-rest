package controllers

import (
	"gin-rest/models"
	"gin-rest/rest/m"
	"gin-rest/rest/r"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Index func(c *gin.Context)
}

type person struct {
	Ager int    `form:"ager" label:"年龄r" validate:"required,gte=10,gtfield=Age"`
	Age  int    `form:"aget" label:"年龄" validate:"required,gte=10"`
	Name string `form:"name" label:"姓名" validate:"required,username"`
}

var IndexController = controller{
	Index: func(c *gin.Context) {
		var data person
		if err := r.Validate(c, &data); err.Err != nil {
			r.Error(c, 11101, err.Err.Error(), err.Data)
			return
		}
		var user models.User
		user.Name = "name"
		m.DB.Save(&user)

		r.Return(c, data)
	},
}
