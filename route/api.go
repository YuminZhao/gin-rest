package route

import (
	"gin-rest/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoute(router *gin.Engine) {

	router.GET("/", controllers.IndexController.Index)
	router.POST("/", controllers.IndexController.Create)
	router.GET("/:id", controllers.IndexController.Show)

	router.GET("/order", controllers.OrderController.Index)
	router.POST("/order", controllers.OrderController.Create)
}
