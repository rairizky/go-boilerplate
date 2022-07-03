package handler

import (
	"boiler-go/constant"
	"boiler-go/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PageNotFound(c *gin.Context) {
    response := helper.APIResponse("URL page not found", http.StatusNotFound, constant.STATUS_FAILED, nil)

	c.JSON(http.StatusNotFound, response)
}
