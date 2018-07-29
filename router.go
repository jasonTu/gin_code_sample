package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/jasonTu/gin_code_sample/apis"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/auditlog", AddAuditlogApi)

	return router
}
