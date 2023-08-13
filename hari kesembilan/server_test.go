package hari_kesembilan

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello word")
	})
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hi")
	})

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, request.Method)
		fmt.Fprint(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: handler,
	}

	server.ListenAndServe()
}
