package main

import (
	"context"
	"fmt"
	"net"

	"github.com/pion/mdns"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", mdns.DefaultAddress)
	if err != nil {
		panic(err)
	}

	l, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	server, err := mdns.Server(l, nil)
	if err != nil {
		panic(err)
	}
	answer, src := server.Query(context.TODO(), "pion-test.local")
	fmt.Println(answer)
	fmt.Println(src)
}
