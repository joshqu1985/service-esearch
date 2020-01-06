package handler

import (
	"google.golang.org/grpc"

	"github.com/joshqu1985/protos/pkg/pb"
	"github.com/joshqu1985/service-esearch/internal/service"
)

func RegisterHandler(grpcSvr *grpc.Server, s *service.Service) {
	pb.RegisterEsearchServer(grpcSvr, &Handler{s})
}

type Handler struct {
	Service *service.Service
}
