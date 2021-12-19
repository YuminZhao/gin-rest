package controllers

import (
	"gin-rest/rest/r"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Index func(c *gin.Context)
}

type Person struct {
	Ager float64 `form:"ager" label:"年龄r" validate:"required,gte=10,gtfield=Age"`
	Age  int     `form:"aget" label:"年龄" validate:"required,gte=10"`
	Name string  `form:"name" label:"姓名" validate:"required,username"`
}

var IndexController = controller{
	Index: func(c *gin.Context) {
		var data Person
		if err := r.Validate(c, &data); err.Err != nil {
			r.Error(c, 11101, err.Err.Error(), err.Data)
			return
		}
		if data.Name == "" {
			r.Return(c, data)
		}
		r.Return(c, data)
	},
}
