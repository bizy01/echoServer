package main

import (
	"io/ioutil"
	"flag"
	"log"
	"github.com/bizy01/echoServer/config"
	"github.com/bizy01/echoServer/core/httpServer"
)

var (
	cfgPath = flag.String("c", "./server.toml", "config path")
    simpleCfg = flag.Bool("init", false, "init config")
    Config *config.EchoServer
)

func main() {
	flag.Parse()

    // init config
	if *simpleCfg  {
		log.Println("init config")

		context := config.InitConfig()
		if err := ioutil.WriteFile("./demo.toml", context, 0644); err != nil {
	        log.Fatalf("WriteFile failure, err=[%v]\n", err)
	    }
	}

	if *cfgPath != "" {
		httpServer.RunHttpServer(*cfgPath)
	}
}


