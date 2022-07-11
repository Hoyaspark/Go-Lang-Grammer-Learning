package main

import (
	"fmt"
	"log"
	"net"
	"net/netip"
)

func main() {
	ip := "127.0.0.1"
	b := []byte(ip)
	tcp, err := net.ListenTCP("tcp", net.TCPAddrFromAddrPort(netip.AddrPortFrom(netip.AddrFrom4([4]byte{b[0], b[1], b[2], b[3]}), 8888)))

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(tcp.Addr())

}
