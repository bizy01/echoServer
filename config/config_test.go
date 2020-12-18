package config

import (
	"testing"
	"reflect"
)

// 疑问点指针类型equal(todo)
func TestGetConfig(t *testing.T) {
	want := GetConfig("./demo.toml")
	expect := &EchoServer{
		HttpServer: []*Server{
			&Server{
				Bind: "0.0.0.0",
				Port: "9090",
				Mock: false,
			},
		},
		GrpcServer: []Server{
			&Server{
				Bind: "0.0.0.0",
				Port: "9091",
				Mock: false,
			},
		},
	}

	if !reflect.DeepEqual(expect, want) {
		t.Fatalf("I want %v, but get %v", expect, want)
	}

	t.Log("OK")
}

func TestInitConfig(t *testing.T) {
	want := InitConfig()

	t.Log("OK")
}