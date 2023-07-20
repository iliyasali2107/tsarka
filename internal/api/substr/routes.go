package substr_handlers

import "github.com/gin-gonic/gin"

func (sh *SubstrHandlers) RegisterCounterRoutes(router *gin.RouterGroup) {
	counterRoute := router.Group("/substr")
	counterRoute.GET("/find", sh.FindHandler)
}
