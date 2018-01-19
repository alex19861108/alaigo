package search

import (
	"google.golang.org/grpc"
	"github.com/prometheus/common/log"
	search_proto "../../worker/search/proto"
	"context"
	"fmt"
)

const (
	address		= "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()

	client := search_proto.NewSearchServiceClient(conn)

	response, _ := client.Search(context.Background(), &search_proto.SearchRequest{
		RequestId: "abc",
	})
	fmt.Printf("%s\n", response.RequestId)
}