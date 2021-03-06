package controllers

import (
	"gin-rest/rest/r"
	"gin-rest/services"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	Index  func(c *gin.Context)
	Create func(c *gin.Context)
}

var OrderController = orderController{
	Index: func(c *gin.Context) {
		var s services.OrderService
		user, err := s.GetOrder()
		if err != nil {
			r.Failed(c, 11101, err.Error())
			return
		}
		r.Return(c, user)
	},
	Create: func(c *gin.Context) {
		var data services.OrderCreate
		if err := r.Validate(c, &data); err.Err != nil {
			r.Error(c, 11103, err.Err.Error(), err.Data)
			return
		}
		var s services.OrderService
		err := s.CreateOrder(data.Price)
		if err != nil {
			r.Failed(c, 11104, err.Error())
			return
		}
		r.Success(c, "添加订单成功")
	},
}
