package main

import (
	"log"
	"net"
)

func main() {
	pc, err := net.ListenPacket("udp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()
	buf := make([]byte, 10)
	for {
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("read from: %s length: %d content: %s", addr.String(), n, string(buf))
		log.Println("len", len(buf))
		log.Println("cap", cap(buf))
	}
}
