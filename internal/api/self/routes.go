package self_handlers

import "github.com/gin-gonic/gin"

func (sh *SelfHandlers) RegisterSelfRoutes(router *gin.RouterGroup) {
	selfRoute := router.Group("/self")
	selfRoute.GET("/find/:substr", sh.FindHandler)
}
