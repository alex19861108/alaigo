package service

import (
	"golang.org/x/net/context"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type SearchServer struct {

}

func (s *SearchServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	request_id := req.RequestId
	return &SearchResponse{
		RequestId: request_id,
	}, nil
}

func InitRpcServer(cfg config.Config) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterSearchServiceServer(s, &SearchServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
