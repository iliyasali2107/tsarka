package iin_handlers

import "github.com/gin-gonic/gin"

func (ih *IINHandlers) RegisterIINRoutes(router *gin.RouterGroup) {
	iinRoute := router.Group("/iin")
	iinRoute.POST("/check", ih.CheckHandler)
}
