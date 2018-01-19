package search

import (
	"golang.org/x/net/context"
	"github.com/alex19861108/alaigo/worker/search/proto"
)

type server struct {

}

func (s *server) Search(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	return SearchItem(req.RequestId)
}

func SearchItem(req string) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{
		RequestId: req,
	}, nil
}
