package main

import (
	"log"

	"tsarka/internal/api"
	counter_handlers "tsarka/internal/api/counter"
	hash_handlers "tsarka/internal/api/hash"
	iin_handlers "tsarka/internal/api/iin"
	self_handlers "tsarka/internal/api/self"
	substr_handlers "tsarka/internal/api/substr"
	"tsarka/internal/config"
	"tsarka/internal/db"
	"tsarka/internal/repository"
	"tsarka/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to config", err)
	}

	db := db.InitRedisDB(cfg)

	// counter
	cr := repository.NewCounterRepository(db)
	cs := service.NewCounterService(cr)
	ch := counter_handlers.NewCounterHandlers(cs)

	// substr
	ss := service.NewSubstrService()
	sh := substr_handlers.NewSubstHandler(ss)

	// self
	selfS := service.NewSelfService()
	selfH := self_handlers.NewSelfHandlers(selfS)

	// hash
	hs := service.NewHashService(cfg)
	hh := hash_handlers.NewHashHandlers(hs)
	go hh.RequestsHandler()

	// iin
	is := service.NewIINService()
	ih := iin_handlers.NewIINHandlers(is)

	// routes
	routes := gin.Default()

	// app
	app := api.Application{Counter: ch, Substr: sh, Self: selfH, Hash: hh, IIN: ih}
	app.RegisterRoutes(routes)

	routes.Run(cfg.ServerUrl)
}
