package httpServer

import (
	"fmt"
	"log"
	"bytes"
	"net/http"
	"strings"
	"sync"
	"github.com/bizy01/echoServer/config"
)

type EchoServer struct{}

var wg sync.WaitGroup

func (server *EchoServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer

	// 写入首行
	headLine := fmt.Sprintf("%s %s %s\n", req.Method, req.URL.Path, req.Proto)
	buf.Write([]byte(headLine))

	// 写入header
	for key, item := range req.Header {
		line := fmt.Sprintf("%s: %s\n", key, strings.Join(item, ","))
		buf.Write([]byte(line))
	}

	buf.Write([]byte("\n"))

	// 写入body
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		log.Println("read body error", err)
	}

	fmt.Fprintf(w, buf.String())
}

func RunHttpServer(path string) {
	cfg := config.GetConfig(path)
	for _, hcfg := range cfg.HttpServer {
		svc := new(EchoServer)
		listen := fmt.Sprintf("%s:%s", hcfg.Bind, hcfg.Port)
		log.Printf("%s http server start and listen port %s...",hcfg.Name, listen)
		wg.Add(1)
		go runServer(svc, listen)
        wg.Wait()
	}
}

func runServer(svc *EchoServer, listen string) {
	log.Fatal(http.ListenAndServe(listen, svc))
	wg.Done()
}