package config

type Service struct {
	Name	string		`yaml:"name"`
	Port	string		`yaml:"port"`
}

type Config struct {
	Service	Service
}
