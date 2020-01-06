package main

import (
	"fmt"
	"net"

	"github.com/joshqu1985/fireman/pkg/configor"
	"github.com/joshqu1985/fireman/pkg/discover"
	"github.com/joshqu1985/fireman/pkg/log"
	"github.com/joshqu1985/fireman/pkg/store/elasticsearch"
	"github.com/joshqu1985/fireman/pkg/tracing"
	"github.com/joshqu1985/fireman/pkg/transport/rpc"

	"github.com/joshqu1985/service-esearch/internal/dao/database"
	"github.com/joshqu1985/service-esearch/internal/handler"
	"github.com/joshqu1985/service-esearch/internal/service"
)

type Config struct {
	Name     string
	Port     int
	Discover discover.Config
	Elastic  elasticsearch.Config
	Log      log.Config
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
	gsvr := rpc.NewUnaryServer()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", Conf.Port))
	if err != nil {
		panic(err)
	}

	if err := discover.Register(Conf.Name, Conf.Port); err != nil {
		panic(err)
	}

	s := service.New(database.NewRepository(elasticsearch.NewPool(Conf.Elastic)))

	handler.RegisterHandler(gsvr, s)

	if err := gsvr.Serve(listener); err != nil {
		panic(err)
	}
}
