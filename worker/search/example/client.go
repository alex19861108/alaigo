package example

import (
	"google.golang.org/grpc"
	"github.com/prometheus/common/log"
	"github.com/alex19861108/alaigo/worker/search/proto"
	"context"
	"fmt"
)

const (
	address		= "localhost:8080"
)

func init() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewSearchServiceClient(conn)

	response, _ := client.Search(context.Background(), &proto.SearchRequest{
		RequestId: "abc",
	})
	fmt.Printf("%s\n", response.RequestId)
}
