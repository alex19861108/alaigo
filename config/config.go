package config

type Service struct {
	Name	string		`yaml:"name"`
	RpcPort	string		`yaml:"rpc_port"`
	HttpPort	string		`yaml:"http_port"`
}

type Config struct {
	Service	Service
}
