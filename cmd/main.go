package main

import (
	"fmt"
	"net"

	"github.com/joshqu1985/fireman/configor"
	"github.com/joshqu1985/fireman/discover"
	"github.com/joshqu1985/fireman/log"
	"github.com/joshqu1985/fireman/store/es"
	"github.com/joshqu1985/fireman/tracing"
	"github.com/joshqu1985/fireman/transport/rpc"

	"github.com/joshqu1985/service-esearch/internal/dao/database"
	"github.com/joshqu1985/service-esearch/internal/handler"
	"github.com/joshqu1985/service-esearch/internal/service"
)

type Config struct {
	Name     string
	Port     int
	Log      log.Config
	Discover discover.Config
	Elastic  es.Config
}

var (
	Conf Config
)

func init() {
	if err := configor.Load("./configs/conf.toml", &Conf); err != nil {
		panic(err)
	}

	log.Init(Conf.Log)

	if _, err := tracing.Init(Conf.Name); err != nil {
		panic(err)
	}

	if err := discover.Init(Conf.Discover); err != nil {
		panic(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", Conf.Port))
	if err != nil {
		panic(err)
	}

	if err := discover.Register(Conf.Name, Conf.Port); err != nil {
		panic(err)
	}

	gsvr := rpc.NewUnaryServer()

	s := service.New(database.NewRepository(es.NewPool(Conf.Elastic)))

	handler.RegisterHandler(gsvr, s)

	if err := gsvr.Serve(listener); err != nil {
		panic(err)
	}
}
