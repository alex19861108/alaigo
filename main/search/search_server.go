package search

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"flag"

	"../../config"
	"../../worker/search"
)


var (
	configFile = flag.String("c", "config.yaml", "yaml config path")
	cfg = config.Config{}
)

func init() {
	flag.Parse()
}

func main() {
	content, _ := ioutil.ReadFile(*configFile)
	cfg := config.Config{}
	yaml.Unmarshal(content, &cfg)

	search.InitRpcServer(cfg.Service.Port)
}
