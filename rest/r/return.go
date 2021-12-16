package r

import (
	"gin-rest/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type D struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func Return(c *gin.Context, data interface{}) {
	var message string
	if config.Server.Language == "zh" {
		message = "请求成功"
	} else {
		message = "Request succeeded"
	}
	c.JSON(http.StatusOK, D{
		Code: 0,
		Msg:  message,
		Data: data,
	})
}

func Success(c *gin.Context, message string) {
	c.JSON(http.StatusCreated, D{
		Code: 0,
		Msg:  message,
	})
}

func Failed(c *gin.Context, code int, message string) {
	c.JSON(http.StatusExpectationFailed, D{
		Code: code,
		Msg:  message,
	})
}

func Deleted(c *gin.Context, message string) {
	c.JSON(http.StatusNoContent, D{
		Code: 0,
		Msg:  message,
	})
}

func Unauthorized(c *gin.Context) {
	var message string
	if config.Server.Language == "zh" {
		message = "认证错误"
	} else {
		message = "Authentication error"
	}
	c.JSON(http.StatusUnauthorized, D{
		Code: 401,
		Msg:  message,
	})
}

func Forbidden(c *gin.Context) {
	var message string
	if config.Server.Language == "zh" {
		message = "没有权限"
	} else {
		message = "No permission"
	}
	c.JSON(http.StatusForbidden, D{
		Code: 403,
		Msg:  message,
	})
}

func NotFound(c *gin.Context) {
	var message string
	if config.Server.Language == "zh" {
		message = "找不到请求"
	} else {
		message = "Request not found"
	}
	c.JSON(http.StatusNotFound, D{
		Code: 404,
		Msg:  message,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusUnprocessableEntity, D{
		Code: code,
		Msg:  message,
	})
}
