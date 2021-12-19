package route

import (
	"gin-rest/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoute(router *gin.Engine) {

	router.GET("/", controllers.IndexController.Index)
}
