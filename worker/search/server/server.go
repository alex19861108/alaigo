package server

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"flag"

	"github.com/alex19861108/alaigo/config"
	"github.com/alex19861108/alaigo/worker/search"
)

var (
	configFile = flag.String("c", "config.yaml", "yaml config path")
	cfg = config.Config{}
)

func init() {
	content, _ := ioutil.ReadFile(*configFile)
	cfg := config.Config{}
	yaml.Unmarshal(content, &cfg)

	search.InitRpcServer(cfg.Service.Port)
}