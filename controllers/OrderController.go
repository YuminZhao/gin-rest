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

type createOrder struct {
	Price uint `form:"price" label:"价格" validate:"required"`
}

var OrderController = orderController{
	Index: func(c *gin.Context) {
		user, err := services.OrderService.GetOrder()
		if err != nil {
			r.Failed(c, 11101, err.Error())
			return
		}
		r.Return(c, user)
	},
	Create: func(c *gin.Context) {
		var data createOrder
		if err := r.Validate(c, &data); err.Err != nil {
			r.Error(c, 11103, err.Err.Error(), err.Data)
			return
		}
		err := services.OrderService.CreateOrder(data.Price)
		if err != nil {
			r.Failed(c, 11104, err.Error())
			return
		}
		r.Success(c, "添加订单成功")
	},
}
