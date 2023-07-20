package hash_handlers

import "github.com/gin-gonic/gin"

func (hh *HashHandlers) RegisterHashRoutes(router *gin.RouterGroup) {
	hashRoute := router.Group("/hash")
	hashRoute.POST("/calc", hh.CalcHandler)
	hashRoute.GET("/result/:id", hh.ResultHandler)
}
