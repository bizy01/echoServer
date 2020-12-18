package config

import (
	"github.com/BurntSushi/toml"
	"sync"
	"path/filepath"
)

const simpleCfg  = `
[echoServer]
[[httpServer]]
   name = "httpServer"
   bind = "0.0.0.0"
   port = "9090"
   mock = false

[[grpcServer]]
   name = "grpcServer"
   bind = "0.0.0.0"
   port = "9091"
   mock = false
`

type EchoServer struct {
	HttpServer	[]*Server   `toml:"httpServer"`
	GrpcServer  []*Server   `toml:"grpcServer"`
}

type Server struct {
	Name string   `toml:"name"`
	Bind string   `toml:"bind"`
	Port string   `toml:"port"`
	Mock bool     `toml:"mock"`
}

var (
	cfg *EchoServer
	once sync.Once
)

func GetConfig(path string) *EchoServer {
	once.Do(func() {
		filePath, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}
		if _ , err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})

	return cfg
}

func InitConfig() []byte {
   return []byte(simpleCfg)
}