package main

import (
	"flag"
	"fmt"

	"github.com/armon/go-socks5"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 1080, "listen port")
	flag.Parse()

	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", fmt.Sprintf(":%v", port)); err != nil {
		panic(err)
	}
}
