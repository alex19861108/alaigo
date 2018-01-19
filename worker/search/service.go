package search

import (
	"golang.org/x/net/context"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"fmt"
	"github.com/alex19861108/alaigo/worker/search/proto"
)

type server struct {

}

func (s *server) Search(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	request_id := req.RequestId
	fmt.Printf("receive: %s\n", request_id)
	return &proto.SearchResponse{
		RequestId: request_id,
	}, nil
}

func InitRpcServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterSearchServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
