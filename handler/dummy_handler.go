package handler

import (
	"boiler-go/constant"
	"boiler-go/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloDummy(c *gin.Context) {
	helloText := helper.APIResponse("Access this page status ok", http.StatusOK, constant.STATUS_SUCCESS, nil)

	c.JSON(http.StatusOK, helloText)
}
