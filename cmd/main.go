package main

import (
	"fmt"
	"net"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/joshqu1985/fireman/pkg/configor"
	"github.com/joshqu1985/fireman/pkg/consul"
	"github.com/joshqu1985/fireman/pkg/locip"
	"github.com/joshqu1985/fireman/pkg/logger"
	"github.com/joshqu1985/fireman/pkg/tracing"

	"github.com/joshqu1985/service-esearch/internal/handler"
	"github.com/joshqu1985/service-esearch/internal/store"
)

type Config struct {
	Name      string
	Port      int
	Discovery consul.Config
	Elastic   store.Config
	Logger    logger.Config
}

var (
	Conf  Config
	LocIP string
)

func init() {
	if err := configor.LoadConfig("./configs/conf.toml", &Conf); err != nil {
		panic(err)
	}

	var err error
	if LocIP, err = locip.GetLocalIP(); err != nil {
		panic(err)
	}

	if _, err := tracing.InitTracing(Conf.Name); err != nil {
		panic(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", Conf.Port))
	if err != nil {
		fmt.Println("failed to listen:", err)
		panic(err)
	}

	grpcSvr := grpc.NewServer(
		grpc.UnaryInterceptor(tracing.GrpcServerInterceptor(opentracing.GlobalTracer())),
	)

	handler.RegisterHandler(grpcSvr,
		store.NewElasticRepo(store.NewElasticClient(Conf.Elastic)),
		logger.InitLogger(Conf.Logger))

	if err := consul.NewClient(Conf.Discovery).
		Register(Conf.Name, LocIP, Conf.Port); err != nil {
		panic(err)
	}
	consul.RegisterGrpcHealth(grpcSvr)

	if err := grpcSvr.Serve(listener); err != nil {
		fmt.Println("failed to serve:", err)
		panic(err)
	}
}
