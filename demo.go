package main

import (
	"fmt"
	"log"
	"bytes"
	"net/http"
	"strings"
)

type echoServer struct{}

func (server *echoServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer

	// 写入首行
	headLine := fmt.Sprintf("%s %s %s\n\n", req.Method, req.URL.Path, req.Proto)
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

func main() {
	svc := new(echoServer)
	log.Println("echoServer start listen port 9999...")
	log.Fatal(http.ListenAndServe(":9999", svc))
}




