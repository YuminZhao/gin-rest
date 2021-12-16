package route

import (
	"gin-rest/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoute(router *gin.Engine) {

	router.POST("/", controllers.IndexController.Index)
}
