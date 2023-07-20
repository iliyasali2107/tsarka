package counter_handlers

import "github.com/gin-gonic/gin"

func (ch *CounterHandlers) RegisterCounterRoutes(router *gin.RouterGroup) {
	counterRoute := router.Group("/counter")
	counterRoute.PUT("/add/:i", ch.IncrementHandler)
	counterRoute.PUT("/sub/:i", ch.DecrementHandler)
	counterRoute.GET("/val", ch.GetValueHandler)
}
