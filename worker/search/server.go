package search

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"flag"

	"github.com/alex19861108/alaigo/config"
	"net"
	"google.golang.org/grpc/reflection"
	"net/http"
	"log"
	"google.golang.org/grpc"
	"github.com/alex19861108/alaigo/worker/search/proto"
)

var (
	pConfigFile = flag.String("c", "config.yaml", "yaml config path")
	cfg = config.Config{}
)

func searchHandler(w http.ResponseWriter, r * http.Request) {
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys) < 1 {
		w.Write([]byte("keys error"))
		return
	}
}


func defaultHandler(w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("Welcome!"))
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

func InitHttpServer(port string) {
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/", defaultHandler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}

func init() {
	flag.Parse()

	content, _ := ioutil.ReadFile(*pConfigFile)
	cfg := config.Config{}
	yaml.Unmarshal(content, &cfg)

	InitRpcServer(cfg.Service.RpcPort)
	InitHttpServer(cfg.Service.HttpPort)
}