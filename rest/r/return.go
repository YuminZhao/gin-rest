package r

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type D struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func Return(c *gin.Context, data D) {
	c.JSON(http.StatusOK, data)
}
