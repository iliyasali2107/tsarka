package api

import (
	counter_handlers "tsarka/internal/api/counter"
	hash_handlers "tsarka/internal/api/hash"
	iin_handlers "tsarka/internal/api/iin"
	self_handlers "tsarka/internal/api/self"
	substr_handlers "tsarka/internal/api/substr"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Counter *counter_handlers.CounterHandlers
	Substr  *substr_handlers.SubstrHandlers
	Self    *self_handlers.SelfHandlers
	Hash    *hash_handlers.HashHandlers
	IIN     *iin_handlers.IINHandlers
}

func (app *Application) RegisterRoutes(routes *gin.Engine) {
	restGroup := routes.Group("/rest")
	app.Counter.RegisterCounterRoutes(restGroup)
	app.Substr.RegisterCounterRoutes(restGroup)
	app.Self.RegisterSelfRoutes(restGroup)
	app.Hash.RegisterHashRoutes(restGroup)
	app.IIN.RegisterIINRoutes(restGroup)
}
